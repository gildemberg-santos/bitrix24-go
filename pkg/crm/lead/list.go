package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) ListLead() {
	pkg.Post(l, pkg.MethodList)
}
