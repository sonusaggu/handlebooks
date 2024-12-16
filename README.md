# handlebooks

curl --location 'http://localhost:8080/graphql' \
--header 'Content-Type: application/json' \
--data '{
  "query": "query { book(id:2) { id title author } }"
}
'


curl --location 'http://localhost:8080/graphql' \
--header 'Content-Type: application/json' \
--data '{
  "query": "mutation { createBook(title: \"The Great Gatsby\", author: \"F. Scott Fitzgerald\") {  title author } }"
}'
