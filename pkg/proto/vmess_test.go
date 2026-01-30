package proto_test

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/drawmoon/xrc/pkg/proto"
)

func TestParseVMessURI_BasicParsing(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		want    *proto.VMess
		wantErr bool
	}{
		{
			name: "basic vmess uri",
			uri: func() string {
				data := map[string]any{
					"v":    "2",
					"ps":   "test-node",
					"add":  "example.com",
					"port": 443,
					"id":   "550e8400-e29b-41d4-a716-446655440000",
				}
				b, _ := json.Marshal(data)
				return "vmess://" + base64.StdEncoding.EncodeToString(b)
			}(),
			want: &proto.VMess{
				Version: "2",
				Name:    "test-node",
				Address: "example.com",
				Port:    443,
				UUID:    "550e8400-e29b-41d4-a716-446655440000",
				Cipher:  "auto", // default cipher
			},
			wantErr: false,
		},
		{
			name: "vmess uri with IP address",
			uri: func() string {
				data := map[string]any{
					"v":    "2",
					"add":  "192.168.1.1",
					"port": 8443,
					"id":   "550e8400-e29b-41d4-a716-446655440000",
				}
				b, _ := json.Marshal(data)
				return "vmess://" + base64.StdEncoding.EncodeToString(b)
			}(),
			want: &proto.VMess{
				Version: "2",
				Address: "192.168.1.1",
				Port:    8443,
				UUID:    "550e8400-e29b-41d4-a716-446655440000",
				Cipher:  "auto",
			},
			wantErr: false,
		},
		{
			name: "vmess uri without padding",
			uri: func() string {
				data := map[string]any{
					"v":    "2",
					"add":  "example.com",
					"port": 443,
					"id":   "550e8400-e29b-41d4-a716-446655440000",
				}
				b, _ := json.Marshal(data)
				encoded := base64.StdEncoding.EncodeToString(b)
				// Remove padding
				return "vmess://" + encoded[:len(encoded)-len(encoded)%4]
			}(),
			want: &proto.VMess{
				Version: "2",
				Address: "example.com",
				Port:    443,
				UUID:    "550e8400-e29b-41d4-a716-446655440000",
				Cipher:  "auto",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseVMessURI(tt.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseVMessURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !vmessEqual(got, tt.want) {
				t.Errorf("ParseVMessURI() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestParseVMessURI_AllFields(t *testing.T) {
	uri := func() string {
		data := map[string]any{
			"v":                "2",
			"ps":               "full-config",
			"add":              "example.com",
			"port":             443,
			"id":               "550e8400-e29b-41d4-a716-446655440000",
			"aid":              32,
			"scy":              "aes-128-gcm",
			"net":              "ws",
			"type":             "http",
			"host":             "example.com",
			"path":             "/vmess",
			"tls":              "tls",
			"sni":              "example.com",
			"fp":               "chrome",
			"alpn":             "h2",
			"skip-cert-verify": true,
		}
		b, _ := json.Marshal(data)
		return "vmess://" + base64.StdEncoding.EncodeToString(b)
	}()

	got, err := proto.ParseVMessURI(uri)
	if err != nil {
		t.Fatalf("ParseVMessURI() unexpected error: %v", err)
	}

	if got.Version != "2" {
		t.Errorf("Version = %v, want 2", got.Version)
	}
	if got.Name != "full-config" {
		t.Errorf("Name = %v, want full-config", got.Name)
	}
	if got.Address != "example.com" {
		t.Errorf("Address = %v, want example.com", got.Address)
	}
	if got.Port != 443 {
		t.Errorf("Port = %v, want 443", got.Port)
	}
	if got.UUID != "550e8400-e29b-41d4-a716-446655440000" {
		t.Errorf("UUID = %v, want 550e8400-e29b-41d4-a716-446655440000", got.UUID)
	}
	if got.AlterID != 32 {
		t.Errorf("AlterID = %v, want 32", got.AlterID)
	}
	if got.Cipher != "aes-128-gcm" {
		t.Errorf("Cipher = %v, want aes-128-gcm", got.Cipher)
	}
	if got.Network != "ws" {
		t.Errorf("Network = %v, want ws", got.Network)
	}
	if got.Type != "http" {
		t.Errorf("Type = %v, want http", got.Type)
	}
	if got.Host != "example.com" {
		t.Errorf("Host = %v, want example.com", got.Host)
	}
	if got.Path != "/vmess" {
		t.Errorf("Path = %v, want /vmess", got.Path)
	}
	if got.TLS != "tls" {
		t.Errorf("TLS = %v, want tls", got.TLS)
	}
	if got.SNI != "example.com" {
		t.Errorf("SNI = %v, want example.com", got.SNI)
	}
	if got.Fingerprint != "chrome" {
		t.Errorf("Fingerprint = %v, want chrome", got.Fingerprint)
	}
	if got.ALPN != "h2" {
		t.Errorf("ALPN = %v, want h2", got.ALPN)
	}
	if got.SkipCertVerify != true {
		t.Errorf("SkipCertVerify = %v, want true", got.SkipCertVerify)
	}
}

func TestParseVMessURI_CipherField(t *testing.T) {
	tests := []struct {
		name       string
		data       map[string]any
		wantCipher string
	}{
		{
			name: "cipher from scy field",
			data: map[string]any{
				"v":    "2",
				"add":  "example.com",
				"port": 443,
				"id":   "550e8400-e29b-41d4-a716-446655440000",
				"scy":  "aes-256-gcm",
			},
			wantCipher: "aes-256-gcm",
		},
		{
			name: "cipher from security field (legacy)",
			data: map[string]any{
				"v":        "2",
				"add":      "example.com",
				"port":     443,
				"id":       "550e8400-e29b-41d4-a716-446655440000",
				"security": "chacha20-poly1305",
			},
			wantCipher: "chacha20-poly1305",
		},
		{
			name: "scy takes precedence over security",
			data: map[string]any{
				"v":        "2",
				"add":      "example.com",
				"port":     443,
				"id":       "550e8400-e29b-41d4-a716-446655440000",
				"scy":      "aes-128-gcm",
				"security": "chacha20-poly1305",
			},
			wantCipher: "aes-128-gcm",
		},
		{
			name: "default to auto when no cipher field",
			data: map[string]any{
				"v":    "2",
				"add":  "example.com",
				"port": 443,
				"id":   "550e8400-e29b-41d4-a716-446655440000",
			},
			wantCipher: "auto",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.data)
			uri := "vmess://" + base64.StdEncoding.EncodeToString(b)

			got, err := proto.ParseVMessURI(uri)
			if err != nil {
				t.Fatalf("ParseVMessURI() unexpected error: %v", err)
			}
			if got.Cipher != tt.wantCipher {
				t.Errorf("Cipher = %v, want %v", got.Cipher, tt.wantCipher)
			}
		})
	}
}

func TestParseVMessURI_NetworkTypes(t *testing.T) {
	networks := []string{"tcp", "ws", "grpc", "http", "quic"}

	for _, net := range networks {
		t.Run("network="+net, func(t *testing.T) {
			data := map[string]any{
				"v":    "2",
				"add":  "example.com",
				"port": 443,
				"id":   "550e8400-e29b-41d4-a716-446655440000",
				"net":  net,
			}
			b, _ := json.Marshal(data)
			uri := "vmess://" + base64.StdEncoding.EncodeToString(b)

			got, err := proto.ParseVMessURI(uri)
			if err != nil {
				t.Fatalf("ParseVMessURI() unexpected error: %v", err)
			}
			if got.Network != net {
				t.Errorf("Network = %v, want %v", got.Network, net)
			}
		})
	}
}

func TestParseVMessURI_InvalidInput(t *testing.T) {
	tests := []struct {
		name   string
		uri    string
		errMsg string
	}{
		{
			name:   "invalid prefix",
			uri:    "vless://invalid",
			errMsg: "invalid vmess prefix",
		},
		{
			name:   "missing prefix",
			uri:    "invalid",
			errMsg: "invalid vmess prefix",
		},
		{
			name:   "invalid base64",
			uri:    "vmess://!!!invalid!!!",
			errMsg: "illegal base64 data",
		},
		{
			name:   "invalid json",
			uri:    "vmess://" + base64.StdEncoding.EncodeToString([]byte("not json")),
			errMsg: "json unmarshal failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseVMessURI(tt.uri)
			if err == nil {
				t.Errorf("ParseVMessURI() expected error, got nil with result %+v", got)
				return
			}
			if !contains(err.Error(), tt.errMsg) {
				t.Errorf("ParseVMessURI() error = %v, want containing %v", err, tt.errMsg)
			}
		})
	}
}

func TestParseVMessURI_EmptyFields(t *testing.T) {
	data := map[string]any{
		"v":    "2",
		"add":  "example.com",
		"port": 443,
		"id":   "550e8400-e29b-41d4-a716-446655440000",
		"ps":   "",
		"net":  "",
		"sni":  "",
	}
	b, _ := json.Marshal(data)
	uri := "vmess://" + base64.StdEncoding.EncodeToString(b)

	got, err := proto.ParseVMessURI(uri)
	if err != nil {
		t.Fatalf("ParseVMessURI() unexpected error: %v", err)
	}

	if got.Version != "2" {
		t.Errorf("Version = %v, want 2", got.Version)
	}
	if got.Address != "example.com" {
		t.Errorf("Address = %v, want example.com", got.Address)
	}
	if got.Port != 443 {
		t.Errorf("Port = %v, want 443", got.Port)
	}
	if got.Name != "" {
		t.Errorf("Name = %v, want empty string", got.Name)
	}
	if got.Network != "" {
		t.Errorf("Network = %v, want empty string", got.Network)
	}
}

func vmessEqual(a, b *proto.VMess) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Version == b.Version &&
		a.Name == b.Name &&
		a.Address == b.Address &&
		a.Port == b.Port &&
		a.UUID == b.UUID &&
		a.AlterID == b.AlterID &&
		a.Cipher == b.Cipher &&
		a.Network == b.Network &&
		a.Type == b.Type &&
		a.Host == b.Host &&
		a.Path == b.Path &&
		a.TLS == b.TLS &&
		a.SNI == b.SNI &&
		a.Fingerprint == b.Fingerprint &&
		a.ALPN == b.ALPN &&
		a.SkipCertVerify == b.SkipCertVerify
}
