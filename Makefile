default: build

build:
	@go build -o ./bin/mason

optimise: build
	@upx --best ./bin/mason	

clean:
	@rm -rf bin/ 

release: darwin freebsd linux openbsd solaris windows

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/mason

freebsd:
	GOOS=freebsd GOARCH=386 go build -o ./bin/freebsd_386/mason
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/freebsd_amd64/mason
	GOOS=freebsd GOARCH=arm go build -o ./bin/freebsd_arm/mason
	
linux:
	GOOS=linux GOARCH=386 go build -o ./bin/linux_386/mason
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/mason
	GOOS=linux GOARCH=arm go build -o ./bin/linux_arm/mason
	
openbsd:
	GOOS=openbsd GOARCH=386 go build -o ./bin/openbsd_386/mason
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/openbsd_amd64/mason

solaris:
	GOOS=solaris GOARCH=amd64 go build -o ./bin/solaris_amd64/mason

windows:	
	GOOS=windows GOARCH=386 go build -o ./bin/windows_386/mason.exe
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows_amd64/mason.exe
