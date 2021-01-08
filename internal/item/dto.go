package item

import "time"

type CreateDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"gt=0"`
}

type ReadDto struct {
	Name  		string  	`json:"name"`
	Owner 		string  	`json:"owner"`
	Price 		float64 	`json:"price"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}
