package repository_test

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"user-master-service/internal/domain"
	"user-master-service/internal/repository"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=beauty_salon_test sslmode=disable")
	if err != nil {
		panic("failed to connect to test db: " + err.Error())
	}
	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

func cleanUsersTable(t *testing.T) {
	_, err := testDB.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("failed to clean users table: %v", err)
	}
}

func TestUserRepository_CreateAndGet(t *testing.T) {
	cleanUsersTable(t)
	repo := repository.NewUserRepository(testDB)
	ctx := context.Background()

	user := &domain.User{
		ID:       uuid.New().String(),
		Name:     "IntTest",
		Email:    "inttest@example.com",
		Password: "password",
		Role:     "client",
	}
	// CreateUser
	if err := repo.CreateUser(ctx, user); err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// GetUserByEmail
	got, err := repo.GetUserByEmail(ctx, "inttest@example.com")
	if err != nil {
		t.Fatalf("GetUserByEmail failed: %v", err)
	}
	if got.ID != user.ID || got.Email != user.Email || got.Name != user.Name {
		t.Errorf("Got wrong user: %+v, want: %+v", got, user)
	}

	// GetUserByID
	gotByID, err := repo.GetUserByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("GetUserByID failed: %v", err)
	}
	if gotByID.Email != user.Email {
		t.Errorf("Got wrong user by ID: %+v", gotByID)
	}
}
