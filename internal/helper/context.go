package helper

import (
	"context"

	"github.com/wallissonmarinho/go-gin-boilerplate/internal/constant"
)

func FromHeader(ctx context.Context) map[string]string {
	from := ctx.Value(constant.RequestHeader("request.headers"))
	if from == nil {
		return map[string]string{}
	}

	values, ok := from.(map[string]string)
	if !ok {
		return map[string]string{}
	}

	return values
}
