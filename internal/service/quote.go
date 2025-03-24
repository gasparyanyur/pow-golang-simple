package service

import (
	"fmt"
	"math/rand"
	"time"

	"wow/internal/pow"
	"wow/internal/repository"
)

type (
	quoteService struct {
		repo repository.QuoteRepository
		pow  pow.ProofOfWorkService
	}

	QuoteService interface {
		GetQuote(challenge, nonce string) (string, error)
	}
)

func NewQuoteService(repo repository.QuoteRepository, pow pow.ProofOfWorkService) QuoteService {
	return &quoteService{
		repo: repo,
		pow:  pow,
	}
}

func (s *quoteService) GetQuote(challenge, nonce string) (string, error) {
	if !s.pow.VerifyProofOfWork(challenge, nonce) {
		return "", fmt.Errorf("invalid Proof of Work")
	}

	quotes, err := s.repo.GetQuotes()
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))], nil
}
