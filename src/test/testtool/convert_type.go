package testtool

import (
	"github.com/go-openapi/strfmt"
)

func Int32Ptr(n int32) *int32 {
    return &n
}

func StringPtr(s string) *string {
    return &s
}

func StrfmtDateTimePtr(p strfmt.DateTime) *strfmt.DateTime {
	return &p
}