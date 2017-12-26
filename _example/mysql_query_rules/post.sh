curl --request POST \
  --url 'http://localhost:3333/api/v1/queryrules?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: d1acc6a4-1479-cf51-02d2-1b8a0ab95120' \
  --data '{"username":"dev","destination_hostgroup":1}'