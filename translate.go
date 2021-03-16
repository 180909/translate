package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Type string `json:"type"`
	ErrorCode int `json:"errorCode"`
	ElapsedTime int `json:"elapsedTime"`
	TranslateResult [][]struct {
		Src string `json:"src"`
		Tgt string `json:"tgt"`
	} `json:"translateResult"`
}

func Translate(str string) string {
	url := "http://fanyi.youdao.com/translate?&doctype=json&type=ZH_CN2EN&i=" + str
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	var data Data
	json.Unmarshal([]byte(html), &data)
	fmt.Println(data.TranslateResult)
	for _, o := range data.TranslateResult{
		for _, o1 := range o{
			return o1.Tgt
		}
	}
	return "http error"
}