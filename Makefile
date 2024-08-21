.PHONY: 

run:
	@CGO_ENABLED=0 go build -o bin/app_prod .
	@echo "compiled you application with all its assets to a single binary => bin/app_prod"
	./bin/app_prod

# docker buildx build --platform linux/arm64 -t go-cache:latest --load .
# start:
# 	go run .

