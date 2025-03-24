package repository

type (
	quoteRepository struct {
		quotes []string
	}

	QuoteRepository interface {
		GetQuotes() ([]string, error)
	}
)

func NewQuoteRepository() QuoteRepository {
	return &quoteRepository{
		quotes: []string{
			"Do not dwell in the past, do not dream of the future, concentrate the mind on the present moment.",
			"Life is 10% what happens to us and 90% how we react to it.",
			"An unexamined life is not worth living.",
		},
	}
}

func (r *quoteRepository) GetQuotes() ([]string, error) {
	return r.quotes, nil
}
