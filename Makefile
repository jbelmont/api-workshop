BIN_DIR := "bin/apid"
APID_MAIN := "cmd/apid/main.go"

all: ensure lint test-cover

build:
		docker-compose build

dev:
		docker-compose up mongo redis apid

destroy-dev:
						docker-compose down

ensure:
				go get -u github.com/mattn/goveralls
				go get -u github.com/philwinder/gocoverage
				go get -u github.com/alecthomas/gometalinter
				go get -u github.com/golang/dep/cmd/dep
				go get -u golang.org/x/tools/cmd/cover
				dep ensure

lint:
			gometalinter --install
			gometalinter ./cmd/... ./internal/...

compile: 	cmd/apid/main.go
						CGO_ENABLED=0 go build -i -o ${BIN_DIR} ${APID_MAIN}

test:
			go test ./... -v

test-cover:
						go test ./... -cover

## Travis automation scripts

travis-install:
								go get -u github.com/mattn/goveralls
								go get -u github.com/philwinder/gocoverage
								go get -u github.com/alecthomas/gometalinter
								go get -u github.com/golang/dep/cmd/dep
								go get -u golang.org/x/tools/cmd/cover
								dep ensure

travis-script:
								set -e
								CGO_ENABLED=0 go build -i -o ${BIN_DIR} ${APID_MAIN}
								gometalinter --install
								gometalinter ./cmd/... ./internal/...
								go test ./... -cover
								gocoverage
								goveralls -coverprofile=profile.cov -repotoken=${COVERALLS_TOKEN}
