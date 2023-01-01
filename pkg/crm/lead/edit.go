package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) EditLead(methodType string) {
	pkg.Send(l, pkg.MethodEdit, methodType)
}
