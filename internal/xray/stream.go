package xray

// StreamSettings represents the stream settings for an outbound connection.
type StreamSettings struct {
	Network         string           `json:"network"`                   // e.g., "tcp", "ws", "http", "quic"
	Security        string           `json:"security,omitempty"`        // e.g., "none", "tls", "reality"
	RealitySettings *RealitySettings `json:"realitySettings,omitempty"` // used if Security is "reality"
}

// RealitySettings represents the settings specific to the Reality security protocol.
type RealitySettings struct {
	ServerName  string `json:"serverName"`            // domain name of the server
	PublicKey   string `json:"publicKey"`             // base64-encoded public key
	ShortID     string `json:"shortId,omitempty"`     // optional short ID
	Fingerprint string `json:"fingerprint,omitempty"` // optional fingerprint type
}
