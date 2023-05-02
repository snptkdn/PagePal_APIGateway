docker-compose -f docker-compose.production.yml build
# Goのビルド(go.sumとbin/mainが作成される)
docker-compose -f docker-compose.production.yml run --rm app make build
# Lambdaへのデプロイ(.srverlessが作成される)
docker-compose -f docker-compose.production.yml run --rm app make deploy
