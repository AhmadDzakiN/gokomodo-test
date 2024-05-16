package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
	"gokomodo-assignment/internal/app/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

type SellerTestSuite struct {
	suite.Suite
	sqlMock        sqlmock.Sqlmock
	repositoryMock ISellerRepository
	loc            *time.Location
	ctx            *gin.Context
}

func (s *SellerTestSuite) SetupTest() {
	sqlDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB, DriverName: "postgres", WithoutQuotingCheck: true}), &gorm.Config{})

	s.repositoryMock = NewSellerRepository(gormdb)
	s.sqlMock = mock
	ctx, _ := gin.CreateTestContext(nil)
	s.ctx = ctx
	loc, _ := time.LoadLocation("UTC")
	s.loc = loc
}

func TestSellerSuite(t *testing.T) {
	suite.Run(t, new(SellerTestSuite))
}

func (s *SellerTestSuite) TestGetByID() {
	type fields struct {
		mock func(id string)
	}

	type args struct {
		ctx *gin.Context
		id  string
	}

	query := `SELECT * FROM sellers WHERE id = $1 ORDER BY sellers.id LIMIT $2`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult entity.Seller
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: s.ctx,
				id:  "1",
			},
			fields: fields{
				mock: func(id string) {
					s.sqlMock.MatchExpectationsInOrder(false)
					s.sqlMock.ExpectQuery(query).WithArgs(id, 1).
						WillReturnRows(s.sqlMock.NewRows([]string{"id", "email", "password", "name", "pickup_address", "created_at", "updated_at"}).
							AddRow("1", "sellermatthew@example.com", "sellergreat213", "matthew", "Jakarta", time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
								time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc)))
				}},
			expectedResult: entity.Seller{
				ID:            "1",
				Email:         "sellermatthew@example.com",
				Password:      "sellergreat213",
				Name:          "matthew",
				PickupAddress: "Jakarta",
				CreatedAt:     time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
				UpdatedAt:     time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: s.ctx,
				id:  "1",
			},
			fields: fields{
				mock: func(id string) {
					s.sqlMock.MatchExpectationsInOrder(false)
					s.sqlMock.ExpectQuery(query).WithArgs(id, 1).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: entity.Seller{},
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		s.Suite.Run(test.name, func() {
			test.fields.mock(test.args.id)

			actualResult, actualErr := s.repositoryMock.GetByID(test.args.ctx, test.args.id)

			assert.Equal(s.T(), test.expectedErr, actualErr)
			assert.Equal(s.T(), test.expectedResult, actualResult)
		})
	}
}

func (s *SellerTestSuite) TestGetByEmail() {
	type fields struct {
		mock func(email string)
	}

	type args struct {
		ctx   *gin.Context
		email string
	}

	query := `SELECT * FROM sellers WHERE email = $1 ORDER BY sellers.id LIMIT $2`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult entity.Seller
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx:   s.ctx,
				email: "test@email.com",
			},
			fields: fields{
				mock: func(email string) {
					s.sqlMock.MatchExpectationsInOrder(false)
					s.sqlMock.ExpectQuery(query).WithArgs(email, 1).
						WillReturnRows(s.sqlMock.NewRows([]string{"id", "email", "password", "name", "pickup_address", "created_at", "updated_at"}).
							AddRow("2", "test@email.com", "sellerjohn2", "john pantau", "Bekasi", time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
								time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc)))
				}},
			expectedResult: entity.Seller{
				ID:            "2",
				Email:         "test@email.com",
				Password:      "sellerjohn2",
				Name:          "john pantau",
				PickupAddress: "Bekasi",
				CreatedAt:     time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
				UpdatedAt:     time.Date(2020, 01, 03, 00, 00, 00, 00, s.loc),
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx:   s.ctx,
				email: "test@email.com",
			},
			fields: fields{
				mock: func(email string) {
					s.sqlMock.MatchExpectationsInOrder(false)
					s.sqlMock.ExpectQuery(query).WithArgs(email, 1).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: entity.Seller{},
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		s.Suite.Run(test.name, func() {
			test.fields.mock(test.args.email)

			actualResult, actualErr := s.repositoryMock.GetByEmail(test.args.ctx, test.args.email)

			assert.Equal(s.T(), test.expectedErr, actualErr)
			assert.Equal(s.T(), test.expectedResult, actualResult)
		})
	}
}
