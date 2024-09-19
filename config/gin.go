package config

type GinConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

func (g *GinConfig) GetAddr() string {
	return g.Host + ":" + g.Port
}
