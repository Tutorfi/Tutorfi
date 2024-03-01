cd ..
docker compose -f ./deployments/docker-compose.yml -f ./deployments/docker-compose.dev.yml up -d
docker exec app go run cmd/dev/main.go --build

