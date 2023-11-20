package jsonrpc2http

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/tnmproject/go-jsonrpc2"
)

type Client struct {
	http.Client
}

func NewClientRequest(url string, message *jsonrpc2.JsonRpcMessage) (*http.Request, error) {
	raw, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", url, bytes.NewReader(raw))
}

func NewClient() *Client {
	return &Client{
		Client: http.Client{},
	}
}
