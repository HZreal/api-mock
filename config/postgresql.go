package config

import "fmt"

type PostgresqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (m *PostgresqlConfig) GetDsn() string {
	// host=127.0.0.1 user=postgres password=szkj1234567890 dbname=eeds port=55433 sslmode=disable TimeZone=Asia/Shanghai

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", m.Host, m.Port, m.Username, m.Dbname)

}
