package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"wow/internal/resthttp/dto"
	"wow/internal/service"
)

type Handler struct {
	service service.QuoteService
}

func NewHandler(service service.QuoteService) *Handler {
	return &Handler{
		service: service,
	}
}

// GetChallenge sends a random challenge string to the client
func (h *Handler) GetChallenge(c *gin.Context) {

	// Generate random challenge
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	challenge := fmt.Sprintf("challenge-%d", randGen.Intn(1000000))

	// Send the challenge as a response
	c.JSON(http.StatusOK, &dto.ChallengeResponse{Challenge: challenge})
	return
}

// GetQuote verifies PoW and sends a quote if valid
func (h *Handler) GetQuote(c *gin.Context) {
	var req dto.QuoteRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quote, err := h.service.GetQuote(req.Challenge, req.Nonce)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &dto.QuoteResponse{Quote: quote})
}
