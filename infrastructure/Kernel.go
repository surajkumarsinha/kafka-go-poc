package infrastructure

type IKernel interface {
	loadEnvVars()
	initializeServiceContainer(config map[string]any)
	initializeRouter()
}

type kernel struct{}

func (k *kernel) loadEnvVars() {
	if err := InitEnvLoader().LoadFromFile(".env"); err != nil {
		println(err)
	}
}

func (k *kernel) initializeServiceContainer(config map[string]any) {
	InitializeServiceContainer(config)
}

func (k *kernel) initializeRouter() {
	ChiRouter()
}
