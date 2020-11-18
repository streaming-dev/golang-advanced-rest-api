package usecase

import (
	"context"
	"fmt"
	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/comments"
	"github.com/AleksK1NG/api-mc/internal/models"
	"github.com/AleksK1NG/api-mc/pkg/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	basePrefix = "api-comments:"
)

// Comments UseCase
type commentsUC struct {
	cfg      *config.Config
	commRepo comments.Repository
}

// Comments UseCase constructor
func NewCommentsUseCase(cfg *config.Config, commRepo comments.Repository) comments.UseCase {
	return &commentsUC{cfg: cfg, commRepo: commRepo}
}

// Create comment
func (u *commentsUC) Create(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	return u.commRepo.Create(ctx, comment)
}

// Update comment
func (u *commentsUC) Update(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	comm, err := u.commRepo.GetByID(ctx, comment.CommentID)
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateIsOwner(ctx, comm.AuthorID.String()); err != nil {
		return nil, errors.Wrap(err, "commentsUC Update ValidateIsOwner")
	}

	updatedComment, err := u.commRepo.Update(ctx, comment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

// Delete comment
func (u *commentsUC) Delete(ctx context.Context, commentID uuid.UUID) error {
	comm, err := u.commRepo.GetByID(ctx, commentID)
	if err != nil {
		return err
	}

	if err = utils.ValidateIsOwner(ctx, comm.AuthorID.String()); err != nil {
		return errors.Wrap(err, "commentsUC Delete ValidateIsOwner")
	}

	if err = u.commRepo.Delete(ctx, commentID); err != nil {
		return err
	}

	return nil
}

// GetByID comment
func (u *commentsUC) GetByID(ctx context.Context, commentID uuid.UUID) (*models.CommentBase, error) {
	return u.commRepo.GetByID(ctx, commentID)
}

// GetAllByNewsID comments
func (u *commentsUC) GetAllByNewsID(ctx context.Context, newsID uuid.UUID, query *utils.PaginationQuery) (*models.CommentsList, error) {
	return u.commRepo.GetAllByNewsID(ctx, newsID, query)
}

func (u *commentsUC) createKey(commentID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, commentID)
}
