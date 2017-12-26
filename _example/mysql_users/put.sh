curl --request PUT \
  --url 'http://localhost:3333/api/v1/users?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: f310cb76-5077-2dec-cbfd-a7202f8b8ab7' \
  --data '{"username":"dev3","password":"1"}'