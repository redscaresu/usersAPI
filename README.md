curl -X POST http://localhost:8080/create \
-H "Content-Type: application/json" \
-d '{
  "first_name": "dave",
  "last_name": "smith",
  "email_address": "a@gemail.com"
}'

curl -X POST http://localhost:8080/create \
-H "Content-Type: application/json" \
-d '{
  "first_name": "john",
  "last_name": "Bull",
  "email_address": "foo@gmail.com"
}'


curl -X POST http://localhost:8080/create \
-H "Content-Type: application/json" \
-d '{
  "first_name": "clive",
  "last_name": "adams",
  "email_address": "foo@googlemail.com"
}'

curl http://localhost:8080/listusers

curl -X PATCH http://localhost:8080/updateuser \
-H "Content-Type: application/json" \
-d '{
  "first_name": "clive",
  "last_name": "xxx",
  "email_address": "yesyes@googlemail.com"
}'
