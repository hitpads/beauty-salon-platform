package usecase

import (
	"context"
	"errors"
	"testing"
	"user-master-service/internal/domain"
)

type mockUserRepository struct {
	CreateUserFunc     func(ctx context.Context, user *domain.User) error
	GetUserByEmailFunc func(ctx context.Context, email string) (*domain.User, error)
	GetUserByIDFunc    func(ctx context.Context, id string) (*domain.User, error)
}

func (m *mockUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return m.CreateUserFunc(ctx, user)
}
func (m *mockUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return m.GetUserByEmailFunc(ctx, email)
}
func (m *mockUserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return m.GetUserByIDFunc(ctx, id)
}

func TestUserUsecase_Register_Success(t *testing.T) {
	mockRepo := &mockUserRepository{
		CreateUserFunc: func(ctx context.Context, user *domain.User) error {
			return nil
		},
	}
	uc := NewUserUsecase(mockRepo)
	user, err := uc.Register(context.Background(), "Test", "test@example.com", "pass", "user")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", user.Email)
	}
}

func TestUserUsecase_Register_Error(t *testing.T) {
	mockRepo := &mockUserRepository{
		CreateUserFunc: func(ctx context.Context, user *domain.User) error {
			return errors.New("repo error")
		},
	}
	uc := NewUserUsecase(mockRepo)
	_, err := uc.Register(context.Background(), "Test", "test@example.com", "pass", "user")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUserUsecase_Login_Success(t *testing.T) {
	mockRepo := &mockUserRepository{
		GetUserByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
			return &domain.User{Email: email, Password: "pass"}, nil
		},
	}
	uc := NewUserUsecase(mockRepo)
	user, err := uc.Login(context.Background(), "test@example.com", "pass")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", user.Email)
	}
}

func TestUserUsecase_Login_UserNotFound(t *testing.T) {
	mockRepo := &mockUserRepository{
		GetUserByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
			return nil, errors.New("not found")
		},
	}
	uc := NewUserUsecase(mockRepo)
	_, err := uc.Login(context.Background(), "test@example.com", "pass")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUserUsecase_Login_InvalidPassword(t *testing.T) {
	mockRepo := &mockUserRepository{
		GetUserByEmailFunc: func(ctx context.Context, email string) (*domain.User, error) {
			return &domain.User{Email: email, Password: "otherpass"}, nil
		},
	}
	uc := NewUserUsecase(mockRepo)
	_, err := uc.Login(context.Background(), "test@example.com", "pass")
	if err == nil || err.Error() != "invalid password" {
		t.Fatalf("expected invalid password error, got %v", err)
	}
}

func TestUserUsecase_GetProfile_Success(t *testing.T) {
	mockRepo := &mockUserRepository{
		GetUserByIDFunc: func(ctx context.Context, id string) (*domain.User, error) {
			return &domain.User{ID: id, Email: "test@example.com"}, nil
		},
	}
	uc := NewUserUsecase(mockRepo)
	user, err := uc.GetProfile(context.Background(), "some-id")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.ID != "some-id" {
		t.Errorf("expected id some-id, got %s", user.ID)
	}
}

func TestUserUsecase_GetProfile_NotFound(t *testing.T) {
	mockRepo := &mockUserRepository{
		GetUserByIDFunc: func(ctx context.Context, id string) (*domain.User, error) {
			return nil, errors.New("not found")
		},
	}
	uc := NewUserUsecase(mockRepo)
	_, err := uc.GetProfile(context.Background(), "some-id")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
