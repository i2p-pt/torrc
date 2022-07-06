
all: test readme index

test:
	go test -v

readme:
	goreadme -import-path github.com/i2p-pt/torrc -methods -types -badge-godoc -badge-goreportcard > README.md

index:
	pandoc -s -t html \
		-c ./css/style.css \
		--highlight-style=tango \
		--metadata title="I2P Pluggable Transport" -o index.html README.md