package providers

import (
	"address-suggesstion-proxy/config"
	"address-suggesstion-proxy/internal/datamodels"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Container struct {
	p map[string]Provider
}

func NewProvidersContainer(cfg *config.Config) Container {
	c := Container{}
	c.p = make(map[string]Provider)

	if cfg.YandexEnabled {
		yp, err := NewYandexProvider(cfg.YandexAPIKey)

		if err != nil {
			log.Fatal(errors.Wrap(err, "error while creating yandex provider"))
		}

		c.p[YandexProviderName] = yp

		log.Debug("yandex provider enabled")
	}

	// todo dadata

	return c
}

// SearchInAvailableProviders
// Единый метод поиска данных в разных провайдерах. Нет смысла мержить результаты т.к. они будут схожими.
// Делаем запрос во все провайдеры, смотрим у кого больше данных, возвращаем результат для дальнейшего кеширования.
func (c *Container) SearchInAvailableProviders(query string) datamodels.Suggestion {
	if len(c.p) == 0 {
		log.Errorln("search unavailable because no one provider is enabled")
	}

	var wg sync.WaitGroup

	ch := make(chan datamodels.Suggestion, 1)

	for _, p := range c.p {
		wg.Add(1)

		go func(pr Provider) {
			defer wg.Done()

			res, err := pr.Search(query)

			if err != nil {
				log.Error(errors.Wrap(err, "error search from provider"))
			}

			ch <- res
		}(p)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var results []datamodels.Suggestion

	for {
		v, ok := <-ch

		if !ok {
			break
		}

		results = append(results, v)
	}

	var result datamodels.Suggestion

	max := 0
	// todo нужно сделать какой-нибудь нормальный фильтр, на основе качества результатов
	for _, v := range results {
		if len(v.Suggestion) > max {
			max = len(v.Suggestion)
			result = v
		}
	}

	return result
}
