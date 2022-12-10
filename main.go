package main

import (
	"fmt"

	"github.com/gildemberg-santos/bitrix24-go/pkg"
)

func main() {
	teste := pkg.Lead{}
	teste.Title = "titulo de teste"
	teste.Name = "nome de teste"
	teste.SecondName = "segundo nome de teste"
	teste.LastName = "ultimo nome de teste"
	teste.StatusId = "1"
	teste.Opened = "Y"
	teste.AssignedById = "1"
	teste.CompanyId = "1"
	teste.Opportunity = "1000"
	teste.Phone = append(teste.Phone, pkg.Phone{Value: "11999999999", ValueType: "WORK"})
	teste.Phone = append(teste.Phone, pkg.Phone{Value: "11999999999", ValueType: "WORK"})
	teste.Web = append(teste.Web, pkg.Web{Value: "www.teste.com.br", ValueType: "WORK"})
	fmt.Println(teste)
}
