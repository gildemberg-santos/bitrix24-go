package pkg

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

func Send(lead any, method any, methodType string) {
	if methodType == MethodTypePost {
		Post(lead, method)
	} else if methodType == MethodTypeGet {
		Get(lead, method)
	}
}

func Post(lead any, method any) {
	fmt.Println("POST")
	fmt.Println(lead)
	fmt.Println(method)
	fmt.Println(Query(lead))
}

func Get(lead any, method any) {
	fmt.Println("GET")
	fmt.Println(lead)
	fmt.Println(method)
	fmt.Println(Query(lead))
}

func Query(lead any) string {
	v, err := query.Values(lead)
	if err != nil {
		return ""
	}

	return v.Encode()
}
