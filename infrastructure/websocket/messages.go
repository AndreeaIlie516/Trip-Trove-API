package websocket

import "Trip-Trove-API/domain/entities"

type EventUpdateNotification struct {
	Action      string               `json:"action"`
	Destination entities.Destination `json:"destination"`
}

type Message struct {
	Action      string               `json:"action"`
	Destination entities.Destination `json:"destination,omitempty"`
	ID          string               `json:"id,omitempty"`
}
