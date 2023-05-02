# ビルドのやり方
以下コマンドを順番に実行
1. docker compose -f docker-compose.yml build
1. docker compose -f docker-compose.yml up -d
1. docker exec -it pagepal-go sh
1. go mod tidy
1. cd handler
1. air

これで、localhost:3000にAPIが立ち上がります。

