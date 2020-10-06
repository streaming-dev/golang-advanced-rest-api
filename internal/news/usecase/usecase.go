package usecase

import (
	"context"
	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/internal/news"
	"github.com/AleksK1NG/api-mc/pkg/logger"
)

// News useCase
type useCase struct {
	logger   *logger.Logger
	cfg      *config.Config
	newsRepo news.Repository
}

// News use case constructor
func NewNewsUseCase(logger *logger.Logger, cfg *config.Config, newsRepo news.Repository) *useCase {
	return &useCase{logger, cfg, newsRepo}
}

// Create news
func (u *useCase) Create(ctx context.Context, news *models.News) (*models.News, error) {
	n, err := u.newsRepo.Create(ctx, news)
	if err != nil {
		return nil, err
	}
	return n, err
}
