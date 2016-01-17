package main

// Channel is a JSON object that represents a Channel.
type Channel struct {
	ID string `json:"id"`
}

// ListChannelsResponse is a JSON object associated with GET /api/channels.list.
type ListChannelsResponse struct {
	Channels []*Channel `json:"channels"`
}
