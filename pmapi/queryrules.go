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
		pmapi.PMdb = "information_schema"

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		aryqrs, err = proxysql.FindAllQr(pmapi.Apidb, limit, skip)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
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
		pmapi.PMdb = "information_schema"

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		if err := c.Bind(&tmpqr); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->CreateOneQr->AddOneQr tmpqr", tmpqr)

		err = tmpqr.AddOneQr(pmapi.Apidb)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
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
		pmapi.PMdb = "information_schema"

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		if err := c.Bind(&tmpqr); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->DeleteOneQr->DeleteOneQr tmpqr", tmpqr)

		err = tmpqr.DeleteOneQr(pmapi.Apidb)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
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
		pmapi.PMdb = "information_schema"

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		if err := c.Bind(&tmpqr); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->UpdateOneQr->UpdateOneQr tmpqr", tmpqr)

		err = tmpqr.UpdateOneQrInfo(pmapi.Apidb)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}
