package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Example = &ExampleController{}

type ExampleController struct {
}

func (ec *ExampleController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Pong",
	})
}
