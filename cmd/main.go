package main

import (
	"fmt"
	"os"

	"github.com/eukarya-inc/jpareacode"
	"github.com/eukarya-inc/jpareacode/jpareacodepref"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("Usage: jpareacode <area code or city name>\n\nFind city info by area code or city name.\n")
		os.Exit(1)
	}

	param := os.Args[1]
	if isNunmeric(param) {
		city := jpareacode.CityByCodeString(param)
		if city != nil {
			printCity(city)
		} else {
			pref := jpareacodepref.PrefectureNameByCodeString(param)
			printPref(jpareacodepref.Prefecture{
				Name:       pref,
				CodeString: param,
			})
		}
	} else {
		found := false
		prefCodes := jpareacodepref.SearchPrefectures(param)
		for _, p := range prefCodes {
			found = true
			printPref(p)
		}

		cities := jpareacode.SearchCitiesByName(param)
		for _, c := range cities {
			found = true
			printCity(&c)
		}

		if !found {
			os.Stderr.WriteString("City not found\n")
			os.Exit(1)
		}
	}
}

func printCity(c *jpareacode.City) {
	if c == nil {
		os.Stderr.WriteString("City not found\n")
		os.Exit(1)
	}
	os.Stdout.WriteString(cityToString(c) + "\n")
}

func printPref(p jpareacode.Prefecture) {
	if p.Name == "" || p.CodeString == "" {
		os.Stderr.WriteString("Prefecture not found\n")
		os.Exit(1)
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
