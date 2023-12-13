package config

type Config struct {
	MysqlHost   string
	MysqlPort   int
	MysqlPasswd string
	MysqlUser   string
}

func GetConfig() Config {
	return Config{
		MysqlHost:   "127.0.0.1",
		MysqlPort:   3306,
		MysqlPasswd: "password",
		MysqlUser:   "root",
	}
}
