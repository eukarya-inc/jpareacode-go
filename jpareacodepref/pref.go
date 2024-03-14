package jpareacodepref

import (
	"fmt"
	"strconv"
)

var PrefectureMinCode = 1
var PrefectureMaxCode = 47

var Prefectures = []string{
	"北海道",
	"青森県",
	"岩手県",
	"宮城県",
	"秋田県",
	"山形県",
	"福島県",
	"茨城県",
	"栃木県",
	"群馬県",
	"埼玉県",
	"千葉県",
	"東京都",
	"神奈川県",
	"新潟県",
	"富山県",
	"石川県",
	"福井県",
	"山梨県",
	"長野県",
	"岐阜県",
	"静岡県",
	"愛知県",
	"三重県",
	"滋賀県",
	"京都府",
	"大阪府",
	"兵庫県",
	"奈良県",
	"和歌山県",
	"鳥取県",
	"島根県",
	"岡山県",
	"広島県",
	"山口県",
	"徳島県",
	"香川県",
	"愛媛県",
	"高知県",
	"福岡県",
	"佐賀県",
	"長崎県",
	"熊本県",
	"大分県",
	"宮崎県",
	"鹿児島県",
	"沖縄県",
}

var prefm = make(map[string]int, len(Prefectures))

func init() {
	for k, v := range Prefectures {
		prefm[v] = k + 1
	}
}

// PrefectureCodeInt は、都道府県名を基に都道府県コードをintで返します。見つからない場合は0が返されます。
func PrefectureCodeInt(name string) int {
	return prefm[name]
}

// PrefectureCodeString は、都道府県名を基に都道府県コードをstringで返します。見つからない場合は空文字列が返されます。
func PrefectureCodeString(name string) string {
	c := PrefectureCodeInt(name)
	if c == 0 {
		return ""
	}
	return FormatPrefectureCode(c)
}

// PrefectureCodeInts は、複数の都道府県名を基に都道府県コードを[]intで返します。見つからない要素は0になります。
func PrefectureCodeInts(names ...string) (r []int) {
	for _, n := range names {
		r = append(r, PrefectureCodeInt(n))
	}
	return
}

// PrefectureCodeStrings は、複数の都道府県名を基に都道府県コードを[]stringで返します。見つからない要素は空文字列になります。
func PrefectureCodeStrings(names ...string) (r []string) {
	for _, n := range names {
		r = append(r, PrefectureCodeString(n))
	}
	return
}

// PrefectureName は、都道府県コードを基に都道府県名をstringで返します。見つからない場合は空文字列が返されます。
func PrefectureName(code int) string {
	if !ValidatePrefectureCode(code) {
		return ""
	}
	return Prefectures[code-1]
}

// PrefectureNames は、複数の都道府県コードを基に都道府県名を[]stringで返します。見つからない要素は空文字列になります。
func PrefectureNames(code ...int) (r []string) {
	for _, n := range code {
		r = append(r, PrefectureName(n))
	}
	return
}

// FormatPrefectureCode は、intの都道府県コードをstringに変換します。無効なコードの場合は空文字列が返されます。
func FormatPrefectureCode(code int) string {
	if !ValidatePrefectureCode(code) {
		return ""
	}
	return fmt.Sprintf("%02d", code)
}

// ParsePrefectureCode は、stringの都道府県コードをintに変換します。パースに失敗した場合や無効なコードの場合は0が返されます。
func ParsePrefectureCode(code string) int {
	c, _ := strconv.Atoi(code)
	if !ValidatePrefectureCode(c) {
		return 0
	}
	return c
}

// ValidatePrefectureCode は、都道府県コードが有効かどうかを返します。
func ValidatePrefectureCode(code int) bool {
	return code >= PrefectureMinCode && code <= PrefectureMaxCode
}
