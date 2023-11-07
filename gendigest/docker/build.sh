cd ..
make clean build shrink
cd docker
cp ../bin/gendigest-linux-amd64 ./gendigest-amd64
cp ../bin/gendigest-linux-arm64 ./gendigest-arm64
docker buildx create --use desktop-linux --name multiarch
docker buildx build --push --platform linux/amd64,linux/arm64/v8 -t robertglonek/gendigest:latest .
