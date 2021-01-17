set -ex

if [ ! -d ./bin ]; then
    mkdir bin
fi

go build -o bin github.com/mohanson/http_server/cmd/http_server
