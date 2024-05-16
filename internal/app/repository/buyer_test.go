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

type BuyerTestSuite struct {
	suite.Suite
	sqlMock        sqlmock.Sqlmock
	repositoryMock IBuyerRepository
	loc            *time.Location
	ctx            *gin.Context
}

func (b *BuyerTestSuite) SetupTest() {
	sqlDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB, DriverName: "postgres", WithoutQuotingCheck: true}), &gorm.Config{})

	b.repositoryMock = NewBuyerRepository(gormdb)
	b.sqlMock = mock
	ctx, _ := gin.CreateTestContext(nil)
	b.ctx = ctx
	loc, _ := time.LoadLocation("UTC")
	b.loc = loc
}

func TestBuyerSuite(t *testing.T) {
	suite.Run(t, new(BuyerTestSuite))
}

func (b *BuyerTestSuite) TestGetByID() {
	type fields struct {
		mock func(id string)
	}

	type args struct {
		ctx *gin.Context
		id  string
	}

	query := `SELECT * FROM buyers WHERE id = $1 ORDER BY buyers.id LIMIT $2`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult entity.Buyer
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: b.ctx,
				id:  "1",
			},
			fields: fields{
				mock: func(id string) {
					b.sqlMock.MatchExpectationsInOrder(false)
					b.sqlMock.ExpectQuery(query).WithArgs(id, 1).
						WillReturnRows(b.sqlMock.NewRows([]string{"id", "email", "password", "name", "shipping_address", "created_at", "updated_at"}).
							AddRow("1", "john_doe@example.com", "passdoe213", "john doe", "Jakarta", time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc),
								time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc)))
				}},
			expectedResult: entity.Buyer{
				ID:              "1",
				Email:           "john_doe@example.com",
				Password:        "passdoe213",
				Name:            "john doe",
				ShippingAddress: "Jakarta",
				CreatedAt:       time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc),
				UpdatedAt:       time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc),
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: b.ctx,
				id:  "1",
			},
			fields: fields{
				mock: func(id string) {
					b.sqlMock.MatchExpectationsInOrder(false)
					b.sqlMock.ExpectQuery(query).WithArgs(id, 1).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: entity.Buyer{},
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		b.Suite.Run(test.name, func() {
			test.fields.mock(test.args.id)

			actualResult, actualErr := b.repositoryMock.GetByID(test.args.ctx, test.args.id)

			assert.Equal(b.T(), test.expectedErr, actualErr)
			assert.Equal(b.T(), test.expectedResult, actualResult)
		})
	}
}

func (b *BuyerTestSuite) TestGetByEmail() {
	type fields struct {
		mock func(email string)
	}

	type args struct {
		ctx   *gin.Context
		email string
	}

	query := `SELECT * FROM buyers WHERE email = $1 ORDER BY buyers.id LIMIT $2`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult entity.Buyer
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx:   b.ctx,
				email: "test@email.com",
			},
			fields: fields{
				mock: func(email string) {
					b.sqlMock.MatchExpectationsInOrder(false)
					b.sqlMock.ExpectQuery(query).WithArgs(email, 1).
						WillReturnRows(b.sqlMock.NewRows([]string{"id", "email", "password", "name", "shipping_address", "created_at", "updated_at"}).
							AddRow("2", "test@email.com", "passscott123", "alex scott", "Depok", time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc),
								time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc)))
				}},
			expectedResult: entity.Buyer{
				ID:              "2",
				Email:           "test@email.com",
				Password:        "passscott123",
				Name:            "alex scott",
				ShippingAddress: "Depok",
				CreatedAt:       time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc),
				UpdatedAt:       time.Date(2020, 01, 03, 00, 00, 00, 00, b.loc),
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx:   b.ctx,
				email: "test@email.com",
			},
			fields: fields{
				mock: func(email string) {
					b.sqlMock.MatchExpectationsInOrder(false)
					b.sqlMock.ExpectQuery(query).WithArgs(email, 1).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: entity.Buyer{},
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		b.Suite.Run(test.name, func() {
			test.fields.mock(test.args.email)

			actualResult, actualErr := b.repositoryMock.GetByEmail(test.args.ctx, test.args.email)

			assert.Equal(b.T(), test.expectedErr, actualErr)
			assert.Equal(b.T(), test.expectedResult, actualResult)
		})
	}
}
