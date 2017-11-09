package plain_from_csv

//go:generate genfront plain --output with_logger.gen.go --template ../../embedded_files/di_logger.gf
type DoStuff struct{}

func (d *DoStuff) Run() {
	//	log.Critical("DoStuff.Run()")
}
