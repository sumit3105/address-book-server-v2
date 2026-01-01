package config

type Config struct {
	DBCfg     *DBConfig
	ServerCfg *ServerConfig
	SMTPCfg   *SMTPConfig
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type ServerConfig struct {
	ServerPort string
	JwtSecret  string
	AppURL     string
}

type SMTPConfig struct {
	SMTPUser string
	SMTPPass string
	SMTPHost string
	SMTPPort string
}
