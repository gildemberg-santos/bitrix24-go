package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) AddLead(methodType string) {
	pkg.Send(l, pkg.MethodAdd, methodType)
}
