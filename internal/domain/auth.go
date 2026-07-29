package domain

import "github.com/munnaMia/ahlan/internal/infra/auth"

type TokenService interface {
	GenerateToken(payload auth.Payload) (string, error)
}
