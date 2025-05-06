ROOT_DIR=$(realpath ../../)
SERVICE_NAME=news-service

build:
	go build -o $(SERVICE_NAME) ./cmd

run: build
	./$(SERVICE_NAME)

docker-build:
	docker build -f $(ROOT_DIR)/services/news/Dockerfile -t $(SERVICE_NAME) $(ROOT_DIR)

docker-run:
	docker run --rm -it --env-file .env -p 8082:8082 $(SERVICE_NAME)

clean:
	rm -f $(SERVICE_NAME)