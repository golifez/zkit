package service

import (
	daws "github.com/golifez/zkit/internal/service/aws"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var SvcProviderSet = wire.NewSet(NewAuthService, daws.NewIamService)
