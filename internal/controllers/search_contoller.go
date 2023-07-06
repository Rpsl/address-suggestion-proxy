package controllers

import (
	"address-suggesstion-proxy/internal/datamodels"
	"address-suggesstion-proxy/internal/providers"
	"address-suggesstion-proxy/internal/services"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func SearchByText(ctx iris.Context, service services.CacheService, cont providers.Container) {
	q := ctx.URLParamEscape("s")

	res, err := service.GetByKey(q)

	if err == nil && res != "" {
		s, err := datamodels.DecodeSuggestion(res)

		if err != nil {
			log.Errorln(errors.Wrap(err, "can't decode object into suggestion"))

			go func() {
				_ = service.DeleteByKey(q)
			}()
		} else {
			ctx.JSON(iris.Map{"status": iris.StatusOK, "data": s})
			return
		}
	}

	s := cont.SearchInAvailableProviders(q)

	str, err := s.Encode()

	if err != nil {
		log.Errorln(errors.Wrap(err, "can't encode suggest to string"))
	} else {
		service.SetByKey(q, str)
	}

	if err != nil {
		ctx.Err()
	}

	ctx.JSON(iris.Map{"status": iris.StatusOK, "data": s})
}
