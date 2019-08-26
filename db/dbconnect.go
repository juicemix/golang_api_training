package db

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Insert to database
func Insert(database string, collection string, i interface{}) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB(database).C(collection)
	err = c.Insert(&i)

	return
}

// Update to database
func Update(database string, collection string, fkey string, fvalue string, key string, value string) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB(database).C(collection)
	selector := bson.M{fkey: fvalue}
	upd := bson.M{"$set": bson.M{key: value, "timestamp": time.Now()}}
	err = c.Update(selector, upd)

	return
}

// Get single data
func GetOne(database string, collection string, fkey string, fvalue string, i interface{}) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB(database).C(collection)
	q := bson.M{fkey: fvalue}
	err = c.Find(q).One(&i)

	return
}

// Get many
func GetMany(database string, collection string, orderby []string, limit int, page int, fkey string, fvalue string, i interface{}) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB(database).C(collection)
	err = c.Find(bson.M{fkey: fvalue}).Limit(limit).Skip((page - 1) * limit).Sort(orderby...).One(&i)

	return
}

// Get all data
func GetAll(database string, collection string, orderby []string, limit int, page int, i interface{}) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return
	}

	defer session.Close()

	c := session.DB(database).C(collection)

	if limit == 0 {
		limit = 9999999
	}

	err = c.Find(bson.M{}).Limit(limit).Skip((page - 1) * limit).Sort(orderby...).All(&i)

	return
}
