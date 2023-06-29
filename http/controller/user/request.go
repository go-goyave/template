package user

import (
	"goyave.dev/goyave/v5"
	v "goyave.dev/goyave/v5/validation"
)

func IndexRequest(_ *goyave.Request) v.RuleSet {
	return v.RuleSet{
		{Path: "page", Rules: v.List{v.Int(), v.Min(1)}},
		{Path: "perPage", Rules: v.List{v.Int(), v.Between(1, 100)}},
	}
}
