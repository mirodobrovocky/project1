package item

type Item struct {
	name  	string  `bson:"name"`
	owner 	string  `bson:"owner"`
	price	float64	`bson:"price"`
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

func NewItem(name string, owner string, price float64) Item {
	return Item{
		name:  name,
		owner: owner,
		price: price,
	}
}
