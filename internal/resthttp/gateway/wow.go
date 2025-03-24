package gateway

import "wow/internal/resthttp/dto"

type (
	wowAPIGateway struct {
		url string
	}

	WOWAPIGateway interface {
		GetChallenge() *dto.ChallengeResponse
		GetQuote(request *dto.QuoteRequest) (*dto.QuoteResponse, error)
	}
)

func NewWowAPIGateway(url string) WOWAPIGateway {
	return &wowAPIGateway{url: url}
}

func (g *wowAPIGateway) GetChallenge() *dto.ChallengeResponse {
	return &dto.ChallengeResponse{}
}

func (g *wowAPIGateway) GetQuote(request *dto.QuoteRequest) (*dto.QuoteResponse, error) {
	return &dto.QuoteResponse{}, nil
}
