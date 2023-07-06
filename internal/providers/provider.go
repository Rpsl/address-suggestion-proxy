package providers

import "address-suggesstion-proxy/internal/datamodels"

type Provider interface {
	Search(query string) (datamodels.Suggestion, error)
}
