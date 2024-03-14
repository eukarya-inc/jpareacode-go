//go:generate go run ./jpareacodegen --output data.csv
package jpareacode

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"

	"github.com/eukarya-inc/jpareacode/jpareacodepref"
)

const tokyo23ku = "東京都特別区部"

//go:embed data.csv
var data []byte

var Cities []City
var citym = map[string]City{}
var citymr = map[int]City{}

type City struct {
	PrefCode int    `json:"prefCode"`
	CityCode int    `json:"cityCode"`
	CityName string `json:"cityName"`
	WardCode int    `json:"wardCode"`
	WardName string `json:"wardName"`
}

func (c *City) Code() int {
	if c == nil {
		return 0
	}
	if c.WardCode > 0 {
		return c.WardCode
	}
	return c.CityCode
}

func init() {
	r := csv.NewReader(bytes.NewBuffer(data))
	data, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var lastCityCode int
	var lastCityName string

	for _, record := range data {
		if len(record) < 3 {
			panic("invalid length")
		}

		pcode, err := strconv.Atoi(record[0])
		if err != nil {
			panic("invalid code")
		}

		code, err := strconv.Atoi(record[2])
		if err != nil {
			panic("invalid code")
		}

		var wardCode int
		var wardName string
		var cityCode int
		var cityName string

		if !isWard(record[1]) {
			cityCode = code
			cityName = record[1]
			lastCityCode = cityCode
			lastCityName = cityName
		} else {
			if isTokyo23ku(code) {
				lastCityCode = 13100
				lastCityName = tokyo23ku
			} else {
				cityCode = lastCityCode
				cityName = lastCityName
			}

			wardCode = code
			wardName = record[1]
		}

		city := City{
			PrefCode: pcode,
			CityCode: cityCode,
			CityName: cityName,
			WardCode: wardCode,
			WardName: wardName,
		}

		Cities = append(Cities, city)
		if isTokyo23ku(code) {
			cityName = ""
		}
		citym[record[0]+cityName+wardName] = city
		citymr[code] = city
	}
}

// CityByName は、都道府県コードと市区町村名を基に市区町村情報を返します。
func CityByName(prefCode int, cityName, wardName string) *City {
	if cityName == tokyo23ku {
		cityName = ""
	}

	k := FormatPrefectureCode(prefCode) + cityName + wardName
	c, ok := citym[k]
	if !ok {
		return nil
	}
	return &c
}

// CitiesByName は、市区町村名に合致する全ての市区町村情報を返します。
func CitiesByName(name string) (res []City) {
	for _, c := range Cities {
		if c.CityName == name || c.WardName == name {
			res = append(res, c)
		}
	}
	return
}

// CityByCode は、市区町村コードを基に市区町村情報を返します。
func CityByCode(code int) *City {
	c, ok := citymr[code]
	if !ok {
		return nil
	}
	return &c
}

// CityByCodeString は、市区町村コードの文字列を基に市区町村情報を返します。
func CityByCodeString(code string) *City {
	c, _ := strconv.Atoi(code)
	return CityByCode(c)
}

// SearchCitiesByName は、都道府県名や市区町村名に部分一致する全ての市区町村情報を返します。
func SearchCitiesByName(name string) (res []City) {
	for _, c := range Cities {
		prefName := jpareacodepref.PrefectureNameByCodeInt(c.PrefCode)
		if strings.Contains(prefName, name) {
			res = append(res, c)
			continue
		}

		if strings.Contains(c.CityName, name) || strings.Contains(c.WardName, name) {
			res = append(res, c)
		}
	}
	return
}

// FormatCityCode は、intの市区町村名コードをstringに変換します。無効なコードの場合は空文字列が返されます。
func FormatCityCode(code int) string {
	if !ValidateCityCode(code) {
		return ""
	}
	return fmt.Sprintf("%05d", code)
}

// ParseCityCode は、stringの市区町村名コードをintに変換します。パースに失敗した場合や無効なコードの場合は0が返されます。
func ParseCityCode(code string) int {
	c, _ := strconv.Atoi(code)
	if !ValidateCityCode(c) {
		return 0
	}
	return c
}

// ValidateCityCode は、市区町村コードが有効かどうかを返します。
func ValidateCityCode(code int) bool {
	return code >= PrefectureMinCode*100 && code <= (PrefectureMaxCode+1)*1000-1
}

func isWard(name string) bool {
	return strings.HasSuffix(name, "区")
}

func isTokyo23ku(code int) bool {
	return code >= 13101 && code < 13199
}
