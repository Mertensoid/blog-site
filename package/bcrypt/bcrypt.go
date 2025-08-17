package bcrypt

import (
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type Crypto struct {
	logger *zerolog.Logger
}

func NewCrypto(logger *zerolog.Logger) *Crypto {
	return &Crypto{
		logger: logger,
	}
}

func (h *Crypto) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Error().Msg(err.Error())
	}
	return string(bytes)
}

func (h *Crypto) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		h.logger.Error().Msg(err.Error())
	}
	return err == nil
}
