package item

type Item struct {
	Name  	string  `bson:"name"`
	Owner 	string  `bson:"owner"`
	Price	float64	`bson:"price"`
}
