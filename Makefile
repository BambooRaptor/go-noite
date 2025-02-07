clean:
	powershell rm .\tmp\ -Force -Recurse

test:
	@go test -v ./...
