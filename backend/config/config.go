package config

type Config struct {
	RDBConfig RDBConfig
	WebConfig WebConfig
	AppConfig AppConfig
}

type RDBConfig struct {
	DBName, User, Password string
}

type WebConfig struct {
	ListenPort string
}

type AppConfig struct {
	// TODO:
}
