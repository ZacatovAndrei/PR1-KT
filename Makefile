run:
	go run main.go waiter.go table.go order.go
build: 
	go build -o Kitchen 
clean: 
	rm Kitchen
docker:
	docker rmi zacatov/pr1kitchen
	docker build -t "zacatov/pr1kitchen" .
