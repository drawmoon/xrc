package proto

// VLess represents the VLESS protocol configuration.
type VLess struct {
	Version    string `json:"v,omitempty"`          // usually "2"
	Name       string `json:"ps,omitempty"`         // node name, subscription shownly
	Address    string `json:"add"`                  // domain or IP
	Port       int    `json:"port"`                 // port number
	UUID       string `json:"id"`                   // user ID
	Encryption string `json:"encryption,omitempty"` // always "none"
	Flow       string `json:"flow,omitempty"`       // e.g. "xtls-rprx-vision"
	Security   string `json:"security,omitempty"`   // e.g. "tls", "reality", "none"
	SNI        string `json:"sni,omitempty"`        // server name indication
	ALPN       string `json:"alpn,omitempty"`       // e.g. "h2", "http/1.1"
	FP         string `json:"fp,omitempty"`         // TLS fingerprint
	PublicKey  string `json:"pbk,omitempty"`        // public key
	ShortID    string `json:"sid,omitempty"`        // short ID
	Network    string `json:"net,omitempty"`        // e.g. "tcp", "ws", "grpc"
	Type       string `json:"type,omitempty"`       // e.g. "none", "http", "srtp", "utp", "wechat-video", "dtls"
	Host       string `json:"host,omitempty"`       // domain for ws, grpc, http/2
	Path       string `json:"path,omitempty"`       // path for ws, grpc, http/2
}
