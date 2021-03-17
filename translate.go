package translate

import (
	"encoding/json"
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

func Translate(str string, types int) string {
	pattern := "ZH_CN2EN"
	switch types {
	case 0:
		pattern = "ZH_CN2EN"		// 中文转英文
	case 1:
		pattern = "ZH_CN2JA"		// 中文转日文
	case 2:
		pattern = "ZH_CN2KR" 		// 中文转韩文
	case 3:
		pattern = "ZH_CN2FR"		// 中文转法文
	case 4:
		pattern = "ZH_CN2RU"		// 中文转俄语
	case 5:
		pattern = "ZH_CN2SP"		// 中文转西语
	case 6:
		pattern = "EN2ZH_CN"		// 英文转中文
	case 7:
		pattern = "JA2ZH_CN"		// 日语转中文
	case 8:
		pattern = "KR2ZH_CN"		// 韩文转中文
	case 9:
		pattern = "FR2ZH_CN"		// 法语转中文
	case 10:
		pattern = "RU2ZH_CN"		// 俄语转中文
	case 11:
		pattern = "SP2ZH_CN"		// 西语转中文
	}
	url := "http://fanyi.youdao.com/translate?&doctype=json&type=" + pattern + "&i=" + str
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	var data Data
	json.Unmarshal([]byte(html), &data)
	for _, o := range data.TranslateResult{
		for _, o1 := range o{
			return o1.Tgt
		}
	}
	return "http error"
}