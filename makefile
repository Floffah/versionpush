winbuild:
	go build -o bin/versionpush.exe src/main.go

windebug:
	go build -o bin/versionpush.exe -x src/main.go

wintest:
	go build -o bin/versionpush.exe src/main.go
	echo "Running"
	"./bin/versionpush" -builder=maven

build:
	go build -o bin/versionpush src/main.go

debug:
	go build -o bin/versionpush -x src/main.go

test:
	go build -o bin/versionpush src/main.go
	echo "Running"
	"./bin/versionpush" -builder=maven