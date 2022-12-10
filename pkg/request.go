package pkg

import (
	"fmt"
)

func Post(lead interface{}, method interface{}) {
	fmt.Println(lead)
	fmt.Println(method)
}
