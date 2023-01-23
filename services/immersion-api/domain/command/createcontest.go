package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tadoku/tadoku/services/common/domain"
)

type ContestCreateRequest struct {
	OwnerUserID             uuid.UUID `validate:"required"`
	OwnerUserDisplayName    string    `validate:"required"`
	ContestStart            time.Time `validate:"required"`
	ContestEnd              time.Time `validate:"required"`
	RegistrationEnd         time.Time `validate:"required"`
	Title                   string    `validate:"required,gt=3"`
	Description             *string
	ActivityTypeIDAllowList []int32 `validate:"required,min=1"`

	// Optional
	Official              bool
	Private               bool
	LanguageCodeAllowList []string
}

type ContestCreateResponse struct {
	ID                      uuid.UUID
	ContestStart            time.Time
	ContestEnd              time.Time
	RegistrationEnd         time.Time
	Title                   string
	Description             *string
	OwnerUserID             uuid.UUID
	OwnerUserDisplayName    string
	Official                bool
	Private                 bool
	LanguageCodeAllowList   []string
	ActivityTypeIDAllowList []int32
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

func (s *ServiceImpl) CreateContest(ctx context.Context, req *ContestCreateRequest) (*ContestCreateResponse, error) {
	// Make sure the user is authorized to create a contest
	if domain.IsRole(ctx, domain.RoleBanned) {
		return nil, ErrForbidden
	}
	if domain.IsRole(ctx, domain.RoleGuest) {
		return nil, ErrUnauthorized
	}
	if req.Official && !domain.IsRole(ctx, domain.RoleAdmin) {
		return nil, ErrForbidden
	}

	// Enrich request with session
	session := domain.ParseSession(ctx)
	if session == nil {
		return nil, ErrUnauthorized
	}
	req.OwnerUserID = uuid.MustParse(session.Subject)
	req.OwnerUserDisplayName = session.DisplayName

	err := s.validate.Struct(req)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("unable to validate: %w", ErrInvalidContest)
	}

	if req.Official && req.Private {
		return nil, fmt.Errorf("official rounds cannot be private: %w", ErrInvalidContest)
	}

	if req.Official && len(req.LanguageCodeAllowList) != 0 {
		return nil, fmt.Errorf("official rounds cannot limit language choice: %w", ErrInvalidContest)
	}

	if req.ContestStart.After(req.ContestEnd) {
		return nil, fmt.Errorf("contest cannot start after it has ended: %w", ErrInvalidContest)
	}

	if !domain.IsRole(ctx, domain.RoleAdmin) {
		now := s.clock.Now()
		if now.After(req.ContestEnd) || now.After(req.ContestStart) {
			return nil, fmt.Errorf("contest cannot be in the past or already have started: %w", ErrInvalidContest)
		}
	}

	return s.r.CreateContest(ctx, req)
}
