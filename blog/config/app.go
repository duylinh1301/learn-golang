package config

var (
	app = NewConfigApp()
)

type configApp struct {
}

func NewConfigApp() *configApp {
	return &configApp{}
}
