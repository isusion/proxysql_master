package pmapi

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
	"github.com/juju/errors"
)

/*与调取器相关的api函数*/
func (pmapi *PMApi) ListAllScheduler(c *gin.Context) {

	var arysch []proxysql.Schedulers

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")
	limit, err := strconv.ParseUint(c.Query("limit"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrorStack(errors.NewBadRequest(err, "limit  must >= 0")))
	} else {
		page, err := strconv.ParseUint(c.Query("page"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.ErrorStack(errors.NewBadRequest(err, "page must > 0")))
		} else {

			if limit == 0 {
				limit = 10
			}

			if page == 0 {
				page = 1
			}

			skip := (page - 1) * limit

			if hostname == "" || hostname == "undefined" || port == "" || port == "undefined" || username == "" || username == "undefined" || password == "" || password == "undefined" {
				c.JSON(http.StatusBadRequest, errors.ErrorStack(errors.New("hostname|port|adminuser|adminpass length is 0 or value is undefined")))
			} else {
				pmapi.PMhost = hostname
				pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
				pmapi.PMuser = username
				pmapi.PMpass = password

				// New connection instance
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
						arysch, pmapi.ApiErr = proxysql.FindAllSchedulerInfo(pmapi.Apidb, limit, skip)
						if pmapi.ApiErr != nil {
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						} else {

							// return success.
							c.JSON(http.StatusOK, arysch)
						}
					}
				}
			}
		}
	}
}

func (pmapi *PMApi) CreateOneScheduler(c *gin.Context) {

	var tmpsch proxysql.Schedulers

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

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		pmapi.PMconn.SetCharset("utf8")
		pmapi.PMconn.SetCollation("utf8_general_ci")
		pmapi.PMconn.MakeDBI()

		if pmapi.ApiErr != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
		} else {
			pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
			if pmapi.ApiErr != nil {
				c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
			} else {

				if pmapi.ApiErr = c.Bind(&tmpsch); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
				} else {
					newsch, _ := proxysql.NewSch(tmpsch.FileName, tmpsch.IntervalMs)
					pmapi.ApiErr = newsch.AddOneScheduler(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
					} else {
						c.JSON(http.StatusCreated, gin.H{"exit": "0", "messages": tmpsch.FileName + "-" + strconv.Itoa(int(tmpsch.IntervalMs)) + " Create Successed!"})
					}

				}

			}
		}
	}
}

func (pmapi *PMApi) DeleteOneScheduler(c *gin.Context) {
	var tmpsch proxysql.Schedulers

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

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		pmapi.PMconn.SetCharset("utf8")
		pmapi.PMconn.SetCollation("utf8_general_ci")
		pmapi.PMconn.MakeDBI()

		if pmapi.ApiErr != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
		} else {
			pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
			if pmapi.ApiErr != nil {
				c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
			} else {

				if pmapi.ApiErr = c.Bind(&tmpsch); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
				} else {
					newsch, _ := proxysql.NewSch(tmpsch.FileName, tmpsch.IntervalMs)
					newsch.SetSchedulerId(tmpsch.Id)
					pmapi.ApiErr = newsch.DeleteOneScheduler(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsNotFound(pmapi.ApiErr):
							c.JSON(http.StatusNotFound, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": strconv.Itoa(int(tmpsch.Id)) + " Delete Successed!"})
					}

				}

			}
		}
	}
}

func (pmapi *PMApi) UpdateOneScheduler(c *gin.Context) {
	var tmpsch proxysql.Schedulers

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	log.Printf("GET api/v1/users?hostname=%s&port=%s&adminuser=%s&adminpass=%s", hostname, port, username, password)

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
				// get args.

				if pmapi.ApiErr = c.Bind(&tmpsch); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
				} else {
					newsch, _ := proxysql.NewSch(tmpsch.FileName, tmpsch.IntervalMs)
					newsch.SetSchedulerId(tmpsch.Id)
					pmapi.ApiErr = newsch.UpdateOneSchedulerInfo(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsNotFound(pmapi.ApiErr):
							c.JSON(http.StatusNotImplemented, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": strconv.Itoa(int(tmpsch.Id)) + " Update Successed!"})
					}

				}
			}
		}
	}
}
