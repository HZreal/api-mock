package config

type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

func (q *RabbitMQConfig) GetUrl() string {
	// amqp://admin:root123456@localhost:5672/%2Flocal
	return "amqp://" + q.Username + ":" + q.Password + "@" + q.Host + ":" + q.Port + "/%2F" + q.Vhost
}
