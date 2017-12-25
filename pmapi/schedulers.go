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
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		arysch, err = proxysql.FindAllSchedulerInfo(pmapi.Apidb, limit, skip)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
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
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		if err := c.Bind(&tmpsch); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->AddOneScheduler->AddOneScheduler tmpsch", tmpsch)

		err = tmpsch.AddOneScheduler(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneScheduler->AddOneScheduler Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
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
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		if err := c.Bind(&tmpsch); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->DeleteOneScheduler->DeleteOneScheduler tmpsch", tmpsch)

		err = tmpsch.DeleteOneScheduler(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneScheduler->DeleteOneScheduler Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
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
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		conn, err := proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		pmapi.Apidb, err = conn.OpenConn()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}

		if err := c.Bind(&tmpsch); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->UpdateOneScheduler->UpdateOneScheduler tmpsch", tmpsch)

		err = tmpsch.UpdateOneSchedulerInfo(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneScheduler->UpdateOneScheduler Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}

}
