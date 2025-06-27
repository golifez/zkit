package utils

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewJwt)
