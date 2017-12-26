curl --request PUT \
  --url 'http://localhost:3333/api/v1/servers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 6adcbd3f-5264-b903-4503-4a16bfe648ee' \
  --data '{"hostgroup_id":999,"hostname":"192.168.100.111","port":7031,"max_connections":999}'