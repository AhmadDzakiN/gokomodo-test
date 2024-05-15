package payloads

type GetProductListResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	Seller      string `json:"seller"`
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
}

type CreateOrderResponse struct {
	ID         uint64 `json:"id"`
	Items      uint64 `json:"items"`
	Quantity   uint   `json:"quantity"`
	Price      uint64 `json:"price"`
	TotalPrice uint64 `json:"total_price"`
}
