package providers

import (
	"address-suggesstion-proxy/internal/datamodels"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const YandexProviderName = "yandex"

type YandexProvider struct {
	apikey string
}

func NewYandexProvider(apikey string) (*YandexProvider, error) {
	yp := YandexProvider{
		apikey: apikey,
	}

	return &yp, nil
}

func (yp *YandexProvider) Search(query string) (datamodels.Suggestion, error) {
	// https://geocode-maps.yandex.ru/1.x/?apikey=YOUR_API_KEY&geocode=Москва,+Тверская+улица,+дом+7&format=json
	params := url.Values{}
	params.Add("apikey", yp.apikey)
	params.Add("geocode", query)
	params.Add("format", "json")

	resp, err := http.Get("https://geocode-maps.yandex.ru/1.x/?" + params.Encode())

	if err != nil {
		return datamodels.Suggestion{}, err
	}

	defer resp.Body.Close()

	var r YandexResponse
	json.NewDecoder(resp.Body).Decode(&r)

	sr, err := yp.convertToSearchResponse(&r)

	return sr, nil
}

func (yp *YandexProvider) convertToSearchResponse(yr *YandexResponse) (datamodels.Suggestion, error) {
	sr := datamodels.Suggestion{
		Meta: struct {
			Request   string `json:"request"`
			Results   int    `json:"results"`
			Timestamp int64  `json:"timestamp"`
		}{},
		Suggestion: nil,
	}

	sr.Meta.Request = yr.Response.GeoObjectCollection.MetaDataProperty.GeocoderResponseMetaData.Request
	sr.Meta.Timestamp = time.Now().Unix()

	for _, v := range yr.Response.GeoObjectCollection.FeatureMember {
		addr := struct {
			Address string `json:"address"`
		}{
			Address: v.GeoObject.MetaDataProperty.GeocoderMetaData.Address.Formatted,
		}

		sr.Suggestion = append(sr.Suggestion, addr)
	}

	sr.Meta.Results = len(sr.Suggestion)

	return sr, nil
}
