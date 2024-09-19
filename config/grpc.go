package config

type GRPCConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (g *GRPCConfig) GetAddr() string {
	return g.Host + ":" + g.Port
}
