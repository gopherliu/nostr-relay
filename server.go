package nostr_relay

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Server struct {
	addr  string
	relay Relay

	router     *mux.Router
	httpServer *http.Server

	clientsMu sync.Mutex
}
