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
	/*delete a users*/
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
					deluser, _ := proxysql.NewUser(tmpusr.Username, tmpusr.Password, tmpusr.DefaultHostgroup, tmpusr.DefaultSchema)
					deluser.SetBackend(1)
					deluser.SetFrontend(1)

					pmapi.ApiErr = deluser.DeleteOneUser(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsNotFound(pmapi.ApiErr):
							c.JSON(http.StatusNotFound, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": tmpusr.Username + " Delete Successed!"})
					}

				}

			}

		}

	}
}

func (pmapi *PMApi) CreateOneUser(c *gin.Context) {
	/*create a user*/
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
					newuser, _ := proxysql.NewUser(tmpusr.Username, tmpusr.Password, tmpusr.DefaultHostgroup, tmpusr.DefaultSchema)
					newuser.SetBackend(1)
					newuser.SetFrontend(1)
					pmapi.ApiErr = newuser.AddOneUser(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsAlreadyExists(pmapi.ApiErr):
							c.JSON(http.StatusFound, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusCreated, gin.H{"exit": "0", "messages": tmpusr.Username + " Create Successed!"})
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

func (pmapi *PMApi) UpdateOneUser(c *gin.Context) {

	/*update a user*/
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
					updateuser, _ := proxysql.NewUser(tmpusr.Username, tmpusr.Password, tmpusr.DefaultHostgroup, tmpusr.DefaultSchema)
					updateuser.SetBackend(1)
					updateuser.SetFrontend(1)
					pmapi.ApiErr = updateuser.UpdateOneUserInfo(pmapi.Apidb)
					if pmapi.ApiErr != nil {
						switch {
						case errors.IsNotFound(pmapi.ApiErr):
							c.JSON(http.StatusNotImplemented, errors.ErrorStack(pmapi.ApiErr))
						default:
							c.JSON(http.StatusInternalServerError, errors.ErrorStack(pmapi.ApiErr))
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"exit": "0", "messages": tmpusr.Username + " Update Successed!"})
					}

				}

			}

		}

	}
}
