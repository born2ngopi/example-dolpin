package check

import "github.com/born2ngopi/example-dolpin/types"

type Coba struct {
	Name string
}

func CheckFunction(msg types.Message, check Coba) string {

	msg.Name = "CheckFunction"

	return msg.Name
}
