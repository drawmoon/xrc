package xray

// VLessSettings represents the settings specific to the VLess protocol.
type VLessSettings struct {
	VNext []VNext `json:"vnext"` // list of VLess server configurations
}

// VNext represents a VLess server configuration.
type VNext struct {
	Address string      `json:"address"` // server address
	Port    int         `json:"port"`    // server port
	Users   []VLessUser `json:"users"`   // list of users for this server
}

// VLessUser represents a user configuration for VLess.
type VLessUser struct {
	ID         string `json:"id"`             // user ID (UUID)
	Encryption string `json:"encryption"`     // encryption method
	Flow       string `json:"flow,omitempty"` // optional flow setting
}
