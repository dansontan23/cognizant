package dao

import (
	"database/sql"
	"elibrary/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewDao(t *testing.T) {
	tests := []struct {
		name string
		want Dao
	}{
		{
			name: "success",
			want: &DaoImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDao(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaoImpl_GetBook(t *testing.T) {
	type args struct {
		db    *sql.DB
		title string
	}
	tests := []struct {
		name  string
		d     *DaoImpl
		args  args
		want  *models.BookDetail
		want1 error
	}{
		{
			name: "success",
			args: args{
				title: "test",
			},
			want: &models.BookDetail{
				Title:           "test",
				AvailableCopies: 10,
			},
			want1: nil,
		},
		{
			name: "fail, error",
			args: args{
				title: "test",
			},
			want:  nil,
			want1: sql.ErrConnDone,
		},
		{
			name: "fail, not found",
			args: args{
				title: "test",
			},
			want:  nil,
			want1: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock DB and mock objects
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("could not create mock db: %v", err)
			}
			// Mock the database query
			switch tt.name {
			case "success":
				// Simulating a book record with 10 available copies
				rows := sqlmock.NewRows([]string{"title", "available_copies"}).
					AddRow("test", 10)
				mock.ExpectQuery("SELECT title, available_copies FROM books WHERE title = ?").
					WithArgs(tt.args.title).
					WillReturnRows(rows)

			case "fail, error":
				// Simulate an error while querying the database
				mock.ExpectQuery("SELECT title, available_copies FROM books WHERE title = ?").
					WithArgs(tt.args.title).
					WillReturnError(sql.ErrConnDone)

			case "fail, not found":
				// Simulate no rows returned (book not found)
				mock.ExpectQuery("SELECT title, available_copies FROM books WHERE title = ?").
					WithArgs(tt.args.title).
					WillReturnError(sql.ErrNoRows)
			}

			// Instantiate DaoImpl and call GetBook
			tt.d = &DaoImpl{}
			got, got1 := tt.d.GetBook(db, tt.args.title)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DaoImpl.GetBook() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DaoImpl.GetBook() got1 = %v, want %v", got1, tt.want1)
			}

		})

	}
}
