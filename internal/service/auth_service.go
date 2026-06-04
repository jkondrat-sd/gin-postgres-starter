// Why: auth service owns authentication rules such as duplicate email checks, password hashing, and JWT creation.
// What to do: add login/register business behavior here instead of putting security logic in handlers.
package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nattakornwarisnarathorn/example-gin/internal/config"
	"github.com/nattakornwarisnarathorn/example-gin/internal/dto"
	"github.com/nattakornwarisnarathorn/example-gin/internal/event"
	"github.com/nattakornwarisnarathorn/example-gin/internal/model"
	"github.com/nattakornwarisnarathorn/example-gin/internal/queue"
	"github.com/nattakornwarisnarathorn/example-gin/internal/repository"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	users  repository.UserRepository
	cfg    *config.Config
	jobs   queue.Publisher
	events event.Publisher
}

func NewAuthService(users repository.UserRepository, cfg *config.Config, jobs queue.Publisher, events event.Publisher) *AuthService {
	return &AuthService{users: users, cfg: cfg, jobs: jobs, events: events}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	if _, err := s.users.FindByEmail(req.Email); err == nil {
		return nil, ErrConflict
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := model.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Name:         req.Name,
		Role:         "user",
	}
	if err := s.users.Create(&user); err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	go s.dispatchRegistrationSideEffects(context.WithoutCancel(ctx), user)

	return &dto.AuthResponse{
		Token: token,
		User:  dto.ToUserResponse(user),
	}, nil
}

func (s *AuthService) dispatchRegistrationSideEffects(ctx context.Context, user model.User) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if s.jobs != nil {
		err := s.jobs.EnqueueWelcomeEmail(ctx, queue.WelcomeEmailPayload{
			UserID: user.ID,
			Email:  user.Email,
			Name:   user.Name,
		})
		if err != nil {
			logger.Log.Warn("welcome email job enqueue failed", zap.Error(err), zap.Uint("user_id", user.ID))
		}
	}

	if s.events != nil {
		err := s.events.PublishUserRegistered(ctx, event.UserRegisteredEvent{
			EventType: event.EventUserRegistered,
			UserID:    user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		})
		if err != nil {
			logger.Log.Warn("user registered event publish failed", zap.Error(err), zap.Uint("user_id", user.ID))
		}
	}
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.users.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidLogin
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidLogin
	}

	token, err := s.generateToken(*user)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User:  dto.ToUserResponse(*user),
	}, nil
}

func (s *AuthService) generateToken(user model.User) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.cfg.JWTExpireHours) * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    s.cfg.AppName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}
