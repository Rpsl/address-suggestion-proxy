package datamodels

import (
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Meta struct {
	Request   string `json:"request"`
	Results   int    `json:"results"`
	Timestamp int64  `json:"timestamp"`
}

// Addresses в виде структуры, чтобы в дальнейшем легко добавлять другие данные
type Addresses []struct {
	Address string `json:"address"`
}

type Suggestion struct {
	Meta       Meta      `json:"meta"`
	Suggestion Addresses `json:"suggestion"`
}

func (sr *Suggestion) Encode() (string, error) {
	str, err := json.Marshal(sr)

	return string(str), err
}

func DecodeSuggestion(obj string) (Suggestion, error) {
	sr := Suggestion{}

	err := json.Unmarshal([]byte(obj), &sr)

	if err != nil {
		log.Error(errors.Wrap(err, "can't decode string into object"))
		return Suggestion{}, err
	}

	return sr, err
}
