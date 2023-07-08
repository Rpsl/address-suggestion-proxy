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
	cl     *http.Client
}

func NewYandexProvider(apikey string) (*YandexProvider, error) {
	yp := YandexProvider{
		apikey: apikey,
	}

	return &yp, nil
}

func (yp *YandexProvider) Search(query string) (datamodels.Suggestion, error) {
	client := yp.getHttpClient()

	resp, err := client.Get(yp.getQueryURI(query))
	defer resp.Body.Close()

	if err != nil {
		return datamodels.Suggestion{}, err
	}

	var r YandexResponse

	_ = json.NewDecoder(resp.Body).Decode(&r)
	sr, _ := yp.convertToSearchResponse(&r)

	return sr, nil
}

func (yp *YandexProvider) getHttpClient() *http.Client {
	if yp.cl == nil {
		tr := &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 10 * time.Second,
		}

		yp.cl = &http.Client{Transport: tr}
	}
	return yp.cl
}

func (yp *YandexProvider) getQueryURI(query string) string {
	// https://geocode-maps.yandex.ru/1.x/?apikey=YOUR_API_KEY&geocode=Москва,+Тверская+улица,+дом+7&format=json
	params := url.Values{}
	params.Add("apikey", yp.apikey)
	params.Add("geocode", query)
	params.Add("format", "json")

	return "https://geocode-maps.yandex.ru/1.x/?" + params.Encode()
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
