package providers

type YandexResponse struct {
	Response struct {
		GeoObjectCollection struct {
			MetaDataProperty struct {
				GeocoderResponseMetaData struct {
					Request string `json:"request"`
					Results int    `json:"results"`
					Suggest string `json:"suggest"`
					Found   string `json:"found"`
				} `json:"GeocoderResponseMetaData"`
			} `json:"metaDataProperty"`
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							Precision string `json:"precision"`
							Text      string `json:"text"`
							Kind      string `json:"kind"`
							Address   struct {
								CountryCode string `json:"country_code"`
								Formatted   string `json:"formatted"`
								Components  []struct {
									Kind string `json:"kind"`
									Name string `json:"name"`
								} `json:"Components"`
							} `json:"Address"`
							AddressDetails struct {
								Country struct {
									AddressLine        string `json:"AddressLine"`
									CountryNameCode    string `json:"CountryNameCode"`
									CountryName        string `json:"CountryName"`
									AdministrativeArea struct {
										AdministrativeAreaName string `json:"AdministrativeAreaName"`
										Locality               struct {
											LocalityName string `json:"LocalityName"`
											Thoroughfare struct {
												ThoroughfareName string `json:"ThoroughfareName"`
											} `json:"Thoroughfare"`
										} `json:"Locality"`
									} `json:"AdministrativeArea"`
								} `json:"Country"`
							} `json:"AddressDetails"`
						} `json:"GeocoderMetaData"`
					} `json:"metaDataProperty"`
					Name        string `json:"name"`
					Description string `json:"description"`
					BoundedBy   struct {
						Envelope struct {
							LowerCorner string `json:"lowerCorner"`
							UpperCorner string `json:"upperCorner"`
						} `json:"Envelope"`
					} `json:"boundedBy"`
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}
