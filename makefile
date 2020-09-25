winbuild:
	go build -o bin/versionpush.exe src/main.go

windebug:
	go build -o bin/versionpush.exe -x src/main.go

wintest:
	go build -o bin/versionpush.exe src/main.go
	echo "Running"
	cd test && "../bin/versionpush" -lang=java -pm=maven -ver=1.3.3

linuxbuild:
	go build -o bin/versionpush src/main.go

linuxdebug:
	go build -o bin/versionpush -x src/main.go

linuxtest:
	go build -o bin/versionpush src/main.go
	echo "Running"
	cd test && "../bin/versionpush" -lang=java -pm=maven -ver=1.3.3
