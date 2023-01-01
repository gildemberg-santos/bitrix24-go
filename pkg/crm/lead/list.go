package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) ListLead(methodType string) {
	pkg.Send(l, pkg.MethodList, methodType)
}
