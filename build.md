docker buildx build --platform linux/arm64 -t go-cache:v-6 --load .

docker tag go-cache:v-6 tgrziminiar/go-cache:v-6

docker push tgrziminiar/go-cache:v-6