package lead

import "github.com/gildemberg-santos/bitrix24-go/pkg"

func (l *Lead) GetLead() {
	pkg.Post(l, pkg.MethodGet)
}
