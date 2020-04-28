export GOPROXY=https://goproxy.io
go build -a -ldflags "-linkmode external -extldflags -static -s -w" -o myblogs