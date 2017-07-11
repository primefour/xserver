package xserver

type ConfigInterface interface {
	SaveConfig(path string, config interface{}) error
	LoadConfig(path string) (error, interface{})
}

type ConfigFile struct {
	ConfigPath string
}
