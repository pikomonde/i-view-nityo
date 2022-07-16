package model

type Config struct {
	App   AppConfig   `json:"app" mapstructure:"app"`
	MySQL MySQLConfig `json:"mysql" mapstructure:"mysql"`
}

type Env string

const (
	Env_Local      Env = "local"
	Env_Staging    Env = "staging"
	Env_Production Env = "production"
)

type DBType string

const (
	DBType_InMemory DBType = "IN_MEMORY"
	DBType_MySQL    DBType = "MYSQL"
)

type AppConfig struct {
	Env          Env    `json:"env" mapstructure:"env"`
	DBType       DBType `json:"db_type" mapstructure:"db_type"`
	Port         int    `json:"port" mapstructure:"port"`
	ReadTimeout  int    `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout int    `json:"write_timeout" mapstructure:"write_timeout"`
}

type MySQLConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	DBName   string `json:"db_name" mapstructure:"db_name"`
}
