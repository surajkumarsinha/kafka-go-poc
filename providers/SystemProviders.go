package providers

import (
	infra "github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	infraInterfaces "github.com/surajkumarsinha/kafka-go-poc/infrastructure/interfaces"
)

func SystemProviders() {
	infra.Register(
		func(params ...any) infraInterfaces.IChiRouter {
			return infra.ChiRouter()
		},
	)
}
