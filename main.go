package main

import (
	"github.com/gildemberg-santos/bitrix24-go/pkg/crm/lead"
)

func main() {
	teste := lead.Lead{}
	teste.AddLead()
	teste.EditLead()
	teste.DeleteLead()
	teste.GetLead()
	teste.ListLead()
}
