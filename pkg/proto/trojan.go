package proto

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// Trojan represents a trojan:// link payload
type Trojan struct {
	Name          string `json:"ps,omitempty"`            // node name, subscription shownly
	Address       string `json:"add"`                     // domain or IP
	Port          int    `json:"port"`                    // port number
	Password      string `json:"password"`                // password
	Security      string `json:"security,omitempty"`      // e.g. "tls", "reality", "none"
	SNI           string `json:"sni,omitempty"`           // server name indication
	ALPN          string `json:"alpn,omitempty"`          // e.g. "h2", "http/1.1"
	FP            string `json:"fp,omitempty"`            // TLS fingerprint
	AllowInsecure bool   `json:"allowInsecure,omitempty"` // skip cert verify
	PublicKey     string `json:"pbk,omitempty"`           // public key
	ShortID       string `json:"sid,omitempty"`           // short ID
	Network       string `json:"net,omitempty"`           // e.g. "tcp", "ws", "grpc"
	Type          string `json:"type,omitempty"`          // e.g. "none", "http", "srtp", "utp", "wechat-video", "dtls"
	Host          string `json:"host,omitempty"`          // domain for ws, grpc, http/2
	Path          string `json:"path,omitempty"`          // path for ws, grpc, http/2
}

// ParseTrojanURI parses a trojan:// link string and returns a Trojan object.
func ParseTrojanURI(uri string) (*Trojan, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("parse url failed: %w", err)
	}

	if !strings.EqualFold(u.Scheme, "trojan") {
		return nil, fmt.Errorf("invalid scheme: %s", u.Scheme)
	}

	// Node name from fragment
	t := &Trojan{Password: u.User.String(), Name: u.Fragment}

	// Parse Address and Port
	host, portStr, err := net.SplitHostPort(u.Host)
	if err != nil {
		// If no port is specified, SplitHostPort will error, fallback to using Host directly
		t.Address = u.Host
		t.Port = 443 // Default port
	} else {
		t.Address = host
		p, _ := strconv.Atoi(portStr)
		t.Port = p
	}

	query := u.Query()
	t.Security = query.Get("security")
	t.SNI = query.Get("sni")
	t.ALPN = query.Get("alpn")
	t.FP = query.Get("fp")
	t.PublicKey = query.Get("pbk")
	t.ShortID = query.Get("sid")
	t.Network = query.Get("net")
	t.Type = query.Get("type")
	t.Host = query.Get("host")
	t.Path = query.Get("path")

	allowInsecureRaw := query.Get("allowInsecure")
	t.AllowInsecure = allowInsecureRaw == "1" || strings.ToLower(allowInsecureRaw) == "true"

	return t, nil
}
