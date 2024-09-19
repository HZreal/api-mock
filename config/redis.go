package config

type RedisConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"DB"`
}

func (r *RedisConfig) GetAddr() string {
	return r.Host + ":" + r.Port
}
