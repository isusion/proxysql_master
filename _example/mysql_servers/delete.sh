curl --request DELETE \
  --url 'http://localhost:3333/api/v1/servers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 9f252ca9-33de-55be-bc54-8ca8b7ac3a6e' \
  --data '{"hostgroup_id":0,"hostname":"dev","port":3306}'