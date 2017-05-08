root_templates: ../../templates/js.qtpl ../../templates/message.qtpl
	qtc --dir ../../templates
templates: root_templates
	qtc --dir templates
build: templates
	go build -v
dev: build
	./shop
