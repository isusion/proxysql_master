package pmapi

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
	"github.com/juju/errors"
)

//查询出所有查询规则
func (pmapi *PMApi) ListAllQueryRules(c *gin.Context) {

	var aryqrs []proxysql.QueryRules

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
				log.Printf(hostname)
				pmapi.PMhost = hostname
				pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
				pmapi.PMuser = username
				pmapi.PMpass = password

				// New connection instance
				pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
				if pmapi.ApiErr != nil {
					c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
				} else {
					// Open Connection.
					pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
					if pmapi.ApiErr != nil {
						c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
					} else {

						aryqrs, pmapi.ApiErr = proxysql.FindAllQr(pmapi.Apidb, limit, skip)
						if pmapi.ApiErr != nil {
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						} else {

							// return success.
							c.JSON(http.StatusOK, aryqrs)
						}
					}
				}
			}
		}
	}
}

func (pmapi *PMApi) CreateOneQueryRules(c *gin.Context) {

	var tmpqr proxysql.QueryRules

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
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
		} else {
			pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
			if pmapi.ApiErr != nil {
				c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
			} else {
				if pmapi.ApiErr = c.Bind(&tmpqr); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
				} else {
					newqr, _ := proxysql.NewQr(tmpqr.Username, tmpqr.Destination_hostgroup)
					pmapi.ApiErr = newqr.AddOneQr(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
					} else {
						c.JSON(http.StatusCreated, gin.H{"exit": "0", "messages": newqr.Username + "-" + strconv.Itoa(int(newqr.Destination_hostgroup)) + " Create Successed!"})
					}

				}

			}
		}
	}
}

func (pmapi *PMApi) DeleteOneQueryRules(c *gin.Context) {

	var tmpqr proxysql.QueryRules

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
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
		} else {
			pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
			if pmapi.ApiErr != nil {
				c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
			} else {

				if pmapi.ApiErr = c.Bind(&tmpqr); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
				} else {
					delqr, _ := proxysql.NewQr(tmpqr.Username, tmpqr.Destination_hostgroup)
					delqr.SetQrRuleid(tmpqr.Rule_id)
					pmapi.ApiErr = delqr.DeleteOneQr(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsNotFound(pmapi.ApiErr):
							c.JSON(http.StatusNotFound, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": strconv.Itoa(int(delqr.Rule_id)) + " Delete Successed!"})
					}

				}

			}
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
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
		} else {
			// Open Connection.
			pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
			if pmapi.ApiErr != nil {
				c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
			} else {
				// get args.

				if pmapi.ApiErr = c.Bind(&tmpqr); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
				} else {
					updateqr, _ := proxysql.NewQr(tmpqr.Username, tmpqr.Destination_hostgroup)
					updateqr.SetQrRuleid(tmpqr.Rule_id)
					pmapi.ApiErr = updateqr.UpdateOneQrInfo(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsNotFound(pmapi.ApiErr):
							c.JSON(http.StatusNotImplemented, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": strconv.Itoa(int(tmpqr.Rule_id)) + " Update Successed!"})
					}

				}

			}
		}
	}
}
