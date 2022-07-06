
all: test readme index

test:
	go test -v

readme:
	goreadme -import-path github.com/i2p-pt/torrc \
		-factories \
		-methods \
		-types \
		-badge-godoc \
		-badge-goreportcard > README.md

index:
	pandoc -s -t html \
		-c ./css/style.css \
		--highlight-style=tango \
		--metadata title="torrc Editor for Go" -o index.html README.md