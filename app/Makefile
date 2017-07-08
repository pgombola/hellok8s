build:
	GOOS=linux CGO_ENABLED=0 go build -o hello
	docker build -t gumby/hellok8s .

push:
	docker push gumby/hellok8s