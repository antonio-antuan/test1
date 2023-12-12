package processor

import (
	"context"

	"github.com/qdm12/go-template/internal/models"
)

type Database interface {
	GetPosts(ctx context.Context) (posts []models.Post, err error)
}
