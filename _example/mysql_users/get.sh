curl --request GET \
  --url 'http://localhost:3333/api/v1/users?hostname=172.18.10.111&port=13306&adminuser=admin&adminpass=admin&limit=2&page=3' \
  --header 'Authorization: Basic Og==' \
  --header 'Cache-Control: no-cache' \
  --header 'Postman-Token: eb55cdb9-0cb7-373c-9783-1e9ee490ebc4' \
  --header 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  --form code=xWnkliVQJURqB2x1 \
  --form grant_type=authorization_code \
  --form redirect_uri=https://www.getpostman.com/oauth2/callback \
  --form client_id=abc123 \
  --form client_secret=ssh-secret