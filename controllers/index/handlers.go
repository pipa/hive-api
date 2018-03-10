package indexController

import (
	"github.com/kataras/iris/context"
	"github.com/pipa/hive-api/utils/jsend"
)

// Index is the homepage handler
func Index(c context.Context) {
	result := map[string]string{
		"foo": "bar",
	}

	c.Application().Logger().Debug("here")
	c.JSON(jsend.Success(result))
}
