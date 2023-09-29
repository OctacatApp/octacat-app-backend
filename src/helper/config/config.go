package config

type AppConfig struct {
	App        App
	SMTP       SMTP
	DataSource DataSource
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
	Default struct {
		Me string
	}
}

type SMTP struct {
	Host     string
	Username string
	Password string
	Port     int64
}

type SQLDB struct {
	Driver   string
	Host     string
	Username string
	Password string
	Database string
	Port     string
	SSL      bool
}

type DataSource struct {
	PostgreSQL SQLDB
}
