package pmapi

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
	"github.com/juju/errors"
)

func (pmapi *PMApi) DeleteOneUser(c *gin.Context) {
	/*新建一个用户实例*/
	var tmpusr proxysql.Users

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []proxysql.Users{})
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

		if pmapi.ApiErr = c.Bind(&tmpusr); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}
		log.Print("pmapi->DeleteOneUser->DeleteOneUser tmpusr", tmpusr)

		pmapi.ApiErr = tmpusr.DeleteOneUser(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			log.Print("pmapi->DeleteOneUser->DeleteOneUser Failed", pmapi.ApiErr)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

func (pmapi *PMApi) CreateOneUser(c *gin.Context) {
	/*新建一个用户实例*/
	var tmpusr proxysql.Users

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
				if pmapi.ApiErr = c.Bind(&tmpusr); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
				} else {

					pmapi.ApiErr = tmpusr.AddOneUser(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": tmpusr.Username + " Create Successed!"})
					}

				}
			}

		}

	}
}

func (pmapi *PMApi) ListAllUsers(c *gin.Context) {

	var aryusr []proxysql.Users

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

			log.Printf("GET api/v1/users?hostname=%s&port=%s&adminuser=%s&adminpass=%s&limit=%d&page=%d", hostname, port, username, password, limit, page)

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
						//Execute Query.
						aryusr, pmapi.ApiErr = proxysql.FindAllUserInfo(pmapi.Apidb, limit, skip)
						if pmapi.ApiErr != nil {
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						} else {

							// return success.
							c.JSON(http.StatusOK, aryusr)
						}

					}

				}

			}

		}
	}

}

/*更新用户信息的patch方法*/
func (pmapi *PMApi) UpdateOneUser(c *gin.Context) {

	/*新建一个用户实例*/
	var tmpusr proxysql.Users

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
				if pmapi.ApiErr = c.Bind(&tmpusr); pmapi.ApiErr != nil {
					c.JSON(http.StatusExpectationFailed, errors.ErrorStack(pmapi.ApiErr))
				} else {
					pmapi.ApiErr = tmpusr.UpdateOneUserInfo(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": tmpusr.Username + " Update Successed!"})
					}

				}

			}

		}

	}
}
