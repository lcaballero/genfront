package doctable_simple2

type MonoStat func()
type CountStat func()

//go:generate genfront doctable --line $GOLINE --output conf-options.json
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

//go:generate genfront plain --output table.html --data-file data:conf-options.json --template conf.t
