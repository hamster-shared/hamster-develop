web:
	rm -rf pkg/controller/dist
	cd frontend && yarn && yarn build

build: web
	go mod tidy
	go build -o aline

macos:
	go mod tidy
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -o aline

linux:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o aline

windows:
	go mod tidy
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o aline.exe


deploy: linux
	scp ./aline ubuntu@ec2-34-232-105-81.compute-1.amazonaws.com:/home/ubuntu/
	ssh ubuntu@ec2-34-232-105-81.compute-1.amazonaws.com "sudo mv /home/ubuntu/aline /usr/local/bin"

