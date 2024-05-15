package payloads

type GetProductListParams struct {
	LastValue uint64 `json:"-"`
	NextToken string `json:"-"`
	Limit     int    `json:"-"`
	Query     string `json:"-"`
	SellerID  string `json:"-"`
}

type GetOrderListParams struct {
	LastValue uint64 `json:"-"`
	NextToken string `json:"-"`
	Limit     int    `json:"-"`
	Query     string `json:"-"`
	Role      string `json:"-"`
}

type BuyerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SellerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
