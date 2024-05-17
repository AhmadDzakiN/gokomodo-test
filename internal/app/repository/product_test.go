package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
	"gokomodo-assignment/internal/app/constant"
	"gokomodo-assignment/internal/app/entity"
	"gokomodo-assignment/internal/app/payloads"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

type ProductTestSuite struct {
	suite.Suite
	sqlMock        sqlmock.Sqlmock
	repositoryMock IProductRepository
	loc            *time.Location
	ctx            *gin.Context
}

func (p *ProductTestSuite) SetupTest() {
	sqlDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB, DriverName: "postgres", WithoutQuotingCheck: true}), &gorm.Config{})

	p.repositoryMock = NewProductRepository(gormdb)
	p.sqlMock = mock
	ctx, _ := gin.CreateTestContext(nil)
	p.ctx = ctx
	loc, _ := time.LoadLocation("UTC")
	p.loc = loc
}

func TestProductSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}

func (p *ProductTestSuite) TestGetByID() {
	type fields struct {
		mock func(id uint64)
	}

	type args struct {
		ctx *gin.Context
		id  uint64
	}

	query := `SELECT * FROM products WHERE id = $1 ORDER BY products.id LIMIT $2`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult entity.Product
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: p.ctx,
				id:  1,
			},
			fields: fields{
				mock: func(id uint64) {
					p.sqlMock.MatchExpectationsInOrder(false)
					p.sqlMock.ExpectQuery(query).WithArgs(id, 1).
						WillReturnRows(p.sqlMock.NewRows([]string{"id", "name", "description", "price", "seller_id", "created_at", "updated_at"}).
							AddRow(1, "Handphone keren", "Handphone android terkeren dan terkece", "1000000", "1", time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc),
								time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc)))
				}},
			expectedResult: entity.Product{
				ID:          1,
				Name:        "Handphone keren",
				Description: "Handphone android terkeren dan terkece",
				Price:       1000000,
				SellerID:    "1",
				CreatedAt:   time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc),
				UpdatedAt:   time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc),
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: p.ctx,
				id:  1,
			},
			fields: fields{
				mock: func(id uint64) {
					p.sqlMock.MatchExpectationsInOrder(false)
					p.sqlMock.ExpectQuery(query).WithArgs(id, 1).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: entity.Product{},
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		p.Suite.Run(test.name, func() {
			test.fields.mock(test.args.id)

			actualResult, actualErr := p.repositoryMock.GetByID(test.args.ctx, test.args.id)

			assert.Equal(p.T(), test.expectedErr, actualErr)
			assert.Equal(p.T(), test.expectedResult, actualResult)
		})
	}
}

func (p *ProductTestSuite) TestGetList() {
	type fields struct {
		mock func(sellerID string, updatedAt uint64, limit int)
	}

	type args struct {
		ctx    *gin.Context
		params payloads.GetProductListParams
	}

	query := `SELECT p.id, p.name, p.description, p.price, p.seller_id, p.updated_at FROM products p WHERE p.seller_id = $1 AND p.updated_at > $2 LIMIT $3`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult []entity.Product
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: p.ctx,
				params: payloads.GetProductListParams{
					Limit:     constant.GetItemListLimit,
					SellerID:  "1",
					LastValue: 1234567890,
				},
			},
			fields: fields{
				mock: func(sellerID string, updatedAt uint64, limit int) {
					p.sqlMock.MatchExpectationsInOrder(false)
					p.sqlMock.ExpectQuery(query).WithArgs(sellerID, time.Unix(int64(updatedAt), 0), limit).
						WillReturnRows(p.sqlMock.NewRows([]string{"id", "name", "description", "price", "seller_id", "created_at", "updated_at"}).
							AddRow(1, "Handphone keren", "Handphone android terkeren dan terkece", "1000000", "1", time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc),
								time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc)))
				},
			},
			expectedResult: []entity.Product{
				{
					ID:          1,
					Name:        "Handphone keren",
					Description: "Handphone android terkeren dan terkece",
					Price:       1000000,
					SellerID:    "1",
					CreatedAt:   time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc),
					UpdatedAt:   time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc),
				},
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: p.ctx,
				params: payloads.GetProductListParams{
					Limit:     constant.GetItemListLimit,
					SellerID:  "1",
					LastValue: 1234567890,
				},
			},
			fields: fields{
				mock: func(sellerID string, updatedAt uint64, limit int) {
					p.sqlMock.MatchExpectationsInOrder(false)
					p.sqlMock.ExpectQuery(query).WithArgs(sellerID, time.Unix(int64(updatedAt), 0), limit).WillReturnError(gorm.ErrUnsupportedDriver)
				},
			},
			expectedResult: []entity.Product(nil),
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		p.Suite.Run(test.name, func() {
			test.fields.mock(test.args.params.SellerID, test.args.params.LastValue, test.args.params.Limit)

			actualResult, actualErr := p.repositoryMock.GetList(test.args.ctx, test.args.params)

			assert.Equal(p.T(), test.expectedErr, actualErr)
			assert.Equal(p.T(), test.expectedResult, actualResult)
		})
	}
}

func (p *ProductTestSuite) TestCreate() {
	type fields struct {
		mock func(product entity.Product)
	}

	type args struct {
		ctx     *gin.Context
		product *entity.Product
	}

	product := entity.Product{
		Name:        "Laptop Gaming",
		Description: "Laptop gaming bisa untuk valorant 200 fps",
		Price:       4500000,
		SellerID:    "2",
	}

	query := `INSERT INTO products (name,description,price,seller_id) VALUES ($1,$2,$3,$4) RETURNING created_at,updated_at,id`

	tests := []struct {
		name        string
		args        args
		fields      fields
		expectedErr error
	}{
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx:     p.ctx,
				product: &product,
			},
			fields: fields{
				mock: func(product entity.Product) {
					p.sqlMock.ExpectBegin()
					p.sqlMock.ExpectQuery(query).
						WithArgs(product.Name, product.Description, product.Price, product.SellerID).
						WillReturnError(gorm.ErrUnsupportedDriver)
					p.sqlMock.ExpectRollback()

				}},
			expectedErr: gorm.ErrUnsupportedDriver,
		},
		{
			name: "Success",
			args: args{
				ctx:     p.ctx,
				product: &product,
			},
			fields: fields{
				mock: func(product entity.Product) {
					p.sqlMock.ExpectBegin()
					p.sqlMock.ExpectQuery(query).
						WithArgs(product.Name, product.Description, product.Price, product.SellerID).
						WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at", "id"}).AddRow(time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc),
							time.Date(2020, 01, 03, 00, 00, 00, 00, p.loc), 1))
					p.sqlMock.ExpectCommit()
				},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		p.Suite.Run(test.name, func() {
			test.fields.mock(*test.args.product)

			actualError := p.repositoryMock.Create(test.args.ctx, test.args.product)

			assert.Equal(p.T(), test.expectedErr, actualError)
		})
	}
}
