package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func GetWeather(req string) string {
	var city string
	if len(req) == 0 {
		return ""
	}
	if value, exists := CITYTABLE[req]; exists {
		city = value
	} else {
		return ""
	}
	currentApi := viper.GetString("CURRENTAPI")
	key := viper.GetString("KEY")
	lang := viper.GetString("LANG")
	uri := fmt.Sprintf("%s?q=%s&appid=%s&lang=%s&units=metric", currentApi, city, key, lang)
	resp, err := http.Get(uri)
	if err != nil {
		log.Println(err)
	}
	var payload WeatherData

	json.NewDecoder(resp.Body).Decode(&payload)
	res := prettyOutput(&payload)
	return res
}

func prettyOutput(payload *WeatherData) string {
	if value, exists := REVERSCITYTABLE[payload.Name]; exists {
		payload.Name = value
	}

	str := fmt.Sprintf(`%s의 날씨입니다.
현재 %s은(는) %s 입니다. 
현재 %s은(는) 최저기온: %.1f°C, 최고기온: %.1f°C 입니다. 
풍속 %.1fm\s 이며, 시간당 %.1fmm의 비가 내리고있습니다.
오늘도 화이팅 하세여`, payload.Name, payload.Name,
		payload.Weather[0].Main,
		payload.Name,
		payload.Main.TempMin, payload.Main.TempMax,
		payload.Wind.Speed, payload.Rain.OneH)
	return str
}
