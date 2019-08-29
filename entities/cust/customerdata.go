package cust

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type address struct {
	City     string `bson:"city" json:"city"`
	Province string `bson:"prov" json:"prov"`
}

type customer struct {
	Id        string    `bson:"id" json:"id"`
	Firstname string    `bson:"fname" json:"fname"`
	Lastname  string    `bson:"lname" json:"lname"`
	Email     string    `bson:"email" json:"email"`
	Home      address   `bson:"home" json:"home"`
	Shipping  address   `bson:"shipping" json:"shipping"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Password  string    `bson:"password" json:"password"`
}

type customerInput struct {
	Id        string  `json:"id"`
	Firstname string  `json:"fname"`
	Lastname  string  `json:"lname"`
	Email     string  `json:"email"`
	Home      address `json:"home"`
	Shipping  address `json:"shipping"`
	Password  string  `json:"password"`
}

func NewCustomer(id string, firstname string, lastname string, email string, homecity string, homeprovince string, shipcity string, shipprovince string) customer {
	return customer{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Home: address{
			City:     homecity,
			Province: homeprovince,
		},
		Shipping: address{
			City:     shipcity,
			Province: shipprovince,
		},
		Timestamp: time.Now(),
	}
}

func NewCustomerInput(id string, firstname string, lastname string, email string, homecity string, homeprovince string, shipcity string, shipprovince string, password string) customerInput {
	return customerInput{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Home: address{
			City:     homecity,
			Province: homeprovince,
		},
		Shipping: address{
			City:     shipcity,
			Province: shipprovince,
		},
		Password: password,
	}
}

func InsertPreparation(cust customerInput) customer {
	return customer{
		Id:        cust.Id,
		Firstname: cust.Firstname,
		Lastname:  cust.Lastname,
		Email:     cust.Email,
		Home: address{
			City:     cust.Home.City,
			Province: cust.Home.Province,
		},
		Shipping: address{
			City:     cust.Shipping.City,
			Province: cust.Shipping.Province,
		},
		Timestamp: time.Now(),
		Password:  cust.Password,
	}
}

func getOne(fkey string, fvalue string) (customer, error) {
	var ret customer

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

func getAll(orderby []string, limit int, page int) ([]customer, error) {
	var i []customer

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

func insert(i customer) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB("golang_training2").C("customers")
	err = c.Insert(&i)

	return
}

func update(i customer) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB("golang_training2").C("customers")
	edit := bson.M{
		"fname": i.Firstname,
		"lname": i.Lastname,
		"email": i.Email,
		"home": bson.M{
			"city": i.Home.City,
			"prov": i.Home.Province,
		},
		"shipping": bson.M{
			"city": i.Shipping.City,
			"prov": i.Shipping.Province,
		},
		"timestamp": time.Now(),
		"password":  i.Password,
	}
	err = c.Update(bson.M{"id": i.Id}, bson.M{"$set": edit})
	return
}
