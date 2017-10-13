package front_simple1

type Rest struct {}
func (r *Rest) Method(method string) *Rest {
	return r
}

type Req struct {}
func (r *Req) Method(method string) *Req {
	return r
}

//go:generate genfront front --input req_rest_methods.fm --output req_rest_methods.gen.go