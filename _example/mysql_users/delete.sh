curl --request DELETE \
  --url 'http://localhost:3333/api/v1/users?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 656f70dd-aba4-df24-9b9b-d7632fb79779' \
  --data '{"username":"dev3"}'