package cust

import "golang_api/db"

func findByKey(key string, value string) (customer, error) {
	var a customer

	err := db.GetOne("golang_training2", "collection", key, value, &a)

	return a, err
}
