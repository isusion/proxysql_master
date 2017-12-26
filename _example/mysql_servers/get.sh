curl --request GET \
  --url 'http://localhost:3333/api/v1/servers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin&limit=5&page=1' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: d68858c4-2bd7-62bc-b6ed-3ff3564c7414' \
  --data '{"hostgroup_id":0,"hostname":"dev","port":3306}'