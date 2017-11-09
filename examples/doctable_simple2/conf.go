package doctable_simple2

type MonoStat func()
type CountStat func()

//go:generate genfront doctable --line $GOLINE --output conf-options.gen.json
type Conf struct {
	// Comment line 1
	// Comment line 2
	MinLength int64
	/* Heres a description */
	MaxLength int64
	// Doc for Max Age
	MaxAge int64
	// Comments ReadQueue
	ReadQueue string `omit:`
}

//go:generate genfront plain --output table.gen.html --data-file data:conf-options.gen.json --template conf.t
