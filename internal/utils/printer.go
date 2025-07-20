package utils

import "github.com/k0kubun/pp/v3"

var PPrint *pp.PrettyPrinter

func InitPPrint() {
	PPrint = pp.New()
	PPrint.SetOmitEmpty(true)
}
