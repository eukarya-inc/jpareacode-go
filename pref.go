package jpareacode

import "github.com/eukarya-inc/jpareacode/jpareacodepref"

// PrefectureCodeInt は、都道府県名を基に都道府県コードをintで返します。見つからない場合は0が返されます。
var PrefectureCodeInt = jpareacodepref.PrefectureCodeInt

// PrefectureCodeString は、都道府県名を基に都道府県コードをstringで返します。見つからない場合は空文字列が返されます。
var PrefectureCodeString = jpareacodepref.PrefectureCodeString

// PrefectureCodeInts は、複数の都道府県名を基に都道府県コードを[]intで返します。見つからない要素は0になります。
var PrefectureCodeInts = jpareacodepref.PrefectureCodeInts

// PrefectureCodeStrings は、複数の都道府県名を基に都道府県コードを[]stringで返します。見つからない要素は空文字列になります。
var PrefectureCodeStrings = jpareacodepref.PrefectureCodeStrings

// PrefectureName は、都道府県コードを基に都道府県名をstringで返します。見つからない場合は空文字列が返されます。
var PrefectureName = jpareacodepref.PrefectureName

// PrefectureNames は、複数の都道府県コードを基に都道府県名を[]stringで返します。見つからない要素は空文字列になります。
var PrefectureNames = jpareacodepref.PrefectureNames

// FormatPrefectureCode は、intの都道府県コードをstringに変換します。
var FormatPrefectureCode = jpareacodepref.FormatPrefectureCode

// ParsePrefectureCode は、stringの都道府県コードをintに変換します。パースに失敗した場合やコードが無効な場合は0が返されます。
var ParsePrefectureCode = jpareacodepref.ParsePrefectureCode

// ValidatePrefectureCode は、都道府県コードが有効かどうかを返します。
var ValidatePrefectureCode = jpareacodepref.ValidatePrefectureCode

var PrefectureMinCode = jpareacodepref.PrefectureMinCode
var PrefectureMaxCode = jpareacodepref.PrefectureMaxCode
var Prefectures = jpareacodepref.Prefectures
