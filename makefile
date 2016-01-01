test: .FORCE
	go test genfront

install: .FORCE
	go install genfront

embed: .FORCE
	go-bindata -nocompress -o process/embedded_template.gen.go -pkg process -prefix .files/ .files/*.fm

.FORCE:


