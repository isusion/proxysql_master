package pmapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
	"github.com/juju/errors"
)

//查询出所有变量的内容
func (pmapi *PMApi) ListPsVariables(c *gin.Context) {
	var aryvars []proxysql.Variables

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if hostname == "" || hostname == "undefined" || port == "" || port == "undefined" || username == "" || username == "undefined" || password == "" || password == "undefined" {
		c.JSON(http.StatusBadRequest, errors.ErrorStack(errors.New("hostname|port|adminuser|adminpass length is 0 or value is undefined")))
	} else {

		pmapi.PMhost = hostname
		pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
		pmapi.PMuser = username
		pmapi.PMpass = password

		// New connection instance.
		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		pmapi.PMconn.SetCharset("utf8")
		pmapi.PMconn.SetCollation("utf8_general_ci")
		pmapi.PMconn.MakeDBI()

		if pmapi.ApiErr != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
		} else {
			// Open Connection.
			pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
			if pmapi.ApiErr != nil {
				c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
			} else {

				aryvars, pmapi.ApiErr = proxysql.GetConfig(pmapi.Apidb)
				if pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
				} else {
					c.JSON(http.StatusOK, aryvars)
				}
			}
		}
	}

}

func (pmapi *PMApi) UpdateOneVariables(c *gin.Context) {

	var tmpvars proxysql.Variables

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if hostname == "" || hostname == "undefined" || port == "" || port == "undefined" || username == "" || username == "undefined" || password == "" || password == "undefined" {
		c.JSON(http.StatusBadRequest, errors.ErrorStack(errors.New("hostname|port|adminuser|adminpass length is 0 or value is undefined")))
	} else {

		pmapi.PMhost = hostname
		pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
		pmapi.PMuser = username
		pmapi.PMpass = password

		// New connection instance.
		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		pmapi.PMconn.SetCharset("utf8")
		pmapi.PMconn.SetCollation("utf8_general_ci")
		pmapi.PMconn.MakeDBI()

		if pmapi.ApiErr != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
		} else {
			// Open Connection.
			pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
			if pmapi.ApiErr != nil {
				c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
			} else {

				if pmapi.ApiErr = c.Bind(&tmpvars); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
				} else {
					pmapi.ApiErr = proxysql.UpdateOneConfig(pmapi.Apidb, tmpvars.VariablesName, tmpvars.Value)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsNotImplemented(pmapi.ApiErr):
							c.JSON(http.StatusNotImplemented, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": tmpvars.VariablesName + " Update Successed!"})
					}

				}

			}

		}
	}
}
