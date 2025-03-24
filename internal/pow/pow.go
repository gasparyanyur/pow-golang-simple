// internal/app/pow/pow.go
package pow

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type (
	proofOfWorkService struct {
		target string
	}

	ProofOfWorkService interface {
		GenerateProofOfWork(challenge string) (string, string)
		VerifyProofOfWork(challenge, nonce string) bool
	}
)

func NewProofOfWorkService(target string) ProofOfWorkService {
	return &proofOfWorkService{target: target}
}

func (pow *proofOfWorkService) GenerateProofOfWork(challenge string) (string, string) {
	nonce := 0
	var hash string
	for {
		data := challenge + strconv.Itoa(nonce)
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		if hash[:len(pow.target)] == pow.target {
			break
		}
		nonce++
	}
	return hash, strconv.Itoa(nonce)
}

func (pow *proofOfWorkService) VerifyProofOfWork(challenge, nonce string) bool {
	data := challenge + nonce
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	return hash[:len(pow.target)] == pow.target
}
