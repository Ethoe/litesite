docker compose up -d --force-recreate --build
docker image prune -f 
docker compose down

set platform linux/arm64