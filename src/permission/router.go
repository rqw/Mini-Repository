package permission

import (
	"Mini-Repository/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterRegister() {
	util.Engine.POST("/permission", _queryPermission)
}
func _queryPermission(c *gin.Context) {
	c.JSON(http.StatusOK, util.SUCCESS(GetAllPermission()))
}
