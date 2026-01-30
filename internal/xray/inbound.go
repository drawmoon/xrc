package xray

// Inbound represents an inbound connection configuration.
type Inbound struct {
	Tag      string        `json:"tag,omitempty"`      // optional tag for the inbound
	Listen   string        `json:"listen,omitempty"`   // optional listen address
	Port     int           `json:"port"`               // port number
	Protocol string        `json:"protocol"`           // e.g., "socks", "http", "vmess"
	Settings *SocksSetting `json:"settings,omitempty"` // used if Protocol is "socks"
}

// SocksSetting represents settings specific to the SOCKS protocol.
type SocksSetting struct {
	UDP bool `json:"udp"` // whether UDP is enabled
}
