# set up
``docker compose up -d --build``

# マイグレーション
``RUN go run cmd/migrate/main.go -exec up``

# サーバ起動
``RUN go run cmd/server/main.go``

**API**

# 登録API
``curl localhost:8080/api/v1/users -X POST -H "Content-Type: application/json" -d '{"name": "Ken Tse", "age":31}'``

# 更新API
``curl localhost:8080/api/v1/users/1 -X PATCH -H "Content-Type: application/json" -d '{"name": "Tse Pak Kwan", "age":31}'``

# 削除API
``curl localhost:8080/api/v1/users/1 -X DELETE``

# リストAPI
``curl localhost:8080/api/v1/users``