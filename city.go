//go:generate go run ./jpareacodegen --output data.csv
package jpareacode

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"strconv"
)

//go:embed data.csv
var data []byte

var Cities []City
var citym = map[string]City{}
var citymr = map[int]City{}

type City struct {
	PrefCode int    `json:"prefCode"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
}

func init() {
	r := csv.NewReader(bytes.NewBuffer(data))
	data, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
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
		city := City{
			PrefCode: pcode,
			Code:     code,
			Name:     record[1],
		}
		Cities = append(Cities, city)
		citym[record[0]+record[1]] = city
		citymr[code] = city
	}
}

// CityByName は、都道府県コードと市区町村名を基に市区町村情報を返します。
func CityByName(prefCode int, name string) *City {
	k := FormatPrefectureCode(prefCode) + name
	c, ok := citym[k]
	if !ok {
		return nil
	}
	return &c
}

// CitiesByName は、市区町村名に合致する全ての市区町村情報を返します。
func CitiesByName(name string) (res []City) {
	for _, c := range Cities {
		if c.Name == name {
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
