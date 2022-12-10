package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) EditLead() {
	pkg.Post(l, pkg.MethodEdit)
}
