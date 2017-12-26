curl --request GET \
  --url 'http://localhost:3333/api/v1/schedulers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin&limit=5&page=3' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 01b3638a-1695-a152-8594-3a08ad92640d' \
  --data '{"rule_id":1}'