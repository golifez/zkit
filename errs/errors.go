package errs

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// v1 "github.com/golifez/zkit/api/helloworld/v1"

var (
	ErrDataAlreadyExists = errors.Errorf(200, "数据已存在", "data already exists")
)
