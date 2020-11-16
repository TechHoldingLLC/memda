build:
	go build -o bin/memda ./src/

compile:
	GOOS=linux GOARCH=386 go build -o bin/memda-linux-386 ./src/
	GOOS=windows GOARCH=386 go build -o bin/memda-windows-386.exe ./src/
	GOOS=darwin GOARCH=amd64 go build -o bin/memda-darwin-amd64 ./src/
	GOOS=linux GOARCH=amd64 go build -o bin/memda-linux-amd64 ./src/
	GOOS=windows GOARCH=amd64 go build -o bin/memda-windows-amd64.exe ./src/

	tar czf bin/memda-linux-386.tar.gz bin/memda-linux-386
	tar czf bin/memda-windows-386.exe.tar.gz bin/memda-windows-386.exe
	tar czf bin/memda-darwin-amd64.tar.gz bin/memda-darwin-amd64
	tar czf bin/memda-linux-amd64.tar.gz bin/memda-linux-amd64
	tar czf bin/memda-windows-amd64.exe.tar.gz bin/memda-windows-amd64.exe

	sha256sum bin/memda-linux-386.tar.gz > bin/memda-linux-386.tar.gz.sha256
	sha256sum bin/memda-windows-386.exe.tar.gz > bin/memda-windows-386.exe.tar.gz.sha256
	sha256sum bin/memda-darwin-amd64.tar.gz > bin/memda-darwin-amd64.tar.gz.sha256
	sha256sum bin/memda-linux-amd64.tar.gz > bin/memda-linux-amd64.tar.gz.sha256
	sha256sum bin/memda-windows-amd64.exe.tar.gz > bin/memda-windows-amd64.exe.tar.gz.sha256
