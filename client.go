package jpush

import (
	"errors"
	"github.com/zwczou/jpush"
)

var defaultClient *Client

func Push(payload *jpush.Payload) (string, error) {
	if defaultClient != nil {
		return "", errors.New("no client found")
	}
	return defaultClient.Push(payload)
}

func New(config Config) *Client {
	return &Client{
		jpush.New(config.Key, config.Secret),
	}
}

// Client see https://github.com/zwczou/jpush
type Client struct {
	*jpush.JpushClient
}
