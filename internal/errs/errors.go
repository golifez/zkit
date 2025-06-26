package errs

import (
	v1 "github.com/golifez/zkit/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)
