package service

import (
	"elibrary/dao"
	"elibrary/dao/mocks"
	"elibrary/models"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestServiceImpl_CheckAvailability(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	type args struct {
		title string
	}
	tests := []struct {
		name    string
		s       *ServiceImpl
		args    args
		want    *models.BookDetail
		wantErr error
	}{
		{
			name: "success",
			args: args{
				title: "test",
			},
			s: &ServiceImpl{
				DB: mockDB,
				Dao: func() dao.Dao {
					mockDao := new(mocks.Dao)
					mockDao.On("GetBook", mockDB, "test").Return(&models.BookDetail{
						Title:           "test",
						AvailableCopies: 10,
					}, nil)
					return mockDao
				}(),
			},
			want: &models.BookDetail{
				Title:           "test",
				AvailableCopies: 10,
			},
			wantErr: nil,
		},
		{
			name: "fail",
			args: args{
				title: "test",
			},
			s: &ServiceImpl{
				DB: mockDB,
				Dao: func() dao.Dao {
					mockDao := new(mocks.Dao)
					// Set up the mock expectation here
					mockDao.On("GetBook", mockDB, "test").Return(
						nil,
						fmt.Errorf("error 1"))
					return mockDao
				}(),
			},
			want:    nil,
			wantErr: fmt.Errorf("error 1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CheckAvailability(tt.args.title)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("ServiceImpl.CheckAvailability() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceImpl.CheckAvailability() = %v, want %v", got, tt.want)
			}
		})
	}

}
