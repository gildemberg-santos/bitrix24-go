package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) GetLead(methodType string) {
	pkg.Send(l, pkg.MethodGet, methodType)
}
