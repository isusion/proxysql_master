curl --request DELETE \
  --url 'http://localhost:3333/api/v1/queryrules?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: f958e37f-3921-c16d-d645-5f0c3c383cb0' \
  --data '{"username":"dev","destination_hostgroup":1,"rule_id":1}'