clear;
docker network inspect service_network >/dev/null 2>&1 || docker network create service_network;
docker network inspect db_network >/dev/null 2>&1 || docker network create db_network;
docker network inspect public_network >/dev/null 2>&1 || docker network create public_network;
docker-compose up --build -d;
