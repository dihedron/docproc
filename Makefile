default: build

build:
	@go build

clean:
	@rm -rf ginkgo release/

release: darwin freebsd linux openbsd solaris windows

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./release/darwin_amd64/ginkgo

freebsd:
	GOOS=freebsd GOARCH=386 go build -o ./release/freebsd_386/ginkgo
	GOOS=freebsd GOARCH=amd64 go build -o ./release/freebsd_amd64/ginkgo
	GOOS=freebsd GOARCH=arm go build -o ./release/freebsd_arm/ginkgo
	
linux:
	GOOS=linux GOARCH=386 go build -o ./release/linux_386/ginkgo
	GOOS=linux GOARCH=amd64 go build -o ./release/linux_amd64/ginkgo
	GOOS=linux GOARCH=arm go build -o ./release/linux_arm/ginkgo
	
openbsd:
	GOOS=openbsd GOARCH=386 go build -o ./release/openbsd_386/ginkgo
	GOOS=openbsd GOARCH=amd64 go build -o ./release/openbsd_amd64/ginkgo

solaris:
	GOOS=solaris GOARCH=amd64 go build -o ./release/solaris_amd64/ginkgo

windows:	
	GOOS=windows GOARCH=386 go build -o ./release/windows_386/ginkgo.exe
	GOOS=windows GOARCH=amd64 go build -o ./release/windows_amd64/ginkgo.exe
