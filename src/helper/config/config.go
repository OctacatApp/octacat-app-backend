package config

type AppConfig struct {
	App  App
	SMTP SMTP
}

type App struct {
	Name   string
	Router struct {
		GQL struct {
			Port string
		}
		WSS struct {
			Port string
		}
	}
}

type SMTP struct {
	Host     string
	Username string
	Password string
	Port     int64
}
