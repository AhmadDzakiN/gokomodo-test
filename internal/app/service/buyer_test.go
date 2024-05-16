package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gokomodo-assignment/internal/app/constant"
	"gokomodo-assignment/internal/app/entity"
	mock_repository "gokomodo-assignment/internal/app/mocks/repository"
	"gokomodo-assignment/internal/app/payloads"
	"gokomodo-assignment/internal/pkg/jwt"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type BuyerTestSuite struct {
	suite.Suite
	serviceMock     IBuyerService
	buyerRepoMock   *mock_repository.MockIBuyerRepository
	productRepoMock *mock_repository.MockIProductRepository
	orderRepoMock   *mock_repository.MockIOrderRepository
	sellerRepoMock  *mock_repository.MockISellerRepository
	validator       *validator.Validate
	ctx             *gin.Context
}

func (b *BuyerTestSuite) SetUpRouter() (router *gin.Engine) {
	router = gin.Default()
	return
}

func (b *BuyerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(b.T())
	defer mockCtrl.Finish()

	gin.SetMode(gin.TestMode)

	b.buyerRepoMock = mock_repository.NewMockIBuyerRepository(mockCtrl)
	b.productRepoMock = mock_repository.NewMockIProductRepository(mockCtrl)
	b.orderRepoMock = mock_repository.NewMockIOrderRepository(mockCtrl)
	b.sellerRepoMock = mock_repository.NewMockISellerRepository(mockCtrl)

	b.validator = validator.New()
	ctx, _ := gin.CreateTestContext(nil)
	b.ctx = ctx

	b.serviceMock = NewBuyerService(b.validator, b.buyerRepoMock, b.productRepoMock, b.orderRepoMock, b.sellerRepoMock)
}

func TestBuyerTestSuite(t *testing.T) {
	suite.Run(t, new(BuyerTestSuite))
}

func (b *BuyerTestSuite) TestGetProductList() {
	productData := []entity.Product{
		{
			ID:          1,
			Name:        "Kipas Angin",
			Description: "Kipas angin dijamin dingin",
			Price:       200000,
			SellerID:    "1",
		},
		{
			ID:          2,
			Name:        "Kulkas 2 Pintu",
			Description: "Kulkas dijamin dingin",
			Price:       5000000,
			SellerID:    "2",
		},
	}

	params := payloads.GetProductListParams{
		Limit: 10,
	}

	type fields struct {
		mock func(ctx *gin.Context)
	}

	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name               string
		args               args
		fields             fields
		expectedStatusCode int
	}{
		{
			name: "Failed to get Product list from repo",
			args: args{ctx: b.ctx},
			fields: fields{mock: func(ctx *gin.Context) {
				b.productRepoMock.EXPECT().GetList(gomock.Any(), params).Return([]entity.Product{}, errors.New("random error"))
			}},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Success",
			args: args{ctx: b.ctx},
			fields: fields{mock: func(ctx *gin.Context) {
				b.productRepoMock.EXPECT().GetList(gomock.Any(), params).Return(productData, nil)
			}},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt.fields.mock(tt.args.ctx)

		r := b.SetUpRouter()
		r.GET("/buyer/products", b.serviceMock.GetProductList)
		req, _ := http.NewRequestWithContext(b.ctx, "GET", "/buyer/products", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(b.T(), tt.expectedStatusCode, w.Code)
	}
}

func (b *BuyerTestSuite) TestLogin() {
	type fields struct {
		mock func(ctx *gin.Context, request payloads.BuyerLoginRequest)
	}

	type args struct {
		ctx        *gin.Context
		reqBody    string
		reqPayload payloads.BuyerLoginRequest
	}

	tests := []struct {
		name               string
		args               args
		fields             fields
		expectedStatusCode int
	}{
		{
			name: "Invalid request body because the body is broken",
			args: args{
				ctx: b.ctx,
				reqBody: `{
							"email": "testbuyer@email.com",
							"password": "test1234",
						  }`,
			},
			fields:             fields{mock: func(ctx *gin.Context, request payloads.BuyerLoginRequest) {}},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid request body the email field is not a valid email",
			args: args{
				ctx: b.ctx,
				reqBody: `{
							"email": "testbuyer",
							"password": "test1234"
						  }`,
			},
			fields:             fields{mock: func(ctx *gin.Context, request payloads.BuyerLoginRequest) {}},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed to get Buyer by Email from repo because buyer is not found",
			args: args{
				ctx: b.ctx,
				reqBody: `{
							"email": "testbuyer@email.com",
							"password": "test1234"
						  }`,
				reqPayload: payloads.BuyerLoginRequest{
					Email:    "testbuyer@email.com",
					Password: "test1234",
				},
			},
			fields: fields{mock: func(ctx *gin.Context, request payloads.BuyerLoginRequest) {
				b.buyerRepoMock.EXPECT().GetByEmail(gomock.Any(), request.Email).Return(entity.Buyer{}, gorm.ErrRecordNotFound)
			}},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "Failed to get Buyer by Email from repo (error in db)",
			args: args{
				ctx: b.ctx,
				reqBody: `{
							"email": "testbuyer@email.com",
							"password": "test1234"
						  }`,
				reqPayload: payloads.BuyerLoginRequest{
					Email:    "testbuyer@email.com",
					Password: "test1234",
				},
			},
			fields: fields{mock: func(ctx *gin.Context, request payloads.BuyerLoginRequest) {
				b.buyerRepoMock.EXPECT().GetByEmail(gomock.Any(), request.Email).Return(entity.Buyer{}, errors.New("random error"))
			}},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Success",
			args: args{
				ctx: b.ctx,
				reqBody: `{
							"email": "testbuyer@email.com",
							"password": "test1234"
						  }`,
				reqPayload: payloads.BuyerLoginRequest{
					Email:    "testbuyer@email.com",
					Password: "test1234",
				},
			},
			fields: fields{mock: func(ctx *gin.Context, request payloads.BuyerLoginRequest) {
				b.buyerRepoMock.EXPECT().GetByEmail(gomock.Any(), request.Email).Return(entity.Buyer{
					ID:   "123",
					Name: "Joko",
				}, nil)
			}},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt.fields.mock(tt.args.ctx, tt.args.reqPayload)

		r := b.SetUpRouter()
		r.POST("/buyer/login", b.serviceMock.Login)
		req, _ := http.NewRequestWithContext(b.ctx, "POST", "/buyer/login", strings.NewReader(tt.args.reqBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(b.T(), tt.expectedStatusCode, w.Code)
	}
}

func (b *BuyerTestSuite) TestGetOrderList() {
	//orderData := []entity.Order{
	//	{
	//		BuyerID:            "1",
	//		SellerID:           "2",
	//		SourceAddress:      "Jakarta",
	//		DestinationAddress: "Bandung",
	//		Items:              3,
	//		Quantity:           2,
	//		Price:              4500000,
	//		TotalPrice:         9000000,
	//		Status:             constant.OrderStatusAccepted,
	//	},
	//}

	params := payloads.GetProductListParams{
		Limit: 10,
	}

	type fields struct {
		mock func(ctx *gin.Context)
	}

	type args struct {
		ctx *gin.Context
	}

	tests := []struct {
		name               string
		args               args
		fields             fields
		expectedStatusCode int
	}{
		{
			name: "Failed to get Product list from repo",
			args: args{ctx: b.ctx},
			fields: fields{mock: func(ctx *gin.Context) {
				ctx.Set("token", jwt.JWTCustomClaims{
					UserID: "1",
					Name:   "Abdul",
					Role:   constant.BuyerRole,
				})

				b.orderRepoMock.EXPECT().GetList(b.ctx, params).Return([]entity.Order{}, errors.New("random error"))
			}},
			expectedStatusCode: http.StatusInternalServerError,
		},
		//{
		//	name: "Success",
		//	args: args{ctx: b.ctx},
		//	fields: fields{mock: func(ctx *gin.Context) {
		//		b.orderRepoMock.EXPECT().GetList(gomock.Any(), params).Return(orderData, nil)
		//	}},
		//	expectedStatusCode: http.StatusOK,
		//},
	}

	for _, tt := range tests {
		tt.fields.mock(tt.args.ctx)

		b.ctx.Set("token", jwt.JWTCustomClaims{
			UserID: "1",
			Name:   "Abdul",
			Role:   constant.BuyerRole,
		})

		r := b.SetUpRouter()
		req, _ := http.NewRequestWithContext(b.ctx, "GET", "/buyer/orders", nil)
		r.GET("/buyer/orders", b.serviceMock.GetOrderList)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(b.T(), tt.expectedStatusCode, w.Code)
		//assert.NotEmpty(b.T(), tasks)
	}
}
