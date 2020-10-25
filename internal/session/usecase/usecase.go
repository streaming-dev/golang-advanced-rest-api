package usecase

import (
	"context"
	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/internal/session"
)

// Session use case
type useCase struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
}

// New session use case constructor
func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config) session.UCSession {
	return &useCase{sessionRepo: sessionRepo, cfg: cfg}
}

// Create new session
func (u *useCase) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	return u.sessionRepo.CreateSession(ctx, session, expire)
}

// Delete session by id
func (u *useCase) DeleteByID(ctx context.Context, sessionID string) error {
	return u.sessionRepo.DeleteByID(ctx, sessionID)
}

// get session by id
func (u *useCase) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {
	return u.sessionRepo.GetSessionByID(ctx, sessionID)
}
