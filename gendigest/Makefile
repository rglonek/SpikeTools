.NOTPARALLEL:

.PHONY: cleanbuild
cleanbuild: clean build shrink package

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: build
build:
	mkdir bin
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o bin/gendigest-linux-amd64
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath -ldflags="-s -w" -o bin/gendigest-linux-arm64
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o bin/gendigest-macos-amd64
	env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -trimpath -ldflags="-s -w" -o bin/gendigest-macos-arm64

.PHONY: shrink
shrink:
	cd bin && upx gendigest-linux-amd64
	cd bin && upx gendigest-linux-arm64

.PHONY: package
package:
	cd bin && mv gendigest-linux-amd64 gendigest && zip gendigest-linux-amd64.zip gendigest && rm gendigest
	cd bin && mv gendigest-linux-arm64 gendigest && zip gendigest-linux-arm64.zip gendigest && rm gendigest
	cd bin && mv gendigest-macos-amd64 gendigest && zip gendigest-macos-amd64.zip gendigest && rm gendigest
	cd bin && mv gendigest-macos-arm64 gendigest && zip gendigest-macos-arm64.zip gendigest && rm gendigest
