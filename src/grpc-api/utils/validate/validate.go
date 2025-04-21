package validate

import protovalidate "github.com/bufbuild/protovalidate-go"

var Validate protovalidate.Validator

func init() {
	v, _ := protovalidate.New()
	Validate = v
}
