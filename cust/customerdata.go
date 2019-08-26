package cust

import "time"

type address struct {
	City     string `json:"city"`
	Province string `json:"prov"`
}

type customer struct {
	Id        string    `json:"id"`
	Firstname string    `json:"fname"`
	Lastname  string    `json:"lname"`
	Email     string    `json:"email"`
	Home      address   `json:"home"`
	Shipping  address   `json:"shipping"`
	Timestamp time.Time `json:"timestamp"`
	Password  string    `json:"password"`
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
