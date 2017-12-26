curl --request DELETE \
  --url 'http://localhost:3333/api/v1/schedulers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 03a5c28d-bc94-8a26-a31e-ae5e8a89becb' \
  --data '{"id":3}'