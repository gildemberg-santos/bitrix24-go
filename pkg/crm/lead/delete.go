package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) DeleteLead() {
	pkg.Post(l, pkg.MethodDelete)
}
