package prod

import (
	"golang_api/entities/cat"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type product struct {
	Product_id   int64         `bson:"prod_id" json:"prod_id"`
	Product_name string        `bson:"prod_name" json:"prod_name"`
	Categories   []cat.ProdCat `bson:"prod_cat" json:"prod_cat"`
	Stock        int64         `bson:"prod_stock" json:"prod_stock"`
	Price        float32       `bson:"prod_price" json:"prod_price"`
	Timestamp    time.Time     `bson:"timestamp" json:"timestamp"`
}

type productInput struct {
	Product_id    int64   `bson:"prod_id" json:"prod_id"`
	Product_name  string  `bson:"prod_name" json:"prod_name"`
	Categories_id []int64 `bson:"prod_cat" json:"prod_cat"`
	Stock         int64   `bson:"prod_stock" json:"prod_stock"`
	Price         float32 `bson:"prod_price" json:"prod_price"`
}

type CartProd struct {
	Product_id   int64   `bson:"prod_id" json:"prod_id"`
	Product_name string  `bson:"prod_name" json:"prod_name"`
	Price        float32 `bson:"prod_price" json:"prod_price"`
}

func NewProduct(id int64, name string, stock int64, price float32, categories ...cat.ProdCat) product {
	return product{
		Product_id:   id,
		Product_name: name,
		Categories:   categories,
		Stock:        stock,
		Price:        price,
		Timestamp:    time.Now(),
	}
}

func NewProductInput(id int64, name string, stock int64, price float32, categories ...int64) productInput {
	return productInput{
		Product_id:    id,
		Product_name:  name,
		Categories_id: categories,
		Stock:         stock,
		Price:         price,
	}
}

func getOne(fkey string, fvalue string) (product, error) {
	var ret product

	session, err := mgo.Dial("localhost")
	if err != nil {
		return ret, err
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("golang_training2").C("products")
	q := bson.M{fkey: fvalue}

	err = c.Find(q).One(&ret)

	return ret, err
}

func getAll(orderby []string, limit int, page int) ([]product, error) {
	var i []product

	session, err := mgo.Dial("localhost")
	if err != nil {
		return i, err
	}

	defer session.Close()

	c := session.DB("golang_training2").C("products")

	if limit == 0 {
		limit = 9999999
	}

	if page == 0 {
		page = 1
	}

	err = c.Find(nil).Limit(limit).Skip((page - 1) * limit).Sort(orderby...).All(&i)

	return i, err
}

func insert(i product) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB("golang_training2").C("products")
	err = c.Insert(&i)

	return
}

func update(id int64, key string, value interface{}) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB("golang_training2").C("products")
	edit := bson.M{
		key: value,
	}
	err = c.Update(bson.M{"id": id}, bson.M{"$set": edit})
	return
}

func addCategory(id int64, ca cat.ProdCat) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB("golang_training2").C("products")
	edit := bson.M{
		"$push": bson.M{
			"prod_cat": bson.M{
				"cat_id":   ca.Category_id,
				"cat_name": ca.Category_name,
			},
		},
	}
	err = c.Update(bson.M{"id": id}, bson.M{"$set": edit})
	return
}

func delCategory(id int64, cat_id int64) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB("golang_training2").C("products")
	edit := bson.M{
		"$pull": bson.M{
			"prod_cat": bson.M{
				"cat_id": cat_id,
			},
		},
	}
	err = c.Update(bson.M{"id": id}, bson.M{"$set": edit})
	return
}
