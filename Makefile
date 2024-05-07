.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags="-X 'main.Version=`grep '^number' .version | cut -d '=' -f2`' -X 'main.BuildTime=`date`' -w -s" -buildvcs=false -o ./bin/ ./...
