package config

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Timeout  string `yaml:"timeout"`
	MaxConn  int    `yaml:"maxConn"`
	MaxOpen  int    `yaml:"maxOpen"`
}

func (m *MysqlConfig) GetDsn() string {
	//
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local&timeout=" + m.Timeout
}
