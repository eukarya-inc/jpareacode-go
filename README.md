# jpareacode

[![Go Reference](https://pkg.go.dev/badge/github.com/eukarya-inc/jpareacode.svg)](https://pkg.go.dev/github.com/eukarya-inc/jpareacode)

日本の都道府県名や市区町村名から、都道府県コードや市区町村コードを取得するライブラリ

```go
package main

import (
  jpareacode "github.com/eukarya-inc/jpareacode"
)

func main() {
  jpareacode.PrefectureCodeInt("北海道") // 1
  jpareacode.PrefectureCodeStrings("北海道") // "01"
  jpareacode.PrefectureCodeInts("北海道", "東京都") // [1, 13]
  jpareacode.PrefectureCodeStrings("北海道", "東京都") // ["01", "13"]
  jpareacode.PrefectureNameByCodeInt(1) // "北海道"
  jpareacode.PrefectureNamesByCodeInt(1, 13) // ["北海道", "東京都"]
  jpareacode.SearchPrefectures // []Prefecture{{Name:"北海道", Code:1}, ...}
  jpareacode.Prefectures // []stirng{"北海道", ...}

  jpareacode.CityByName(13, "", "千代田区") // &City{PrefCode:13, CityName:"東京都特別区部", CityCode:13100, WardName:"千代田区", WardCode:13101}
  jpareacode.CityByNames("北区") // []City{...}
  jpareacode.SearchCitiesByName("北区") // []City{...}
  jpareacode.CityByCodeInt(13101) // &City{PrefCode:13, CityName:"東京都特別区部", CityCode:13100, WardName:"千代田区", WardCode:13101}
  jpareacode.CityByCodeString("13101") // &City{PrefCode:13, CityName:"東京都特別区部", CityCode:13100, WarName:"千代田区", WardCode:13101}
  jpareacode.Cities // []City{{PrefCode:13, CityName:"東京都特別区部", CityCode:13100, WardName:"千代田区", WardCode:13101}, ...}
}
```

データは国土交通省が公開する「都道府県内市区町村一覧取得API」を基にしています。詳しくは以下をご覧ください。

[https://www.land.mlit.go.jp/webland/api.html](https://www.land.mlit.go.jp/webland/api.html)

最終データ更新日：2023/02/20

[MIT License](LICENSE)
