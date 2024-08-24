docker buildx build --platform linux/arm64 -t go-cache:latest --load .

docker tag go-cache:latest ghcr.io/tgrziminiar/go-cache:latest

docker push ghcr.io/tgrziminiar/go-cache:latest