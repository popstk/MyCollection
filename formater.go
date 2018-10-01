package main

import (
	"errors"
	"strconv"
	"strings"
)

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
	"332NAMA",
	"SSNI",
	"ENFD",
	"XAM",
	"IDBD",
	"AVZG",
	"CHN",
	"INCT",
	"KWBD",
	"PornoEigakan",
}

// FormatName format names
func FormatName(name string) (string, error) {
	name = strings.ToUpper(name)

	for _, sep := range []byte{'-', '_'} {
		name = strings.Replace(name, string(sep), "", -1)
	}

	prefix := ""
	for _, pre := range prefixs {
		if strings.HasPrefix(name, strings.ToUpper(pre)) {
			prefix = pre
			break
		}
	}

	if prefix == "" {
		return "", errors.New("no match prefix")
	}

	rest := name[len(prefix):]
	n, err := strconv.Atoi(rest)
	if err != nil {
		return "", errors.New("not number:" + rest)
	}

	return prefix + "-" + strconv.Itoa(n), nil
}
