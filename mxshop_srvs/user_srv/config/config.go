package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type SrvConfig struct {
	Ip   string `mapstructure:"ip" json:"ip"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Srv        SrvConfig    `mapstructure:"user-srv" json:"user-srv"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}
