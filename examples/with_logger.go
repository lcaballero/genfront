package examples

//go:generate genfront plain --output with_logger.gen.go --template ../embedded_files/di_logger.fm
type DoStuff struct{}

func (d *DoStuff) Run() {
	//	log.Critical("DoStuff.Run()")
}
