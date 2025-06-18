package response

import (
	pkgSices "github.com/sreekar2307/khata/pkg/slices"
	"github.com/sreekar2307/khata/service"
)

type Balance struct {
	Owes  []Owe  `json:"owes,omitempty"`
	Lends []Lend `json:"lends,omitempty"`
}

func NewBalance(owes []service.Owes, lends []service.Lends) Balance {
	return Balance{
		Owes: pkgSices.Map(owes, func(owe service.Owes) Owe {
			return NewOwe(owe)
		}),
		Lends: pkgSices.Map(lends, func(lend service.Lends) Lend {
			return NewLend(lend)
		}),
	}
}
