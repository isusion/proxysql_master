curl --request POST \
  --url 'http://localhost:3333/api/v1/users?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: ff9f89a6-02fc-87a8-d1a3-a953d49adfdf' \
  --data '{"username":"dev3","password":"dev","default_schema":"dev"}'