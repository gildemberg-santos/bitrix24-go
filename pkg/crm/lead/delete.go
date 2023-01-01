package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) DeleteLead(methodType string) {
	pkg.Send(l, pkg.MethodDelete, methodType)
}
