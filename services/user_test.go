package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mauroccvieira/restapi-echo-go/db"
	"github.com/mauroccvieira/restapi-echo-go/models"
	"github.com/mauroccvieira/restapi-echo-go/stores"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceContext_GetUsers_Success(t *testing.T) {
	mockDB, mock := db.Mock()
	defer mockDB.Close()

	s := stores.New(mockDB)
	services := New(s)

	rows := sqlmock.NewRows([]string{"id", "name", "password", "username"}).
		AddRow(1, "Testname", "Passwordtest", "Usernametest").
		AddRow(2, "Testname2", "Passwordtest2", "Usernametest2")

	mock.
		ExpectQuery("SELECT id, name, password, username from users").
		WillReturnRows(rows)

	r, err := services.User.GetUsers()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(r))
}

func TestUserServiceContext_CreateUser_Success(t *testing.T) {
	mockDB, mock := db.Mock()
	defer mockDB.Close()

	s := stores.New(mockDB)
	services := New(s)

	a := &models.User{
		Name:     "Testname",
		Username: "Usernametest",
		Password: "Passwordtest",
	}

	mock.
		ExpectQuery("INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id").
		WithArgs(a.Name, a.Username, a.Password).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1),
		)

	r, err := services.User.CreateUser(a)

	a.ID = r.ID

	assert.NoError(t, err)
	assert.Equal(t, *a, r)
}
