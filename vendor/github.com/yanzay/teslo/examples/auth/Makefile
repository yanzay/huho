root_templates:
	qtc ../../templates
templates: ../../templates/js.qtpl.go ../../templates/message.qtpl.go
	qtc templates
build: templates
	go build
dev: build
	./auth
