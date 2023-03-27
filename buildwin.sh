go-winres simply --icon images/icon.ico
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc \
    CXX=x86_64-w64-mingw32-g++ \
    go build -buildvcs=false -ldflags="-H=windowsgui -extldflags=-static" \
		-o minicalc.exe .
strip minicalc.exe
