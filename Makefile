services-up:
	@echo "Running Service = ${SERVICE}"
	docker-compose -f docker-compose.yaml up --build ${SERVICE}
services-stop:
	docker compose stop
services-down:
	docker compose docker-compose.yaml down
init:
	go mod vendor
	echo "Project initiated"
run:
	go run application/rest/*.go
rest-dev:
	air -c .air.toml -- -h
gen-proto:
	protoc -I proto proto/*.proto --gofast_out=plugins=grpc:proto
docker:
	docker build -t tes/bist .
	docker run -d -p 3381:3381 --name tes_bist tes/bist