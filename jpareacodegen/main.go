package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const minArea = 1
const maxArea = 47
const areaKey = "area"

var mlitAPIURL *url.URL

func init() {
	mlitAPIURL, _ = url.Parse("https://www.land.mlit.go.jp/webland/api/CitySearch")
}

type Config struct {
	Area    string
	AreaInt int `opts:"-"`
	Output  string
	Quiet   bool
}

func main() {
	area := flag.Int("area", 0, "")
	output := flag.String("output", "", "")
	quiet := flag.Bool("quiet", false, "")
	flag.Parse()

	if err := main2(*area, *output, *quiet); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
}

func main2(area int, output string, quiet bool) error {
	var results []Result
	var err error

	if area == 0 {
		results, err = All(!quiet)
		if err != nil {
			return err
		}
	} else {
		results, err = MLITAPI(area, !quiet)
		if err != nil {
			return err
		}
	}

	var w io.Writer
	if output == "" {
		w = os.Stdout
	} else {
		f, err := os.Create(output)
		if err != nil {
			return err
		}

		w = f
		defer func() {
			_ = f.Close()
		}()
	}

	for _, r := range results {
		if _, err := io.WriteString(w, fmt.Sprintf("%s,%s,%s\n", r.PrefCode, r.Name, r.Code)); err != nil {
			return err
		}
	}

	return nil
}

type Result struct {
	PrefCode string
	Name     string
	Code     string
}

func All(log bool) ([]Result, error) {
	var results []Result
	for c := minArea; c <= maxArea; c++ {
		r, err := MLITAPI(c, log)
		if err != nil {
			return nil, err
		}
		results = append(results, r...)
	}
	return results, nil
}

func MLITAPI(area int, log bool) (results []Result, _ error) {
	prefCode := fmt.Sprintf("%02d", area)
	u := *mlitAPIURL
	q := url.Values{}
	q.Set(areaKey, prefCode)
	u.RawQuery = q.Encode()
	us := u.String()

	if log {
		os.Stderr.WriteString(us)
		os.Stderr.WriteString("\n")
	}

	res, err := http.Get(us)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	var rawRes struct {
		Status string `json:"status"`
		Data   []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&rawRes); err != nil {
		return nil, err
	}
	if rawRes.Status != "OK" {
		return nil, fmt.Errorf("status is %s", rawRes.Status)
	}

	for _, r := range rawRes.Data {
		results = append(results, Result{
			PrefCode: prefCode,
			Name:     r.Name,
			Code:     r.ID,
		})
	}
	return
}
