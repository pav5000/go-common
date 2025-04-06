// srvwrapper is handy when you need to write simple JSON-RPC handlers with POST requests only.
// Just create a function that receives some struct that have Validate() method
// and returns any JSON-compatible struct and wrap it with New(...).
// srvwrapper will handle all the boilerplate for you.
// You may return ErrNotFound or ErrBadRequest to change the default 500 error code
package srvwrapper

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler[Req Validatable, Res any] struct {
	handleFn func(context.Context, Req) (Res, error)

	emptyRequestMode bool
}

type Validatable interface {
	Validate() error
}

type EmptyRequest struct{}

func (e EmptyRequest) Validate() error {
	return nil
}

func New[Req Validatable, Res any](handleFn func(context.Context, Req) (Res, error)) *Handler[Req, Res] {
	var req Req
	var reqCast interface{} = req
	_, emptyReqMode := reqCast.(EmptyRequest)
	return &Handler[Req, Res]{
		handleFn: handleFn,

		emptyRequestMode: emptyReqMode,
	}
}

func (h Handler[Req, Res]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("only POST allowed here"))
		return
	}

	var req Req
	if !h.emptyRequestMode {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeErrorText(w, "unmarshaling JSON", err)
			return
		}

		err = req.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeErrorText(w, "validating request", err)
			return
		}
	}

	res, err := h.handleFn(ctx, req)
	if err != nil {
		statusCode := CodeFromError(err)
		w.WriteHeader(statusCode)
		writeErrorText(w, http.StatusText(statusCode), err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(res)
	if err != nil {
		slog.Error("cannot marshal JSON", "error", err)
	}
}

func writeErrorText(w http.ResponseWriter, text string, err error) {
	buf := bytes.NewBufferString(text)
	buf.WriteString(": ")
	buf.WriteString(err.Error())
	buf.WriteByte('\n')

	_, _ = w.Write(buf.Bytes())
}
