package validator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-library/lib/presenter"
	"strconv"
)

func SetDefaultQueryParamNumber(ctx *gin.Context, key string, dval string) (int, error) {
	p := ctx.Query(key)
	if p == "" {
		p = dval
	}

	param, err := strconv.Atoi(p)
	if err != nil {
		return 0, fmt.Errorf("'%s' %s", key, presenter.RESPONSE_ERROR_INVALID_PARAM_NUMBER)
	}

	return param, nil
}
