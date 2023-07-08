package providers

import (
	"github.com/h2non/gock"
	"testing"
)

const defResp = "{ \"response\": { \"GeoObjectCollection\": { \"metaDataProperty\": { \"GeocoderResponseMetaData\": { \"request\": \"Москва, улица Новый Арбат, дом 24\", \"found\": \"1\", \"results\": \"10\" } }, \"featureMember\": [ { \"GeoObject\": { \"metaDataProperty\": { \"GeocoderMetaData\": { \"kind\": \"house\", \"text\": \"Россия, Москва, улица Новый Арбат, 24\", \"precision\": \"exact\", \"Address\": { \"country_code\": \"RU\", \"postal_code\": \"119019\", \"formatted\": \"Москва, улица Новый Арбат, 24\", \"Components\": [ { \"kind\": \"country\", \"name\": \"Россия\" }, { \"kind\": \"province\", \"name\": \"Центральный федеральный округ\" }, { \"kind\": \"province\", \"name\": \"Москва\" }, { \"kind\": \"locality\", \"name\": \"Москва\" }, { \"kind\": \"street\", \"name\": \"улица Новый Арбат\" }, { \"kind\": \"house\", \"name\": \"24\" } ] }, \"AddressDetails\": { \"Country\": { \"AddressLine\": \"Москва, улица Новый Арбат, 24\", \"CountryNameCode\": \"RU\", \"CountryName\": \"Россия\", \"AdministrativeArea\": { \"AdministrativeAreaName\": \"Москва\", \"Locality\": { \"LocalityName\": \"Москва\", \"Thoroughfare\": { \"ThoroughfareName\": \"улица Новый Арбат\", \"Premise\": { \"PremiseNumber\": \"24\", \"PostalCode\": { \"PostalCodeNumber\": \"119019\" } } } } } } } } }, \"description\": \"Москва, Россия\", \"name\": \"улица Новый Арбат, 24\", \"boundedBy\": { \"Envelope\": { \"lowerCorner\": \"37.583508 55.750768\", \"upperCorner\": \"37.591719 55.755398\" } }, \"Point\": { \"pos\": \"37.587614 55.753083\" } } } ] } } }"

func TestYandexProvider_Search(t *testing.T) {
	defer gock.Off()
	defer gock.EnableNetworking()

	yp := &YandexProvider{apikey: "MY_KEY"}

	gock.InterceptClient(yp.getHttpClient())
	gock.DisableNetworking()
	gock.New("https://geocode-maps.yandex.ru").
		Get("/1.x/").
		Reply(200).
		JSON(defResp)

	value, err := yp.Search("арбат")

	if err != nil {
		t.Errorf("not expected errors', got %s", err)
	}

	expected := "Москва, улица Новый Арбат, 24"

	if value.Suggestion[0].Address != expected {
		t.Errorf("Expected '%s', got '%s'", expected, value.Suggestion[0].Address)
	}
}
