package envvar

import (
	"os"

	"github.com/daverussell13/to-do-list-bts/internal/domain"
	"github.com/joho/godotenv"
)

type Provider interface {
	Get(key string) (string, error)
}

type Configuration struct {
	provider Provider
}

func Load(filename string) error {
	if err := godotenv.Load(filename); err != nil {
		return domain.NewErrorf(domain.ErrorCodeUnknown, "loading env var file")
	}
	return nil
}

func New(provider Provider) *Configuration {
	return &Configuration{
		provider: provider,
	}
}

func (c *Configuration) Get(key string) (string, error) {
	res := os.Getenv(key)
	return res, nil
}
