package indexController

import (
	"github.com/kataras/iris/context"
	"github.com/pipa/hive-api/app/models/ip2loc"
	"github.com/pipa/hive-api/app/utils/jsend"
)

// Index is the homepage handler
func Index(c context.Context) {
	ip := c.Params().Get("ip")
	result := ip2loc.GetIP(ip)

	c.JSON(jsend.Success(result))
}
