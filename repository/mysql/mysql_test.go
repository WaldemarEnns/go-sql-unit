package mysql

import (
	"database/sql"
	"log"
	"testing"

	r "github.com/moemoe89/go-unit-tes-sql/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var u = &r.UserModel{
	ID:    uuid.New().String(),
	Name:  "Momo",
	Email: "momo@mail.com",
	Phone: "01729810391",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection.", err)
	}

	return db, mock
}

func TestFindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, email, phone FROM users WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(u.ID, u.Name, u.Email, u.Phone)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.FindById(u.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
