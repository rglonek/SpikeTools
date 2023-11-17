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

.PHONY: sign
sign:
	cd bin && codesign --verbose --deep --timestamp --force --options runtime --sign '${applesigner}' gendigest-macos-amd64
	cd bin && codesign --verbose --verify gendigest-macos-amd64
	cd bin && codesign --verbose --deep --timestamp --force --options runtime --sign '${applesigner}' gendigest-macos-arm64
	cd bin && codesign --verbose --verify gendigest-macos-arm64

.PHONY: notarize
notarize: notarize-amd64 notarize-arm64

.PHONY: distbuild
distbuild: clean build shrink sign package notarize

.PHONY: notarize-amd64
notarize-amd64:
	rm -f notarize_result_amd64
	xcrun notarytool submit --apple-id ${appleid} --password '${applepw}' --team-id ${appleteamid} -f json --wait --timeout 60m bin/gendigest-macos-amd64.zip |tee notarize_result_amd64
	if [ "$$(cat notarize_result_amd64 |jq -r .status)" != "Accepted" ] ;\
	then \
		echo "ZIP-AMD FAILED TO NOTARIZE" ;\
		cat notarize_result_amd64 ;\
		exit 1 ;\
	else \
		echo "ZIP-AMD NOTARIZE SUCCESS" ;\
	fi

.PHONY: notarize-arm64
notarize-arm64:
	rm -f notarize_result_arm64
	xcrun notarytool submit --apple-id ${appleid} --password '${applepw}' --team-id ${appleteamid} -f json --wait --timeout 60m bin/gendigest-macos-arm64.zip |tee notarize_result_arm64
	if [ "$$(cat notarize_result_arm64 |jq -r .status)" != "Accepted" ] ;\
	then \
		echo "ZIP-ARM FAILED TO NOTARIZE" ;\
		cat notarize_result_arm64 ;\
		exit 1 ;\
	else \
		echo "ZIP-ARM NOTARIZE SUCCESS" ;\
	fi
