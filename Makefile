default: build

build:
	@go build -o ./bin/ginkgo

optimise: build
	@upx --best ./bin/ginkgo	

clean:
	@rm -rf bin/ 

release: darwin freebsd linux openbsd solaris windows

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/ginkgo

freebsd:
	GOOS=freebsd GOARCH=386 go build -o ./bin/freebsd_386/ginkgo
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/freebsd_amd64/ginkgo
	GOOS=freebsd GOARCH=arm go build -o ./bin/freebsd_arm/ginkgo
	
linux:
	GOOS=linux GOARCH=386 go build -o ./bin/linux_386/ginkgo
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/ginkgo
	GOOS=linux GOARCH=arm go build -o ./bin/linux_arm/ginkgo
	
openbsd:
	GOOS=openbsd GOARCH=386 go build -o ./bin/openbsd_386/ginkgo
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/openbsd_amd64/ginkgo

solaris:
	GOOS=solaris GOARCH=amd64 go build -o ./bin/solaris_amd64/ginkgo

windows:	
	GOOS=windows GOARCH=386 go build -o ./bin/windows_386/ginkgo.exe
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows_amd64/ginkgo.exe
