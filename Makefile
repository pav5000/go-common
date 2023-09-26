BINDIR=$(CURDIR)/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
LINTVER=v1.54.2
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}

precommit: format test lint

test:
	go test -timeout 20s ./...

lint: install-lint
	${LINTBIN} run

format: install-smartimports
	${SMARTIMPORTS}

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

bindir:
	mkdir -p ${BINDIR}
