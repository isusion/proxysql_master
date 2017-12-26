curl --request PUT \
  --url 'http://localhost:3333/api/v1/queryrules?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: de7d7f95-fab3-8353-e36f-8ee8a4de1aae' \
  --data '{"username":"dev","destination_hostgroup":100,"rule_id":3}'