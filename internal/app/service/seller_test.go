package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gokomodo-assignment/internal/app/entity"
	mock_repository "gokomodo-assignment/internal/app/mocks/repository"
	"gokomodo-assignment/internal/app/payloads"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type SellerTestSuite struct {
	suite.Suite
	serviceMock     ISellerService
	productRepoMock *mock_repository.MockIProductRepository
	orderRepoMock   *mock_repository.MockIOrderRepository
	sellerRepoMock  *mock_repository.MockISellerRepository
	validator       *validator.Validate
	ctx             *gin.Context
}

func (s *SellerTestSuite) SetUpRouter() (router *gin.Engine) {
	router = gin.Default()
	return
}

func (s *SellerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()

	s.productRepoMock = mock_repository.NewMockIProductRepository(mockCtrl)
	s.orderRepoMock = mock_repository.NewMockIOrderRepository(mockCtrl)
	s.sellerRepoMock = mock_repository.NewMockISellerRepository(mockCtrl)

	s.validator = validator.New()
	ctx, _ := gin.CreateTestContext(nil)
	s.ctx = ctx

	s.serviceMock = NewSellerService(s.validator, s.sellerRepoMock, s.productRepoMock, s.orderRepoMock)
}

func TestSellerTestSuite(t *testing.T) {
	suite.Run(t, new(SellerTestSuite))
}

func (s *SellerTestSuite) TestLogin() {
	type fields struct {
		mock func(ctx *gin.Context, request payloads.SellerLoginRequest)
	}

	type args struct {
		ctx        *gin.Context
		reqBody    string
		reqPayload payloads.SellerLoginRequest
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
				ctx: s.ctx,
				reqBody: `{
							"email": "seller@email.com",
							"password": "test1234",
						  }`,
			},
			fields:             fields{mock: func(ctx *gin.Context, request payloads.SellerLoginRequest) {}},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid request body the email field is not a valid email",
			args: args{
				ctx: s.ctx,
				reqBody: `{
							"email": "seller",
							"password": "test1234"
						  }`,
			},
			fields:             fields{mock: func(ctx *gin.Context, request payloads.SellerLoginRequest) {}},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Failed to get Buyer by Email from repo because buyer is not found",
			args: args{
				ctx: s.ctx,
				reqBody: `{
							"email": "seller@email.com",
							"password": "test1234"
						  }`,
				reqPayload: payloads.SellerLoginRequest{
					Email:    "seller@email.com",
					Password: "test1234",
				},
			},
			fields: fields{mock: func(ctx *gin.Context, request payloads.SellerLoginRequest) {
				s.sellerRepoMock.EXPECT().GetByEmail(gomock.Any(), request.Email).Return(entity.Seller{}, gorm.ErrRecordNotFound)
			}},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "Failed to get Buyer by Email from repo (error in db)",
			args: args{
				ctx: s.ctx,
				reqBody: `{
							"email": "seller@email.com",
							"password": "test1234"
						  }`,
				reqPayload: payloads.SellerLoginRequest{
					Email:    "seller@email.com",
					Password: "test1234",
				},
			},
			fields: fields{mock: func(ctx *gin.Context, request payloads.SellerLoginRequest) {
				s.sellerRepoMock.EXPECT().GetByEmail(gomock.Any(), request.Email).Return(entity.Seller{}, errors.New("random error"))
			}},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "Success",
			args: args{
				ctx: s.ctx,
				reqBody: `{
							"email": "seller@email.com",
							"password": "test1234"
						  }`,
				reqPayload: payloads.SellerLoginRequest{
					Email:    "seller@email.com",
					Password: "test1234",
				},
			},
			fields: fields{mock: func(ctx *gin.Context, request payloads.SellerLoginRequest) {
				s.sellerRepoMock.EXPECT().GetByEmail(gomock.Any(), request.Email).Return(entity.Seller{
					ID:   "123",
					Name: "Joko",
				}, nil)
			}},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt.fields.mock(tt.args.ctx, tt.args.reqPayload)

		r := s.SetUpRouter()
		r.POST("/seller/login", s.serviceMock.Login)
		req, _ := http.NewRequestWithContext(s.ctx, "POST", "/seller/login", strings.NewReader(tt.args.reqBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(s.T(), tt.expectedStatusCode, w.Code)
	}
}
