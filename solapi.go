package solapi

import (
	"github.com/moon-planet/solapi-go/cash"
	"github.com/moon-planet/solapi-go/messages"
	"github.com/moon-planet/solapi-go/storage"
)

// Client struct
type Client struct {
	Messages messages.Messages
	Storage  storage.Storage
	Cash     cash.Cash
}

// NewClient return a new client
func NewClient() *Client {
	client := Client{}
	return &client
}
