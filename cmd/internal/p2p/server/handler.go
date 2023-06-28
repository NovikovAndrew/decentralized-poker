package server

import "fmt"

type Handler interface {
	HandleMessage(message *Message) error
}

type DefaultHandler struct {
	Version string
}

func NewHandler() Handler {
	return &DefaultHandler{}
}

func (h *DefaultHandler) HandleMessage(message *Message) error {
	fmt.Printf("handling the message from %s, payload %s\n", message.From.String(), string(message.payload))
	return nil
}
