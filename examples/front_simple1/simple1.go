package front_simple1

//go:generate genfront front --input req_rest_methods.fm --output req_rest_methods.gen.go

type Rest int
func (r *Rest) Method(s string) *Rest {
	return r
}


type Req int
func (r *Req) Method(s string) *Req {
	return r
}