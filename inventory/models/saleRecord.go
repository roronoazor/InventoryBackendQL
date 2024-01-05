package inventory

import (
	"time"

	auth "github.com/roronoazor/InventoryBackendQL/auth/models"
)

type SaleRecord struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	TotalAmount   int    `json:"totalAmount"`
	CustomerName  string `json:"customerName"`
	CustomerPhone string `json:"customerPhone"`
	CustomerEmail string `json:"customerEmail"`
	ItemsSold     []StoreItem
	CreatedBy     auth.User
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedBy     auth.User
}
