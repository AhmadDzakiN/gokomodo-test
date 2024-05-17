package payloads

type GetProductListResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	Seller      string `json:"seller"`
	UpdatedAt   int64  `json:"updated_at"`
}

type GetOrderListResponse struct {
	ID                 uint64 `json:"id"`
	BuyerID            string `json:"buyer_id"`
	SellerID           string `json:"seller_id"`
	SourceAddress      string `json:"source_address"`
	DestinationAddress string `json:"destination_address"`
	Items              uint64 `json:"items"`
	Quantity           uint   `json:"quantity"`
	Price              uint64 `json:"price"`
	TotalPrice         uint64 `json:"total_price"`
	Status             string `json:"status"`
	UpdatedAt          int64  `json:"updated_at"`
}

type CreateOrderResponse struct {
	ID         uint64 `json:"id"`
	Items      uint64 `json:"items"`
	Quantity   uint   `json:"quantity"`
	Price      uint64 `json:"price"`
	TotalPrice uint64 `json:"total_price"`
}

type CreateProductResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	SellerID    string `json:"seller_id"`
}

type SellerLoginResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type BuyerLoginResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type AcceptOrderResponse struct {
	ID    uint64 `json:"id"`
	Items uint64 `json:"items"`
}
