curl --request PUT \
  --url 'http://localhost:3333/api/v1/variables?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 7a4e9e35-37d8-b332-d6bb-e7a2a6106326' \
  --data '{"variable_name":"admin-mysql_ifaces","variable_value":"0.0.0.0:1234"}'