package processor

import (
	"context"
	"github.com/qdm12/go-template/internal/models"
)

func (p *Processor) GetPosts(ctx context.Context) (posts []models.Post, err error) {
	posts, err = p.db.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
