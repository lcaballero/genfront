package examples

//go:generate genfront plain --output with_logger.gen.go --template ./di_logger.fm --data-file table:values.csv
type DoStuff struct{}

func (d *DoStuff) Run() {
	//	log.Critical("DoStuff.Run()")
}
