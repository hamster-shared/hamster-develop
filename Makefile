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
	timestamp=$$(date +"%Y%m%d%H%M%S"); \
	docker build -t hamstershare/hamster-develop:$$timestamp . && \
	docker push hamstershare/hamster-develop:$$timestamp

clean:
	rm -rf aline aline-test aline-worker aline-worker-test
