## これは何か

GraphQL をお試し

### 実行

```
dep ensure
go run main
```

### リクエスサンプル

```
# query
curl -X POST -d '{ hello }' http://localhost:8000/graphql
# mutation
curl -X POST -d 'mutation{ hello(message: "aaaa"){message} }' http://localhost:8000/graphql
```
