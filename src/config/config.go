package config

type AppConfig struct {
	App App
}

type App struct {
	Router struct {
		GQL struct {
			Port string
		}
		WSS struct {
			Port string
		}
	}
}
