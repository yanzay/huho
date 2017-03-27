huho: main.go storage.go state.go templates/page.qtpl.go templates/viewmodels.go
	go build -v
templates/page.qtpl.go: templates/page.qtpl
	qtc --dir templates
dev: huho
	./huho --log-level=trace
build:
	go build -v
race:
	go build -v -race
install:
	go install -v
test:
	go test
clean:
	rm huho
