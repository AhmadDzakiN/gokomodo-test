package repository

import (
	"fmt"
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

type OrderTestSuite struct {
	suite.Suite
	sqlMock        sqlmock.Sqlmock
	repositoryMock IOrderRepository
	loc            *time.Location
	ctx            *gin.Context
}

func (o *OrderTestSuite) SetupTest() {
	sqlDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB, DriverName: "postgres", WithoutQuotingCheck: true}), &gorm.Config{})

	o.repositoryMock = NewOrderRepository(gormdb)
	o.sqlMock = mock
	ctx, _ := gin.CreateTestContext(nil)
	o.ctx = ctx
	loc, _ := time.LoadLocation("UTC")
	o.loc = loc
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}

func (o *OrderTestSuite) TestGetList() {
	type fields struct {
		mock func(sellerID string, updatedAt uint64, limit int)
	}

	type args struct {
		ctx    *gin.Context
		params payloads.GetOrderListParams
	}

	query := `SELECT o.id, o.buyer_id, o.seller_id, o.source_address, o.destination_address,o.items, o.quantity, o.price, o.total_price, o.status 
FROM orders o WHERE %s AND o.updated_at > $2 LIMIT $3`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult []entity.Order
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: o.ctx,
				params: payloads.GetOrderListParams{
					Limit:     constant.GetItemListLimit,
					UserID:    "1",
					Role:      constant.BuyerRole,
					LastValue: 1234567890,
				},
			},
			fields: fields{
				mock: func(sellerID string, updatedAt uint64, limit int) {
					updatedQuery := fmt.Sprintf(query, "o.buyer_id = $1")
					o.sqlMock.MatchExpectationsInOrder(false)
					o.sqlMock.ExpectQuery(updatedQuery).WithArgs(sellerID, time.Unix(int64(updatedAt), 0), limit).
						WillReturnRows(o.sqlMock.NewRows([]string{"id", "seller_id", "buyer_id", "source_address", "destination_address", "items", "quantity", "price", "total_price", "status"}).
							AddRow(1, "1", "2", "Depok", "Bekasi", 4, 2, 1000000, 2000000, constant.OrderStatusPending))
				},
			},
			expectedResult: []entity.Order{
				{
					ID:                 1,
					SellerID:           "1",
					BuyerID:            "2",
					SourceAddress:      "Depok",
					DestinationAddress: "Bekasi",
					Items:              4,
					Quantity:           2,
					Price:              1000000,
					TotalPrice:         2000000,
					Status:             constant.OrderStatusPending,
				},
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: o.ctx,
				params: payloads.GetOrderListParams{
					Limit:     constant.GetItemListLimit,
					UserID:    "1",
					Role:      constant.SellerRole,
					LastValue: 1234567890,
				},
			},
			fields: fields{
				mock: func(sellerID string, updatedAt uint64, limit int) {
					updatedQuery := fmt.Sprintf(query, "o.seller_id = $1")
					o.sqlMock.MatchExpectationsInOrder(false)
					o.sqlMock.ExpectQuery(updatedQuery).WithArgs(sellerID, time.Unix(int64(updatedAt), 0), limit).WillReturnError(gorm.ErrUnsupportedDriver)
				},
			},
			expectedResult: []entity.Order(nil),
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		o.Suite.Run(test.name, func() {
			test.fields.mock(test.args.params.UserID, test.args.params.LastValue, test.args.params.Limit)

			actualResult, actualErr := o.repositoryMock.GetList(test.args.ctx, test.args.params)

			assert.Equal(o.T(), test.expectedErr, actualErr)
			assert.Equal(o.T(), test.expectedResult, actualResult)
		})
	}
}

func (o *OrderTestSuite) TestCreate() {
	type fields struct {
		mock func(product entity.Order)
	}

	type args struct {
		ctx   *gin.Context
		order *entity.Order
	}

	order := entity.Order{
		BuyerID:            "1",
		SellerID:           "2",
		SourceAddress:      "Jakarta",
		DestinationAddress: "Bandung",
		Items:              3,
		Quantity:           2,
		Price:              4500000,
		TotalPrice:         9000000,
		Status:             constant.OrderStatusPending,
	}

	query := `INSERT INTO orders (buyer_id,seller_id,source_address,destination_address,items,quantity,price,total_price,status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING created_at,updated_at,id`

	tests := []struct {
		name        string
		args        args
		fields      fields
		expectedErr error
	}{
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx:   o.ctx,
				order: &order,
			},
			fields: fields{
				mock: func(order entity.Order) {
					o.sqlMock.ExpectBegin()
					o.sqlMock.ExpectQuery(query).
						WithArgs(order.BuyerID, order.SellerID, order.SourceAddress, order.DestinationAddress, order.Items, order.Quantity, order.Price, order.TotalPrice, order.Status).
						WillReturnError(gorm.ErrUnsupportedDriver)
					o.sqlMock.ExpectRollback()

				}},
			expectedErr: gorm.ErrUnsupportedDriver,
		},
		{
			name: "Success",
			args: args{
				ctx:   o.ctx,
				order: &order,
			},
			fields: fields{
				mock: func(product entity.Order) {
					o.sqlMock.ExpectBegin()
					o.sqlMock.ExpectQuery(query).
						WithArgs(order.BuyerID, order.SellerID, order.SourceAddress, order.DestinationAddress, order.Items, order.Quantity, order.Price, order.TotalPrice, order.Status).
						WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at", "id"}).AddRow(time.Date(2020, 01, 03, 00, 00, 00, 00, o.loc),
							time.Date(2020, 01, 03, 00, 00, 00, 00, o.loc), 1))
					o.sqlMock.ExpectCommit()
				},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		o.Suite.Run(test.name, func() {
			test.fields.mock(*test.args.order)

			actualError := o.repositoryMock.Create(test.args.ctx, test.args.order)

			assert.Equal(o.T(), test.expectedErr, actualError)
		})
	}
}

func (o *OrderTestSuite) TestAccept() {
	type fields struct {
		mock func(id uint64, sellerID string)
	}

	type args struct {
		ctx      *gin.Context
		id       uint64
		sellerID string
	}

	query := `UPDATE orders SET status=$1 WHERE id = $2 AND seller_id = $3`

	tests := []struct {
		name        string
		args        args
		fields      fields
		expectedErr error
	}{
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx:      o.ctx,
				id:       1,
				sellerID: "2",
			},
			fields: fields{
				mock: func(id uint64, sellerID string) {
					o.sqlMock.ExpectBegin()
					o.sqlMock.ExpectExec(query).
						WithArgs(constant.OrderStatusAccepted, id, sellerID).
						WillReturnError(gorm.ErrUnsupportedDriver)
					o.sqlMock.ExpectRollback()
				},
			},
			expectedErr: gorm.ErrUnsupportedDriver,
		},
		{
			name: "Success",
			args: args{
				ctx:      o.ctx,
				id:       3,
				sellerID: "4",
			},
			fields: fields{
				mock: func(id uint64, sellerID string) {
					o.sqlMock.ExpectBegin()
					o.sqlMock.ExpectExec(query).
						WithArgs(constant.OrderStatusAccepted, id, sellerID).
						WillReturnResult(sqlmock.NewResult(1, 1))
					o.sqlMock.ExpectCommit()
				},
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		o.Suite.Run(test.name, func() {
			test.fields.mock(test.args.id, test.args.sellerID)

			actualError := o.repositoryMock.Accept(test.args.ctx, test.args.id, test.args.sellerID)

			assert.Equal(o.T(), test.expectedErr, actualError)
		})
	}
}

func (o *OrderTestSuite) TestGetByID() {
	type fields struct {
		mock func(id uint64)
	}

	type args struct {
		ctx *gin.Context
		id  uint64
	}

	query := `SELECT * FROM orders WHERE id = $1 ORDER BY orders.id LIMIT $2`

	tests := []struct {
		name           string
		args           args
		fields         fields
		expectedResult entity.Order
		expectedErr    error
	}{
		{
			name: "Success",
			args: args{
				ctx: o.ctx,
				id:  1,
			},
			fields: fields{
				mock: func(id uint64) {
					o.sqlMock.MatchExpectationsInOrder(false)
					o.sqlMock.ExpectQuery(query).WithArgs(id, 1).
						WillReturnRows(o.sqlMock.NewRows([]string{"id", "seller_id", "buyer_id", "source_address", "destination_address", "items", "quantity", "price", "total_price", "status", "created_at", "updated_at"}).
							AddRow(1, "1", "2", "Depok", "Bekasi", 4, 2, 1000000, 2000000, constant.OrderStatusPending, time.Date(2020, 01, 03, 00, 00, 00, 00, o.loc),
								time.Date(2020, 01, 03, 00, 00, 00, 00, o.loc)))
				}},
			expectedResult: entity.Order{
				ID:                 1,
				SellerID:           "1",
				BuyerID:            "2",
				SourceAddress:      "Depok",
				DestinationAddress: "Bekasi",
				Items:              4,
				Quantity:           2,
				Price:              1000000,
				TotalPrice:         2000000,
				Status:             constant.OrderStatusPending,
				CreatedAt:          time.Date(2020, 01, 03, 00, 00, 00, 00, o.loc),
				UpdatedAt:          time.Date(2020, 01, 03, 00, 00, 00, 00, o.loc),
			},
			expectedErr: nil,
		},
		{
			name: "Failed, theres an error in db",
			args: args{
				ctx: o.ctx,
				id:  1,
			},
			fields: fields{
				mock: func(id uint64) {
					o.sqlMock.MatchExpectationsInOrder(false)
					o.sqlMock.ExpectQuery(query).WithArgs(id, 1).WillReturnError(gorm.ErrUnsupportedDriver)
				}},
			expectedResult: entity.Order{},
			expectedErr:    gorm.ErrUnsupportedDriver,
		},
	}

	for _, test := range tests {
		o.Suite.Run(test.name, func() {
			test.fields.mock(test.args.id)

			actualResult, actualErr := o.repositoryMock.GetByID(test.args.ctx, test.args.id)

			assert.Equal(o.T(), test.expectedErr, actualErr)
			assert.Equal(o.T(), test.expectedResult, actualResult)
		})
	}
}
