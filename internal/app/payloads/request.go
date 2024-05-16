package payloads

type GetProductListParams struct {
	LastValue uint64 `json:"-"`
	NextToken string `json:"-"`
	Limit     int    `json:"-"`
	SellerID  string `json:"-"`
}

type GetOrderListParams struct {
	LastValue uint64 `json:"-"`
	NextToken string `json:"-"`
	Limit     int    `json:"-"`
	Role      string `json:"-"`
	UserID    string `json:"-"`
}

type BuyerLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SellerLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateOrderRequest struct {
	Items      uint64 `json:"items" validate:"required,number"`
	Quantity   uint   `json:"quantity" validate:"required,number"`
	Price      uint64 `json:"price" validate:"required,number"`
	TotalPrice uint64 `json:"total_price" validate:"required,number"`
}

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       uint64 `json:"price" validate:"required,number"`
}
