build:
	go build -o bin/versionpush.exe src/main.go

debug:
	go build -o bin/versionpush.exe -x src/main.go

test:
	go build -o bin/versionpush.exe src/main.go
	echo "Running"
	"./bin/versionpush" -builder=maven