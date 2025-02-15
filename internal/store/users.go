package store

import (
	"context"
	"database/sql"

	"github.com/KengoWada/go-todos/internal/models"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *models.User, userProfile *models.UserProfile) error {
	return withTransaction(ctx, s.db, func(tx *sql.Tx) error {
		if err := createUser(ctx, tx, user); err != nil {
			return err
		}

		userProfile.UserID = user.ID
		if err := createUserProfile(ctx, tx, userProfile); err != nil {
			return err
		}

		return nil
	})

}

func createUser(ctx context.Context, tx *sql.Tx, user *models.User) error {
	query := `
		INSERT INTO users(email, password)
		VALUES($1, $2)
		RETURNING id, version, created_at, updated_at, deleted_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := tx.QueryRowContext(ctx, query, user.Email, user.Password).Scan(
		&user.ID,
		&user.Version,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		// TODO: Check the type of error.
		return err
	}

	return nil
}

func createUserProfile(ctx context.Context, tx *sql.Tx, userProfile *models.UserProfile) error {
	query := `
		INSERT INTO user_profiles(name, user_id)
		VALUES($1, $2)
		RETURNING id, version, created_at, updated_at, deleted_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := tx.QueryRowContext(ctx, query, userProfile.Name, userProfile.UserID).Scan(
		&userProfile.ID,
		&userProfile.Version,
		&userProfile.CreatedAt,
		&userProfile.UpdatedAt,
		&userProfile.DeletedAt,
	)

	if err != nil {
		// TODO: Check the type of error.
		return err
	}

	return nil
}
