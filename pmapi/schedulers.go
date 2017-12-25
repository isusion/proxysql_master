package pmapi

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
)

/*与调取器相关的api函数*/
func (pmapi *PMApi) ListAllScheduler(c *gin.Context) {

	var arysch []proxysql.Schedulers

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
		c.JSON(http.StatusOK, []proxysql.Schedulers{})
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

		arysch, pmapi.ApiErr = proxysql.FindAllSchedulerInfo(pmapi.Apidb, limit, skip)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}
		c.JSON(http.StatusOK, arysch)
	}
}

func (pmapi *PMApi) CreateOneScheduler(c *gin.Context) {

	var tmpsch proxysql.Schedulers

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 || hostname == "undefined" {
		c.JSON(http.StatusOK, []proxysql.Schedulers{})
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

		if pmapi.ApiErr = c.Bind(&tmpsch); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}
		log.Print("pmapi->AddOneScheduler->AddOneScheduler tmpsch", tmpsch)

		pmapi.ApiErr = tmpsch.AddOneScheduler(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			log.Print("pmapi->CreateOneScheduler->AddOneScheduler Failed", pmapi.ApiErr)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) DeleteOneScheduler(c *gin.Context) {
	var tmpsch proxysql.Schedulers

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 || hostname == "undefined" {
		c.JSON(http.StatusOK, []proxysql.Schedulers{})
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

		if pmapi.ApiErr = c.Bind(&tmpsch); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}
		log.Print("pmapi->DeleteOneScheduler->DeleteOneScheduler tmpsch", tmpsch)

		pmapi.ApiErr = tmpsch.DeleteOneScheduler(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			log.Print("pmapi->CreateOneScheduler->DeleteOneScheduler Failed", pmapi.ApiErr)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) UpdateOneScheduler(c *gin.Context) {
	var tmpsch proxysql.Schedulers

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 || hostname == "undefined" {
		c.JSON(http.StatusOK, []proxysql.Schedulers{})
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

		if pmapi.ApiErr = c.Bind(&tmpsch); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}
		log.Print("pmapi->UpdateOneScheduler->UpdateOneScheduler tmpsch", tmpsch)

		pmapi.ApiErr = tmpsch.UpdateOneSchedulerInfo(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			log.Print("pmapi->CreateOneScheduler->UpdateOneScheduler Failed", pmapi.ApiErr)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}

}
