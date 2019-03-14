package api_test

import (
	"encoding/json"
	"reflect"
	"strings"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mullvad/wireguard-manager/api"
)

var fixture = api.WireguardPeerList{
	api.WireguardPeer{
		IPLeastsig: 1,
		Ports:      []int{1234, 4321},
		Pubkey:     strings.Repeat("a", 44),
	},
}

func TestAPI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		bytes, _ := json.Marshal(fixture)
		rw.Write(bytes)
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := api.API{
		BaseURL: server.URL,
		Client:  server.Client(),
	}

	peers, err := api.GetWireguardPeers()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if !reflect.DeepEqual(peers, fixture) {
		t.Errorf("got unexpected result, wanted %+v, got %+v", peers, fixture)
	}
}
