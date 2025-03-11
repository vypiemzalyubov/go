package storage

import (
	"context"
	"database/sql"
	"fmt"

	"gitlab.ozon.dev/route256/wallet/internal/pkg/domain"
)

func (s *storage) CreateUser(ctx context.Context, user *domain.User) error {
	stmt := s.Builder().Insert("users").
		Columns("name", "lastname", "age", "phone", "password_hash", "level").
		Values(user.Name, user.Lastname, user.Age, user.Phone, user.PasswordHash, user.Level).
		Suffix("RETURNING id")

	req, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	err = s.db.GetContext(ctx, &user.ID, req, args...)

	return err
}

func (s *storage) UpgradeUser(ctx context.Context, userID string, level string) error {
	stmt := s.Builder().
		Update("users").
		Set("level", level).
		Where("id = ?", userID)

	req, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, req, args...)

	return err
}

func (s *storage) LogIn(ctx context.Context, userID, token string) error {
	stmt := s.Builder().
		Insert("sessions").
		Columns("user_id", "token").
		Values(userID, token)

	req, args, err := stmt.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, req, args...)

	return err
}

func (s *storage) GetUser(ctx context.Context, userID string) (*domain.UserWithAccounts, error) {
	user, err := s.getUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	accounts, err := s.GetUserAccounts(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.UserWithAccounts{
		User:     *user,
		Accounts: accounts,
	}, nil
}

func (s *storage) getUser(ctx context.Context, userID string) (*domain.User, error) {
	var user domain.User

	stmt := s.Builder().
		Select("*").
		From("users").
		Where("id = ?", userID)

	req, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.GetContext(ctx, &user, req, args...)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *storage) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	var user domain.User

	stmt := s.Builder().
		Select("*").
		From("users").
		Where("phone = ?", phone)

	req, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.GetContext(ctx, &user, req, args...)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *storage) GetUserSessions(ctx context.Context, userID string) ([]string, error) {
	var tokens []string

	stmt := s.Builder().
		Select("token").
		From("sessions").
		Where("user_id = ?", userID)

	req, args, err := stmt.ToSql()
	if err != nil {
		return nil, err
	}

	err = s.db.SelectContext(ctx, &tokens, req, args...)

	if err == sql.ErrNoRows {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
