package xray

// Routing represents the routing configuration for the Xray application.
type Routing struct {
	DomainStrategy string `json:"domainStrategy,omitempty"` // e.g., "AsIs", "IPIfNonMatch", "IPOnDemand"
	Rules          []Rule `json:"rules,omitempty"`          // list of routing rules
}

// Rule represents a single routing rule.
type Rule struct {
	Type        string   `json:"type"`             // e.g., "field", "chinaip", "geosite"
	IP          []string `json:"ip,omitempty"`     // list of IPs or IP ranges
	Domain      []string `json:"domain,omitempty"` // list of domains
	OutboundTag string   `json:"outboundTag"`      // tag of the outbound to route to
}
