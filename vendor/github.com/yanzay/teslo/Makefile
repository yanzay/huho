build: templates
	go build -v
templates: templates/js.qtpl.go templates/message.qtpl.go
templates/js.qtpl.go: templates/js.qtpl templates/main.js
	qtc --dir templates
templates/message.qtpl.go: templates/message.qtpl
	qtc --dir templates
