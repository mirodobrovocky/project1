package item

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Item struct {
	name  		string
	owner 		string
	price		float64
	createdAt 	time.Time
}

type itemBson struct {
	Name  		string  	`bson:"name"`
	Owner 		string  	`bson:"owner"`
	Price		float64		`bson:"price"`
	CreatedAt 	time.Time 	`bson:"createdAt"`
}

func (i Item) Name() string {
	return i.name
}

func (i Item) Owner() string {
	return i.owner
}

func (i Item) Price() float64 {
	return i.price
}

func (i Item) CreatedAt() time.Time {
	return i.createdAt
}

func (i *Item) BeforeSave() {
	i.createdAt = time.Now().Local()
}

func (i Item) MarshalBSON() ([]byte, error) {
	return bson.Marshal(&itemBson{
		Name: i.name,
		Owner: i.owner,
		Price: i.price,
		CreatedAt: i.createdAt,
	})
}

func (i *Item) UnmarshalBSON(data []byte) error {
	temp := &itemBson{}

	if err := bson.Unmarshal(data, temp); err != nil {
		return err
	}

	i.name = temp.Name
	i.owner = temp.Owner
	i.price = temp.Price
	i.createdAt = temp.CreatedAt

	return nil
}

func NewItem(name string, owner string, price float64) Item {
	return Item{
		name:  name,
		owner: owner,
		price: price,
	}
}
