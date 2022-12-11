package pkg

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

func Post(lead interface{}, method interface{}) {
	fmt.Println(lead)
	fmt.Println(method)
	fmt.Println(Query(lead))
}

func Query(lead interface{}) string {
	v, err := query.Values(lead)
	if err != nil {
		return ""
	}

	return v.Encode()
}
