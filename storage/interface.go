package storage

import (
	"github.com/nbd-wtf/go-nostr"
)

type Storage interface {
	QueryEvents(filter *nostr.Filter) (events []nostr.Event, err error)
	DeleteEvent(id string, pubKey string) error
	SaveEvent(event *nostr.Event) error
}
