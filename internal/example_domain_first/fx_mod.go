package example_domain_first

import (
	"github.com/neiasit/service-boilerplate/internal/example_domain_first/delivery"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"domain_name_example",
	fx.Invoke(
		delivery.RegisterHandlers,
	),
)
