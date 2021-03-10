build:
	cd src; go build -o ../bin/memda .

compile:
	cd src; GOOS=darwin GOARCH=amd64 go build -o ../bin/memda-darwin-amd64 .
	cd src; GOOS=darwin GOARCH=arm64 go build -o ../bin/memda-darwin-arm64 .
	cd src; GOOS=linux GOARCH=amd64 go build -o ../bin/memda-linux-amd64 .
	cd src; GOOS=windows GOARCH=amd64 go build -o ../bin/memda-windows-amd64.exe .

	tar czf bin/memda-darwin-amd64.tar.gz bin/memda-darwin-amd64
	tar czf bin/memda-darwin-arm64.tar.gz bin/memda-darwin-arm64
	tar czf bin/memda-linux-amd64.tar.gz bin/memda-linux-amd64
	tar czf bin/memda-windows-amd64.exe.tar.gz bin/memda-windows-amd64.exe

	sha256sum bin/memda-darwin-amd64.tar.gz > bin/memda-darwin-amd64.tar.gz.sha256
	sha256sum bin/memda-darwin-arm64.tar.gz > bin/memda-darwin-arm64.tar.gz.sha256
	sha256sum bin/memda-linux-amd64.tar.gz > bin/memda-linux-amd64.tar.gz.sha256
	sha256sum bin/memda-windows-amd64.exe.tar.gz > bin/memda-windows-amd64.exe.tar.gz.sha256
