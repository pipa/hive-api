package jwtAuth

import (
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/pipa/hive-api/app/utils/jsend"
)

// Serve creates the jwt-middleware and serves it
func Serve(ctx iris.Context) {
	m := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("dHJ1c3RubzE="), nil
		},
		ErrorHandler:  onError,
		SigningMethod: jwt.SigningMethodHS256,
	})

	if err := m.CheckJWT(ctx); err != nil {
		ctx.StopExecution()
		return
	}

	// If everything ok then call next.
	ctx.Next()
}

// onError is the error handler for jwt auths.
//  in here, we will return jsend format when an error occurs
func onError(c iris.Context, err string) {
	c.StatusCode(iris.StatusUnauthorized)
	c.JSON(jsend.ErrorCode(err, iris.StatusUnauthorized))
}
