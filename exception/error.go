package exception

import (
	"fmt"
)

func PanicLogging(err interface{}) {
	if err != nil {
		fmt.Println("ada error", err)
		panic(err)
	}
}
