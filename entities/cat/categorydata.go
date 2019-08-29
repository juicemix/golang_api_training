package cat

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type category struct {
	Category_id   int64     `bson:"cat_id" json:"cat_id"`
	Category_name string    `bson:"cat_name" json:"cat_name"`
	IsActive      bool      `bson:"active" json:"active"`
	Timestamp     time.Time `bson:"timestamp" json:"timestamp"`
}

type ProdCat struct {
	Category_id   int64  `bson:"cat_id" json:"cat_id"`
	Category_name string `bson:"cat_name" json:"cat_name"`
}

func NewCategory(id int64, name string, isActive bool) category {
	return category{
		Category_id:   id,
		Category_name: name,
		IsActive:      isActive,
		Timestamp:     time.Now(),
	}
}

func getOne(fkey string, fvalue string) (category, error) {
	var ret category

	session, err := mgo.Dial("localhost")
	if err != nil {
		return ret, err
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("golang_training2").C("customers")
	q := bson.M{fkey: fvalue}

	err = c.Find(q).One(&ret)

	return ret, err
}

func getAll(orderby []string, limit int, page int) ([]category, error) {
	var i []category

	session, err := mgo.Dial("localhost")
	if err != nil {
		return i, err
	}

	defer session.Close()

	c := session.DB("golang_training2").C("customers")

	if limit == 0 {
		limit = 9999999
	}

	if page == 0 {
		page = 1
	}

	err = c.Find(nil).Limit(limit).Skip((page - 1) * limit).Sort(orderby...).All(&i)

	return i, err
}
