run:
	go run main.go waiter.go table.go order.go
build: 
	go build -o Kitchen 
clean: 
	rm Kitchen
docker:
	if [ -n $(docker image ls | grep zacatov/pr1kitchen)]; then docker rmi zacatov/pr1kitchen;fi
	docker build -t "zacatov/pr1kitchen" .
