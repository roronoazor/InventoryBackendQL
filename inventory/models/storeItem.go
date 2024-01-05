package inventory

import (
	"time"

	auth "github.com/roronoazor/InventoryBackendQL/auth/models"
)

type StoreItem struct {
	ID              string `json:"id,omitempty" bson:"_id,omitempty"`
	Text            string `json:"text"`
	Description     string `json:"description"`
	CurrentQuantity int    `json:"currentQuantity"`
	CreatedBy       auth.User
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedBy       auth.User
}

// enum StockAction {
// 	SALES
// 	ADD_TO_STOCK
// 	REMOVE_FROM_STOCK
//   }

//   type StockHistory {
// 	action: StockAction!
// 	quantity: Int!
// 	createdBy: User!
// 	updatedBy: User!
//   }
