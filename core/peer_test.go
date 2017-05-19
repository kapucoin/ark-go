package core

import (
	"log"
	"testing"
)

func TestListPeers(t *testing.T) {
	arkapi := NewArkClient(nil)

	params := PeerQueryParams{Status: "OK", Limit: "10"}

	peers, _, err := arkapi.ListPeers(&params)
	if peers.Success {
		log.Println(t.Name(), "Success, returned ", len(peers.Peers), "peers")

	} else {
		t.Error(err.Error())
	}
}