package examples

//go:generate genfront doctable --line $GOLINE --output conf-options.json
type Conf2 struct {
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
