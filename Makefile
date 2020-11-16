build:
	go build -o bin/memda ./src/

compile:
	GOOS=linux GOARCH=386 go build -o bin/memda-linux-386 ./src/
	GOOS=windows GOARCH=386 go build -o bin/memda-windows-386 ./src/
	GOOS=darwin GOARCH=amd64 go build -o bin/memda-darwin-amd64 ./src/
	GOOS=linux GOARCH=amd64 go build -o bin/memda-linux-amd64 ./src/
	GOOS=windows GOARCH=amd64 go build -o bin/memda-windows-amd64 ./src/
