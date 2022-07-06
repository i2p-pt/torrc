package torrc

import (
	"testing"
)

func TestTorRCRead(t testing.T) {
	// read torrc.test into a TorRC struct
	torrc, err := ReadTorRC("torrc.test")
	if err != nil {
		t.Error(err)
	}
	t.Log(torrc.String())
}
