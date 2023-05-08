package stores

import (
	"testing"

	"github.com/mauroccvieira/restapi-echo-go/db"
	"github.com/mauroccvieira/restapi-echo-go/models"
	"github.com/stretchr/testify/assert"
)

func TestUserStore_GetSuccessCase(t *testing.T) {
	mockDB, mock := db.Mock()
	defer mockDB.Close()

	users := []models.User{
		{ID: 1, Username: "kvothe", Name: "Kvothe, the bloodless", Password: "bloodless"},
		{ID: 2, Username: "denna", Name: "Denna", Password: "denna"},
	}

	rows := mock.NewRows([]string{"id", "name", "password", "username"})
	for _, a := range users {
		rows.AddRow(a.ID, a.Name, a.Password, a.Username)
	}

	mock.
		ExpectQuery("SELECT id, name, password, username from users").
		WillReturnRows(rows)

	s := New(mockDB)

	r, err := s.User.Get(nil)

	assert.NoError(t, err)
	assert.Equal(t, users, r)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// func TestUserStore_CreateSuccessCase(t *testing.T) {
// 	mockDB, mock := db.Mock()
// 	defer mockDB.Close()

// 	a := &models.User{
// 		ID:      1,
// 		Name:    "test",
// 		Country: "US",
// 	}

// 	mock.NewRows([]string{"id", "name", "country"})
// 	mock.
// 		ExpectQuery("INSERT INTO users (name, country) VALUES ($1, $2) RETURNING id").
// 		WithArgs(a.Name, a.Country).
// 		WillReturnRows(
// 			sqlmock.NewRows([]string{"id"}).AddRow(1),
// 		)

// 	s := stores.New(mockDB)

// 	r, err := s.User.Create(nil, a)

// 	assert.NoError(t, err)
// 	assert.Equal(t, int64(1), r)
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestUserStore_UpdateByIdSuccessCase(t *testing.T) {
// 	mockDB, mock := db.Mock()
// 	defer mockDB.Close()

// 	a := &models.User{
// 		ID:      1,
// 		Name:    "test",
// 		Country: "US",
// 	}

// 	mock.NewRows([]string{"id", "name", "country"}).AddRow(a.ID, a.Name, a.Country)

// 	a.Name = "new name"
// 	a.Country = "new country"

// 	pr := mock.ExpectPrepare("UPDATE users SET name = $1, country = $2 WHERE users.id = $3 RETURNING id")
// 	pr.
// 		ExpectQuery().
// 		WithArgs(a.Name, a.Country, a.ID).
// 		WillReturnRows(
// 			sqlmock.NewRows([]string{"id"}).AddRow(a.ID),
// 		)

// 	s := stores.New(mockDB)

// 	r, err := s.User.UpdateById(nil, a)

// 	assert.NoError(t, err)
// 	assert.Equal(t, int64(a.ID), r)
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestUserStore_DeleteByIdSuccessCase(t *testing.T) {
// 	mockDB, mock := db.Mock()
// 	defer mockDB.Close()

// 	users := []models.User{
// 		{ID: 1, Name: "test1", Country: "US"},
// 		{ID: 2, Name: "test2", Country: "UK"},
// 	}

// 	deletingID := users[0].ID

// 	rows := mock.NewRows([]string{"id", "name", "country"})
// 	for _, a := range users {
// 		rows.AddRow(a.ID, a.Name, a.Country)
// 	}
// 	mock.
// 		ExpectExec("DELETE FROM users WHERE users.id = $1 RETURNING users.id").
// 		WithArgs(deletingID).
// 		WillReturnResult(sqlmock.NewResult(int64(deletingID), 1))
// 	mock.
// 		ExpectExec("DELETE FROM users WHERE users.id = $1 RETURNING users.id").
// 		WithArgs(deletingID).
// 		WillReturnResult(sqlmock.NewResult(int64(deletingID), 0))

// 	s := stores.New(mockDB)

// 	assert.NoError(t, s.User.DeleteById(nil, deletingID))
// 	assert.Equal(t, s.User.DeleteById(nil, deletingID), sql.ErrNoRows)
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }
