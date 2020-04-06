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

func (ec *ExampleController) NotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}

func (ec *ExampleController) OK(c *gin.Context) {
	c.Status(http.StatusOK)
}
