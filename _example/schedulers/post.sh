curl --request POST \
  --url 'http://localhost:3333/api/v1/schedulers?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin' \
  --header 'Cache-Control: no-cache' \
  --header 'Content-Type: application/json' \
  --header 'Postman-Token: 92002821-afb0-84ee-87a2-72162699e4ee' \
  --data '{"filename":"/bin/ls","interval_ms":1000}'