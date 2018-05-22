package config

type Config struct {
	// The root directory for all data.
	// This should be set in viper so it can unmarshal into this struct
	RootDir string `mapstructure:"home"`
	ChainID string `mapstructure:"chain_id "`
}

// Set the RootDir for all Config structs
func (cfg *Config) SetRoot(root string) *Config {
	cfg.RootDir = root
	return cfg
}

func DefaultConfig() *Config {
	return &Config{}
}
