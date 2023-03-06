web:
	rm -rf pkg/controller/dist
	cd frontend && yarn && yarn build

build: web
	go mod tidy
	go build -o aline

macos:
	go mod tidy
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -ldflags="-s -w" -o aline
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -ldflags="-s -w" -o aline-worker ./bin/aline-worker

linux:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags="-s -w" -o aline
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags="-s -w" -o aline-worker ./bin/aline-worker

linux-test:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags="-s -w" -o aline-test
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags="-s -w" -o aline-worker-test ./bin/aline-worker

windows:
	go mod tidy
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags="-s -w" -o aline.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -ldflags="-s -w" -o aline-worker.exe ./bin/aline-worker

docker: linux-test
	docker build -t hamstershare/hamster-develop:latest .
	docker push hamstershare/hamster-develop:latest


deploy: linux
	scp ./aline ubuntu@ec2-34-232-105-81.compute-1.amazonaws.com:/home/ubuntu/
	ssh ubuntu@ec2-34-232-105-81.compute-1.amazonaws.com "sudo mv /home/ubuntu/aline /usr/local/bin"

