package session

import (
	"context"
	"time"

	"github.com/jestradaramos/group-lift/backend/pkg/domain"
	"github.com/jestradaramos/group-lift/backend/pkg/repo"
	"github.com/jestradaramos/group-lift/backend/pkg/repo/bun"
	"github.com/jestradaramos/group-lift/backend/pkg/repo/models"
)

type LiftSessionService struct {
	domain.UnimplementedLiftSessionServiceServer
	DB repo.Repository
}

func NewSessionServiceServer(db *bun.DB) *LiftSessionService {
	return &LiftSessionService{
		DB: db,
	}
}

func (s LiftSessionService) AddSession(ctx context.Context, req *domain.AddSessionRequest) (*domain.AddSessionResponse, error) {
	liftSession := &models.LiftSession{
		Date:   time.Now(),
		UserID: req.UserId,
		Lift:   []*models.Lift{},
	}

	liftSession, err := s.DB.AddLiftSession(context.Background(), liftSession)
	if err != nil {
		return nil, err
	}

	domainSession := sessionFromModelToDomain(liftSession)
	res := &domain.AddSessionResponse{Session: domainSession}
	return res, nil
}

func (s LiftSessionService) AddLift(ctx context.Context, req *domain.AddLiftRequest) (*domain.AddLiftResponse, error) {
	lift := &models.Lift{
		SessionID: req.SessionId,
		Lift:      req.Workout,
		Weight:    int(req.Weight),
		Feel:      req.Feel.String(),
	}
	lift, err := s.DB.AddLift(context.Background(), lift)
	if err != nil {
		return nil, err
	}

	domainLift := liftFromModelToDomain(lift)
	res := &domain.AddLiftResponse{Lift: domainLift}
	return res, nil
}

func (s LiftSessionService) GetSession(ctx context.Context, req *domain.GetSessionRequest) (*domain.GetSessionResponse, error) {
	liftSession, err := s.DB.GetLiftSession(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	domainLifts := liftsFromModelToDomain(liftSession.Lift)
	res := &domain.GetSessionResponse{
		Id:     liftSession.ID,
		Date:   liftSession.Date.String(),
		UserId: liftSession.UserID,
		Lifts:  domainLifts,
	}
	return res, nil
}

func (s LiftSessionService) ListSessionsByUser(ctx context.Context, req *domain.ListSessionsByUserRequest) (*domain.ListSessionByUserResponse, error) {
	sessions, err := s.DB.ListLiftSessionsByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	domainSessions := sessionsFromModelToDomain(sessions)
	res := &domain.ListSessionByUserResponse{
		UserId:   req.UserId,
		Sessions: domainSessions,
	}
	return res, nil
}

func sessionFromModelToDomain(liftSession *models.LiftSession) *domain.Session {
	domainUser := domain.Session{
		Date:   liftSession.Date.Format(time.UnixDate),
		SeshId: liftSession.ID,
	}

	return &domainUser
}

func sessionsFromModelToDomain(sessions []*models.LiftSession) []*domain.Session {
	domainSessions := make([]*domain.Session, 0)
	for _, s := range sessions {
		domainLifts := liftsFromModelToDomain(s.Lift)
		session := domain.Session{
			Date:   s.Date.String(),
			SeshId: s.ID,
			UserId: s.UserID,
			Lifts:  domainLifts,
		}
		domainSessions = append(domainSessions, &session)
	}
	return domainSessions
}

func liftFromModelToDomain(lift *models.Lift) *domain.Lift {
	domainUser := domain.Lift{
		SessionId: lift.SessionID,
		Lift:      lift.Lift,
		Weight:    int64(lift.Weight),
		Feel:      lift.Feel,
	}

	return &domainUser
}

func liftsFromModelToDomain(lifts []*models.Lift) []*domain.Lift {
	domainLifts := make([]*domain.Lift, 0)
	for _, l := range lifts {
		lift := domain.Lift{
			Lift:   l.Lift,
			Weight: int64(l.Weight),
			Feel:   l.Feel,
		}
		domainLifts = append(domainLifts, &lift)
	}

	return domainLifts
}
