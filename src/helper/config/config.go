package config

type AppConfig struct {
	App App
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
