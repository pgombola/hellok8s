build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags '-s -w' -o hello
	docker build -t gumby/hellok8s .
	rm hello

push:
	docker push gumby/hellok8s