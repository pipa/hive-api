package onErrorCode

import (
	"github.com/kataras/iris"
	"github.com/pipa/hive-api/utils/jsend"
)

// OnAnyErrorCode handles any iris app error
func OnAnyErrorCode(c iris.Context) {
	c.JSON(jsend.ErrorCode(c.Values().GetString("message"), c.GetStatusCode()))
	return
}
