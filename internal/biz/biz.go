package biz

import (
	bizaws "github.com/golifez/zkit/internal/biz/aws"
	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(NewAutherUsecase, bizaws.NewAwsIamUsecase)
