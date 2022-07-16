package fsmv

// Config for server
type Config struct {
	Port         int
	NativeScript bool    // Import native Mermaid(Drawing Graph) script for some reason
	Mode         *string // gin.ReleaseMode for default
}

var globalConfig *Config

func GetGlobalConfig() *Config {
	return globalConfig
}

func SetGlobalConfig(config *Config) {
	globalConfig = config
}
