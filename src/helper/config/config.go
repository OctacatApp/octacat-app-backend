package config

type AppConfig struct {
	App        App
	SMTP       SMTP
	DataSource DataSource
}

type App struct {
	Name    string
	Version string
	Router  struct {
		GQL struct {
			Port string
		}
		WSS struct {
			Port                     string
			ReadBufferSize           int64
			WriteBufferSize          int64
			HandshakeTimeoutInMinute int64
			EnableCompression        bool
		}
	}
	Default struct {
		Me string
	}
	JWT struct {
		Secret      string
		ExpInMinute int64
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
	Port     int64
	SSL      bool
}

type DataSource struct {
	PostgreSQL SQLDB
}
