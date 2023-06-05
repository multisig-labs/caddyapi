package caddyapi

type Config struct {
	Admin   AdminConfig   `json:"admin"`
	Logging LoggingConfig `json:"logging"`
	Storage StorageConfig `json:"storage"`
	Apps    AppsConfig    `json:"apps"`
}

type AdminConfig struct {
	Disabled      bool           `json:"disabled"`
	Listen        string         `json:"listen"`
	EnforceOrigin bool           `json:"enforce_origin"`
	Origins       []string       `json:"origins"`
	Config        AdminAPIConfig `json:"config"`
	Identity      IdentityConfig `json:"identity"`
	Remote        RemoteConfig   `json:"remote"`
}

type AdminAPIConfig struct {
	Persist bool        `json:"persist"`
	Load    interface{} `json:"load"` // This can be fulfilled by modules, so it's kept as interface{}
}

type IdentityConfig struct {
	Identifiers []string      `json:"identifiers"`
	Issuers     []interface{} `json:"issuers"` // This can be fulfilled by modules, so it's kept as interface{}
}

type RemoteConfig struct {
	Listen        string          `json:"listen"`
	AccessControl []AccessControl `json:"access_control"`
}

type AccessControl struct {
	PublicKeys  []string     `json:"public_keys"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	Paths   []string `json:"paths"`
	Methods []string `json:"methods"`
}

type LoggingConfig struct {
	Sink SinkConfig           `json:"sink"`
	Logs map[string]LogConfig `json:"logs"`
}

type SinkConfig struct {
	Writer interface{} `json:"writer"` // This can be fulfilled by modules, so it's kept as interface{}
}

type LogConfig struct {
	Writer   interface{} `json:"writer"`  // This can be fulfilled by modules, so it's kept as interface{}
	Encoder  interface{} `json:"encoder"` // This can be fulfilled by modules, so it's kept as interface{}
	Level    string      `json:"level"`
	Sampling Sampling    `json:"sampling"`
	Include  []string    `json:"include"`
	Exclude  []string    `json:"exclude"`
}

type Sampling struct {
	Interval   int `json:"interval"`
	First      int `json:"first"`
	Thereafter int `json:"thereafter"`
}

type StorageConfig struct {
	// This can be fulfilled by modules, so it's kept as interface{}
	StorageRaw interface{} `json:"storage"`
}

type AppsConfig struct {
	// This can be fulfilled by modules, so it's kept as interface{}
	AppsRaw interface{} `json:"apps"`
}
