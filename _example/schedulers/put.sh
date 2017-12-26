curl --request PUT \
  --url 'http://localhost:3333/api/v1/schedulers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 06b4e5da-6598-a833-5853-7d7472457f93' \
  --data '{"filename":"/bin/ls","interval_ms":1111,"id":3}'