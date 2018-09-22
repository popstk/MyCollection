package main

import "strings"

var prefixs = []string{
	"IPTD",
	"SIRO",
	"IPX",
	"SNIS",
	"BAGBD",
	"200GANA",
	"ABP",
	"SNIS",
	"IPZ",
	"KAWD",
	"300MIUM",
	"259LUXU",
	"FSET",
	"PT",
	"MIRD",
	"MIDE",
	"RBD",
	"WWW",
	"LOVE",
	"REAL",
	"WANZ",
	"MIAD",
	"TokyoHot",
	"RKI",
	"HUNT",
	"URMC",
	"SDMU",
	"LALS",
	"APKH",
	"OKSN",
	"ZEX",
	"CLUB",
	"MIDE",
	"OFJE",
	"ONEZ",
	"PGD",
	"HIZ",
	"261ARA",
	"MUGON",
	"STAR",
	"YRH",
	"CWP",
	"AKB",
	"HND",
}

// FormatName format names
func FormatName(name string) string {
	name = strings.ToUpper(name)
	name = strings.Replace(name, "-", "", -1)
	for _, pre := range prefixs {
		if strings.HasPrefix(name, pre) {
			name = pre + "-" + name[len(pre):]
			return name
		}
	}
	return ""
}
