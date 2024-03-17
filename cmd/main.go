package main

import (
	"fmt"
	"os"

	"github.com/eukarya-inc/jpareacode"
	"github.com/eukarya-inc/jpareacode/jpareacodepref"
)

func main() {
	invalid := false
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" || len(os.Args[1]) == 0 {
		invalid = true
	}

	if invalid {
		os.Stderr.WriteString("Usage: jpareacode <area code or city name>\n\nFind city info by area code or city name.\n")
		os.Exit(1)
	}

	params := os.Args[1:]

	if isNunmeric(params[0][:1]) {
		for _, code := range params {
			city := jpareacode.CityByCodeString(code)
			if city != nil {
				printCity(city, len(params) > 1)
			} else {
				pref := jpareacodepref.PrefectureNameByCodeString(code)
				printPref(jpareacodepref.Prefecture{
					Name:       pref,
					CodeString: code,
				}, len(params) > 1)
			}
		}
	} else {
		param := params[0]

		found := false
		prefCodes := jpareacodepref.SearchPrefectures(param)
		for _, p := range prefCodes {
			found = true
			printPref(p, false)
		}

		cities := jpareacode.SearchCitiesByName(param)
		for _, c := range cities {
			found = true
			printCity(&c, false)
		}

		if !found {
			os.Stderr.WriteString("City not found\n")
			os.Exit(1)
		}
	}
}

func printCity(c *jpareacode.City, ignoreErr bool) {
	if c == nil {
		if !ignoreErr {
			os.Stderr.WriteString("City not found\n")
			os.Exit(1)
		}
	}
	os.Stdout.WriteString(cityToString(c) + "\n")
}

func printPref(p jpareacode.Prefecture, ignoreErr bool) {
	if p.Name == "" || p.CodeString == "" {
		if !ignoreErr {
			os.Stderr.WriteString("Prefecture not found\n")
			os.Exit(1)
		}
	}
	os.Stdout.WriteString(prefToString(p.Name, p.CodeString) + "\n")
}

const sep = " "

func prefToString(name, code string) string {
	return fmt.Sprintf("%s%s(%s)", name, sep, code)
}

func cityToString(c *jpareacode.City) string {
	cityCode := jpareacode.FormatCityCode(c.CityCode)
	var ward string
	if c.WardName != "" {
		wardCode := jpareacode.FormatCityCode(c.WardCode)
		ward = fmt.Sprintf("%s%s%s(%s)", sep, c.WardName, sep, wardCode)
	}

	prefName := jpareacodepref.PrefectureNameByCodeInt(c.PrefCode)
	prefCode := jpareacodepref.FormatPrefectureCode(c.PrefCode)
	return fmt.Sprintf("%s%s%s%s(%s)%s", prefToString(prefName, prefCode), sep, c.CityName, sep, cityCode, ward)
}

func isNunmeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
