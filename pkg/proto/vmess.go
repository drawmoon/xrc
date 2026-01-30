package proto

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// VMess represents a vmess:// link payload (legacy support).
type VMess struct {
	Version        string `json:"v"`
	Name           string `json:"ps"`
	Address        string `json:"add"`
	Port           int    `json:"port"`
	UUID           string `json:"id"`
	AlterID        int    `json:"aid,omitempty"` // legacy, usually 0
	Cipher         string `json:"scy,omitempty"` // e.g. "auto", "aes-128-gcm"
	Network        string `json:"net,omitempty"` // e.g. "tcp", "ws", "grpc"
	Type           string `json:"type,omitempty"`
	Host           string `json:"host,omitempty"`
	Path           string `json:"path,omitempty"`
	TLS            string `json:"tls,omitempty"`
	SNI            string `json:"sni,omitempty"`
	Fingerprint    string `json:"fp,omitempty"`
	ALPN           string `json:"alpn,omitempty"`
	SkipCertVerify bool   `json:"skip-cert-verify,omitempty"`
}

// ParseVMessURI parses a vmess:// link string and returns a VMess object.
func ParseVMessURI(uri string) (*VMess, error) {
	if !strings.HasPrefix(uri, "vmess://") {
		return nil, fmt.Errorf("invalid vmess prefix")
	}
	uri = strings.TrimPrefix(uri, "vmess://")

	// Many clients generate Base64 without "=" padding, standard library decoding will fail
	if i := len(uri) % 4; i != 0 {
		uri += strings.Repeat("=", 4-i)
	}

	b, err := base64.StdEncoding.DecodeString(uri)
	if err != nil {
		return nil, err
	}

	// First unmarshal into a raw struct to handle legacy fields.
	var temp map[string]any
	if err := json.Unmarshal(b, &temp); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	v := &VMess{}
	mapToVMess(temp, v)
	return v, nil
}

func mapToVMess(m map[string]any, v *VMess) {
	getStr := func(key string) string {
		if v, ok := m[key].(string); ok {
			return v
		}
		return ""
	}

	v.Version = getStr("v")
	v.Name = getStr("ps")
	v.Address = getStr("add")
	v.UUID = getStr("id")
	v.Network = getStr("net")
	v.Type = getStr("type")
	v.Host = getStr("host")
	v.Path = getStr("path")
	v.TLS = getStr("tls")
	v.SNI = getStr("sni")
	v.Fingerprint = getStr("fp")
	v.ALPN = getStr("alpn")

	// Handle legacy "security" field.
	cipher := "auto"
	scy := getStr("scy")
	if scy != "" {
		cipher = scy
	} else {
		security := getStr("security")
		if security != "" {
			cipher = security
		}
	}
	v.Cipher = cipher

	if port, ok := m["port"].(float64); ok {
		v.Port = int(port)
	}
	if aid, ok := m["aid"].(float64); ok {
		v.AlterID = int(aid)
	}
	if skipCertVerify, ok := m["skip-cert-verify"].(bool); ok {
		v.SkipCertVerify = skipCertVerify
	}
}

// func VmessAsOutbound(lk *Link, mux int16) conf.OutboundDetourConfig {
// 	c := conf.OutboundDetourConfig{
// 		Tag:      lk.Tag,
// 		Protocol: lk.Protocol,
// 	}

// 	settings := json.RawMessage([]byte(fmt.Sprintf(`{
// 		"vnext": [
// 			{
// 				"address": "%s",
// 				"port": %v,
// 				"users": [{"id": "%s", "level": 0, "security": "%s"}]
// 			}
// 		]
// 	}`, lk.Address, lk.Port, lk.Id, lk.Security)))
// 	c.Settings = &settings

// 	c.StreamSetting = func() *conf.StreamConfig {
// 		network := conf.TransportProtocol(lk.Network)
// 		stream := &conf.StreamConfig{
// 			Network:  &network,
// 			Security: lk.StreamSecurity,
// 		}

// 		switch lk.Network {
// 		case "ws":
// 			stream.WSSettings = &conf.WebSocketConfig{
// 				Path: lk.Path,
// 				Headers: map[string]string{
// 					"host": lk.Host,
// 				},
// 			}
// 		}

// 		if lk.StreamSecurity == "tls" {
// 			stream.TLSSettings = &conf.TLSConfig{
// 				Insecure: lk.AllowInsecure,
// 			}
// 			if len(lk.Host) > 0 {
// 				stream.TLSSettings.ServerName = lk.Host
// 			}
// 		}

// 		return stream
// 	}()

// 	c.MuxSettings = &conf.MuxConfig{}
// 	if mux > 0 {
// 		c.MuxSettings.Enabled = true
// 		c.MuxSettings.Concurrency = mux
// 	}

// 	return c
// }
