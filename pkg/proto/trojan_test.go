package proto_test

import (
	"testing"

	"github.com/drawmoon/xrc/pkg/proto"
)

func TestParseTrojanURI_BasicParsing(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		want    *proto.Trojan
		wantErr bool
	}{
		{
			name:    "basic trojan uri",
			uri:     "trojan://password@example.com:443",
			want:    &proto.Trojan{Password: "password", Address: "example.com", Port: 443},
			wantErr: false,
		},
		{
			name:    "trojan uri with node name",
			uri:     "trojan://mypass@example.com:8443#my-node",
			want:    &proto.Trojan{Name: "my-node", Password: "mypass", Address: "example.com", Port: 8443},
			wantErr: false,
		},
		{
			name:    "trojan uri with URL encoded node name",
			uri:     "trojan://password@example.com:443#%E9%A6%99%E6%B8%AF",
			want:    &proto.Trojan{Name: "香港", Password: "password", Address: "example.com", Port: 443},
			wantErr: false,
		},
		{
			name:    "trojan uri with IP address",
			uri:     "trojan://pass@192.168.1.1:443",
			want:    &proto.Trojan{Password: "pass", Address: "192.168.1.1", Port: 443},
			wantErr: false,
		},
		{
			name:    "trojan uri without port defaults to 443",
			uri:     "trojan://password@example.com",
			want:    &proto.Trojan{Password: "password", Address: "example.com", Port: 443},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseTrojanURI(tt.uri)
			if err != nil {
				t.Fatalf("ParseTrojanURI() unexpected error: %v", err)
			}
			if !trojanEqual(got, tt.want) {
				t.Errorf("ParseTrojanURI() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestParseTrojanURI_QueryParameters(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want *proto.Trojan
	}{
		{
			name: "trojan uri with all query parameters",
			uri:  "trojan://pass@example.com:443?security=tls&sni=example.com&alpn=h2&fp=chrome&pbk=publickey&sid=shortid&net=ws&type=http&host=example.com&path=/proxy#node",
			want: &proto.Trojan{
				Name:          "node",
				Password:      "pass",
				Address:       "example.com",
				Port:          443,
				Security:      "tls",
				SNI:           "example.com",
				ALPN:          "h2",
				FP:            "chrome",
				PublicKey:     "publickey",
				ShortID:       "shortid",
				Network:       "ws",
				Type:          "http",
				Host:          "example.com",
				Path:          "/proxy",
				AllowInsecure: false,
			},
		},
		{
			name: "trojan uri with security=reality",
			uri:  "trojan://pass@example.com:443?security=reality&pbk=key&sid=id",
			want: &proto.Trojan{
				Password:  "pass",
				Address:   "example.com",
				Port:      443,
				Security:  "reality",
				PublicKey: "key",
				ShortID:   "id",
			},
		},
		{
			name: "trojan uri with empty query parameters",
			uri:  "trojan://pass@example.com:443?security=&alpn=",
			want: &proto.Trojan{
				Password: "pass",
				Address:  "example.com",
				Port:     443,
				Security: "",
				ALPN:     "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseTrojanURI(tt.uri)
			if err != nil {
				t.Fatalf("ParseTrojanURI() unexpected error: %v", err)
			}
			if !trojanEqual(got, tt.want) {
				t.Errorf("ParseTrojanURI() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestParseTrojanURI_AllowInsecure(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		wantFlag bool
	}{
		{
			name:     "allowInsecure=1",
			uri:      "trojan://pass@example.com:443?allowInsecure=1",
			wantFlag: true,
		},
		{
			name:     "allowInsecure=true",
			uri:      "trojan://pass@example.com:443?allowInsecure=true",
			wantFlag: true,
		},
		{
			name:     "allowInsecure=True (case insensitive)",
			uri:      "trojan://pass@example.com:443?allowInsecure=True",
			wantFlag: true,
		},
		{
			name:     "allowInsecure not set defaults to false",
			uri:      "trojan://pass@example.com:443",
			wantFlag: false,
		},
		{
			name:     "allowInsecure=false",
			uri:      "trojan://pass@example.com:443?allowInsecure=false",
			wantFlag: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseTrojanURI(tt.uri)
			if err != nil {
				t.Fatalf("ParseTrojanURI() unexpected error: %v", err)
			}
			if got.AllowInsecure != tt.wantFlag {
				t.Errorf("AllowInsecure = %v, want %v", got.AllowInsecure, tt.wantFlag)
			}
		})
	}
}

func TestParseTrojanURI_SpecialCharacters(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		wantPass string
	}{
		{
			name:     "special characters in password",
			uri:      "trojan://p%40ssw0rd@example.com:443",
			wantPass: "p%40ssw0rd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseTrojanURI(tt.uri)
			if err != nil {
				t.Fatalf("ParseTrojanURI() unexpected error: %v", err)
			}
			if got.Password != tt.wantPass {
				t.Errorf("Password = %v, want %v", got.Password, tt.wantPass)
			}
		})
	}
}

func TestParseTrojanURI_CaseInsensitivity(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{
			name: "lowercase trojan scheme",
			uri:  "trojan://pass@example.com:443",
		},
		{
			name: "uppercase TROJAN scheme",
			uri:  "TROJAN://pass@example.com:443",
		},
		{
			name: "mixed case Trojan scheme",
			uri:  "Trojan://pass@example.com:443",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseTrojanURI(tt.uri)
			if err != nil {
				t.Fatalf("ParseTrojanURI() unexpected error: %v", err)
			}
			if got.Password != "pass" || got.Address != "example.com" || got.Port != 443 {
				t.Errorf("ParseTrojanURI() failed to parse correctly with case variation")
			}
		})
	}
}

func TestParseTrojanURI_InvalidInput(t *testing.T) {
	tests := []struct {
		name   string
		uri    string
		errMsg string
	}{
		{
			name:   "invalid scheme",
			uri:    "http://password@example.com:443",
			errMsg: "invalid scheme",
		},
		{
			name:   "invalid url format",
			uri:    "trojan://[invalid",
			errMsg: "parse url failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := proto.ParseTrojanURI(tt.uri)
			if err == nil {
				t.Errorf("ParseTrojanURI() expected error, got nil with result %+v", got)
				return
			}
			if !contains(err.Error(), tt.errMsg) {
				t.Errorf("ParseTrojanURI() error = %v, want containing %v", err, tt.errMsg)
			}
		})
	}
}

func trojanEqual(a, b *proto.Trojan) bool {
	return a.Name == b.Name &&
		a.Address == b.Address &&
		a.Port == b.Port &&
		a.Password == b.Password &&
		a.Security == b.Security &&
		a.SNI == b.SNI &&
		a.ALPN == b.ALPN &&
		a.FP == b.FP &&
		a.AllowInsecure == b.AllowInsecure &&
		a.PublicKey == b.PublicKey &&
		a.ShortID == b.ShortID &&
		a.Network == b.Network &&
		a.Type == b.Type &&
		a.Host == b.Host &&
		a.Path == b.Path
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
