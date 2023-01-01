package pkg

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

func Send(lead interface{}, method interface{}, methodType string) {
	if methodType == "POST" {
		Post(lead, method)
	} else if methodType == "GET" {
		Get(lead, method)
	}
}

func Post(lead interface{}, method interface{}) {
	fmt.Println("POST")
	fmt.Println(lead)
	fmt.Println(method)
	fmt.Println(Query(lead))
}

func Get(lead interface{}, method interface{}) {
	fmt.Println("GET")
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
