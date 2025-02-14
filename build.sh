go build -o minicalc .
strip minicalc
upx -q --best --lzma minicalc
