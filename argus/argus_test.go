package argus

import (
	"fmt"
	"testing"
)

func TestArgus(t *testing.T) {
	argus, err := NewArgusService("rpc", "argusAddress", "execPrivateKey")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//callData := MakeCallData(nil, nil, nil, nil)
	tx, err := argus.ExecTransaction(nil, CallData{})
	fmt.Println(tx)
}
