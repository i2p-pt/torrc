package torrc

import (
	"testing"
)

func TestTorRCRead(t *testing.T) {
	// read torrc.test into a TorRC struct
	torrc, err := ReadTorRC("torrc.test")
	if err != nil {
		t.Error(err)
	}
	t.Log(torrc.String())
	if err = torrc.Write("torrc.test.out"); err != nil {
		t.Error(err)
	}
}

func TestTorRCEditClient(t *testing.T) {
	torrc, err := ReadTorRC("torrc.test")
	if err != nil {
		t.Error(err)
	}
	torrc.SetField("UseBridges", []string{"1"})
	torrc.SetField("Bridge", []string{"i2p", "i2pbase32addressesarefiftytwocharacterslongenoughok.b32.i2p"})
	torrc.SetField("ClientTransportPlugin", []string{"i2p", "exec", "i2p-client"})
	t.Log(torrc.String())
	if err = torrc.Write("torrc.test.i2p-client"); err != nil {
		t.Error(err)
	}
}

func TestTorRCEditBridge(t *testing.T) {
	torrc, err := ReadTorRC("torrc.test")
	if err != nil {
		t.Error(err)
	}
	torrc.SetField("BridgeRelay", []string{"1"})
	torrc.SetField("ORPort", []string{"9001"})
	torrc.SetField("ExtORPort", []string{"7689"})
	torrc.SetField("ServerTransportPlugin", []string{"i2p", "exec", "i2p-server"})
	t.Log(torrc.String())
	if err = torrc.Write("torrc.test.i2p-bridge"); err != nil {
		t.Error(err)
	}
}
