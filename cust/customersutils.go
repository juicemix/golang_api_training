package cust

import "golang_api/db"

func findByKey(key string, value string) (customer, error) {
	var a customer

	err := db.GetOne("golang_training2", "customers", key, value, &a)

	return a, err
}
