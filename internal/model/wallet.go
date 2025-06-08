package model

type Wallet struct {
	ID      int64 `json:"id"`
	OwnerID int64 `json:"owner_id"`
	Balance int64 `json:"balance"`
}
