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

// TODO: need to specify the tag for input validation

type BuyerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SellerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateOrderRequest struct {
	Items      uint64 `json:"items"`
	Quantity   uint   `json:"quantity"`
	Price      uint64 `json:"price"`
	TotalPrice uint64 `json:"total_price"`
}

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
}
