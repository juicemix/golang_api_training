package prod

import (
	"golang_api/entities/cat"
	"time"
)

type product struct {
	Product_id   int64         `bson:"prod_id" json:"prod_id"`
	Product_name string        `bson:"prod_name" json:"prod_name"`
	Categories   []cat.ProdCat `bson:"prod_cat" json:"prod_cat"`
	Stock        int64         `bson:"prod_stock" json:"prod_stock"`
	Price        float32       `bson:"prod_price" json:"prod_price"`
	Timestamp    time.Time     `bson:"timestamp" json:"timestamp"`
}

type CartProd struct {
	Product_id   int64   `bson:"prod_id" json:"prod_id"`
	Product_name string  `bson:"prod_name" json:"prod_name"`
	Price        float32 `bson:"prod_price" json:"prod_price"`
}
