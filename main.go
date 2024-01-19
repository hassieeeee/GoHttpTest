package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.open-meteo.com/v1/forecast?latitude=34.40&longitude=132.71&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// レスポンスボディの取得(ReadAll()でポインタの値を読み込んでいるぽい)
	// bodyBytesは　[]byte型
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// データの受け皿の型をインスタンス化？
	// dataは上で宣言した構造体
	data := new(WeatherData)

	// 作った受け皿dataに、読み込んだbodyBytesの中身を代入する関数Unmarshal()
	if err := json.Unmarshal(bodyBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		if err, ok := err.(*json.SyntaxError); ok {
			fmt.Println(string(bodyBytes[err.Offset-1:]))
		}
		log.Fatal(err)
		return
	}
	fmt.Println(data.Current.Time)
	fmt.Println(data.Current.Temperature2m)

}
