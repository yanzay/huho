huho: main.go storage.go state.go templates/page.qtpl.go templates/viewmodels.go
	go build -v -race
templates/page.qtpl.go: templates/page.qtpl
	qtc --dir templates
dev: huho
	./huho --log-level=trace
build:
	go build -v -race
install:
	go install -v
clean:
	rm huho
