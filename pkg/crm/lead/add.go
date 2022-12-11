package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) AddLead() {
	l.Title = "titulo de teste"
	l.Name = "nome de teste"
	l.SecondName = "segundo nome de teste"
	l.LastName = "ultimo nome de teste"
	l.StatusId = "1"
	l.Opened = "Y"
	l.AssignedById = "1"
	l.CompanyId = "1"
	l.Opportunity = "1000"
	l.Phone = append(l.Phone, Phone{Value: "11999999999", ValueType: "WORK"})
	l.Phone = append(l.Phone, Phone{Value: "11999999999", ValueType: "WORK"})
	l.Web = append(l.Web, Web{Value: "www.teste.com.br", ValueType: "WORK"})
	pkg.Post(l, pkg.MethodAdd)
}
