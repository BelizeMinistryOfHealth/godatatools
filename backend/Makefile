deployCases:
	gcloud functions deploy CasesByOutbreak --entry-point HandlerCasesByOutbreak --runtime go113 --trigger-http \
--set-env-vars MONGO_DB=go-data --set-env-vars MONGO_URI=mongodb://godata-dev.epi.openstep.bz:27017/go-data \
--region us-east1 --memory 1024MB --service-account epi-belize@epi-belize.iam.gserviceaccount.com \
--allow-unauthenticated

tidy:
	go mod tidy

# ==============================================================================
# Build
build-linux:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/godatatools cmd/server/main.go
	chmod +x bin/godatatools

build-macos:
	export GO111MODULE=on
	env GOOS=darwin go build -o bin/godatatools cmd/server/main.go
	chmod +x bin/godatatools


clean:
	rm -rf ./bin Gopkg.lock
