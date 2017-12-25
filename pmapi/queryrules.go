package pmapi

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
)

//查询出所有查询规则
func (pmapi *PMApi) ListAllQueryRules(c *gin.Context) {

	var aryqrs []proxysql.QueryRules

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	if len(hostname) == 0 || hostname == "undefined" {
		c.JSON(http.StatusOK, []proxysql.QueryRules{})
	} else {
		pmapi.PMhost = hostname
		pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
		pmapi.PMuser = username
		pmapi.PMpass = password

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		aryqrs, pmapi.ApiErr = proxysql.FindAllQr(pmapi.Apidb, limit, skip)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}
		c.JSON(http.StatusOK, aryqrs)
	}

}

func (pmapi *PMApi) CreateOneQueryRules(c *gin.Context) {

	var tmpqr proxysql.QueryRules

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []proxysql.QueryRules{})
	} else {
		pmapi.PMhost = hostname
		pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
		pmapi.PMuser = username
		pmapi.PMpass = password

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		if pmapi.ApiErr = c.Bind(&tmpqr); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}
		log.Print("pmapi->CreateOneQr->AddOneQr tmpqr", tmpqr)

		pmapi.ApiErr = tmpqr.AddOneQr(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) DeleteOneQueryRules(c *gin.Context) {

	var tmpqr proxysql.QueryRules

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []proxysql.QueryRules{})
	} else {
		pmapi.PMhost = hostname
		pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
		pmapi.PMuser = username
		pmapi.PMpass = password

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		if pmapi.ApiErr = c.Bind(&tmpqr); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}
		log.Print("pmapi->DeleteOneQr->DeleteOneQr tmpqr", tmpqr)

		pmapi.ApiErr = tmpqr.DeleteOneQr(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

/*更新一个新的查询规则*/
func (pmapi *PMApi) UpdateOneQueryRules(c *gin.Context) {
	var tmpqr proxysql.QueryRules

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []proxysql.QueryRules{})
	} else {
		pmapi.PMhost = hostname
		pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
		pmapi.PMuser = username
		pmapi.PMpass = password

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		if pmapi.ApiErr = c.Bind(&tmpqr); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}
		log.Print("pmapi->UpdateOneQr->UpdateOneQr tmpqr", tmpqr)

		pmapi.ApiErr = tmpqr.UpdateOneQrInfo(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}
