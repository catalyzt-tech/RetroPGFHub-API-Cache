docker buildx build --platform linux/arm64 -t go-cache:v-5 --load .

docker tag go-cache:v-5 tgrziminiar/go-cache:v-5

docker push tgrziminiar/go-cache:v-5