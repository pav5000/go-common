package longdur

import (
	"encoding/json"
)

var _ json.Unmarshaler = &Duration{}

func (d *Duration) UnmarshalJSON(rawValue []byte) error {
	var str string
	err := json.Unmarshal(rawValue, &str)
	if err != nil {
		return err
	}

	parsed, err := Parse(str)
	if err != nil {
		return err
	}

	*d = parsed
	return nil
}

var _ json.Marshaler = Duration{}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
