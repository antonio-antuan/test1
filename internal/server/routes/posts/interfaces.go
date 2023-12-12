package posts

import (
	"context"

	"github.com/qdm12/go-template/internal/models"
)

type Logger interface {
	Debugf(format string, args ...any)
	Error(s string)
}

type Processor interface {
	GetPosts(ctx context.Context) (user []models.Post, err error)
}
