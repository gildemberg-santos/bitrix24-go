package main

import (
	"github.com/gildemberg-santos/bitrix24-go/pkg"
	"github.com/gildemberg-santos/bitrix24-go/pkg/crm/lead"
)

func main() {
	teste := lead.Lead{
		Title:        "titulo de teste",
		Name:         "nome de teste",
		SecondName:   "segundo nome de teste",
		LastName:     "ultimo nome de teste",
		StatusId:     "1",
		Opened:       "Y",
		AssignedById: "1",
		CompanyId:    "1",
		Opportunity:  "1000",
		Phone:        []lead.Phone{{Value: "11999999999", ValueType: "WORK"}, {Value: "11999999999", ValueType: "WORK"}},
		Web:          []lead.Web{{Value: "www.teste.com.br", ValueType: "WORK"}},
	}

	teste.AddLead(pkg.MethodTypeGet)
}
