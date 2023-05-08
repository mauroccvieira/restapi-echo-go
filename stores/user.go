package stores

import (
	"database/sql"

	"github.com/mauroccvieira/restapi-echo-go/logger"
	"github.com/mauroccvieira/restapi-echo-go/models"
	"go.uber.org/zap"
)

type (
	UserStore interface {
		Get(tx *sql.Tx) ([]models.User, error)
		// Create(tx *sql.Tx, user *models.User) (int64, error)
		// UpdateById(tx *sql.Tx, user *models.User) (int64, error)
		// DeleteById(tx *sql.Tx, id int) error
	}

	userStore struct {
		*sql.DB
	}
)

func (s *userStore) Get(tx *sql.Tx) ([]models.User, error) {

	users := make([]models.User, 0)

	rows, err := s.Query("SELECT id, name, password, username from users")

	if err != nil {
		logger.Error("Failed to query users: ", zap.Error(err))
		return nil, err
	}

	for rows.Next() {

		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Username)

		if err != nil {
			logger.Error("Failed to query users: ", zap.Error(err))
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
