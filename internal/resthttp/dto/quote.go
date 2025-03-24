package dto

type (
	QuoteRequest struct {
		Challenge string `json:"challenge"`
		Nonce     string `json:"nonce"`
	}

	QuoteResponse struct {
		Quote string `json:"quote"`
	}

	ChallengeResponse struct {
		Challenge string `json:"challenge"`
	}
)
