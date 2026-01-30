package xray

// Outbound represents an outbound connection configuration.
type Outbound struct {
	Tag            string          `json:"tag,omitempty"`            // optional tag for the outbound
	Protocol       string          `json:"protocol"`                 // e.g., "vless", "vmess", "trojan"
	Settings       *VLessSettings  `json:"settings,omitempty"`       // used if Protocol is "vless"
	StreamSettings *StreamSettings `json:"streamSettings,omitempty"` // stream settings for the outbound
}
