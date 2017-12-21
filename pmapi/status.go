package pmapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
)

//查询出ProxySQL状态信息
func (pmapi *PMApi) ListPStatus(c *gin.Context) {
	ps := new(proxysql.PsStatus)

	c.JSON(http.StatusOK, ps.GetProxySqlStatus(pmapi.Apidb))
}
