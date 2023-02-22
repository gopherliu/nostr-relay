package nostr_relay

import (
	"github.com/nbd-wtf/go-nostr"

	"github.com/gopherliu/nostr-relay/storage"
)

type Relay interface {
	AcceptEvent(*nostr.Event) bool
	Name() string
	Storage() storage.Storage
}
