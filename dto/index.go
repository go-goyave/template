package dto

import "goyave.dev/goyave/v5/util/typeutil"

type Index struct {
	Page    typeutil.Undefined[int] `json:"page"`
	PerPage typeutil.Undefined[int] `json:"perPage"`
}
