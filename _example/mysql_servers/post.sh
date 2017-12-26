curl --request POST \
  --url 'http://localhost:3333/api/v1/servers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 8f32dd65-3cee-a9f9-01d5-8549b2b30312' \
  --data '{"hostgroup_id":1,"hostname":"dev","port":3306}'