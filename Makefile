huho: main.go storage.go templates/page.qtpl.go templates/viewmodels.go
	go build -v
templates/page.qtpl.go: templates/page.qtpl
	qtc --dir templates
dev: huho
	./huho
build:
	go build -v
clean:
	rm huho
