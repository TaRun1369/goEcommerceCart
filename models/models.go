package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id",bson:"_id"`
	First_Name      *string            `json:"first_name" 	validate:"required,min=2,max=100"`
	Last_Name       *string            `json:"last_name" 	validate:"required,min=2,max=100"`
	Password        *string            `json:"password" 	validate:"required,min=6"`
	Email           *string            `json:"email" 		validate:"required,email"`
	Phone           *string            `json:"phone" 		validate:"required,min=10,max=10"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart",bson:"usercart"`
	Address_Details []Address          `json:"address",bson:"address"`
	Order_Status    []Order            `json:"orders",bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"_id",bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint64            `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `json:"_id",bson:"_id"`
	Product_Name *string            `json:"product_name",bson:"product_name"`
	Price        uint64             `json:"price",bson:"price"`
	Rating       *uint64            `json:"rating",bson:"rating"`
	Image        *string            `json:"image",bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `json:"_id",bson:"_id"`
	House      *string            `json:"house",bson:"house"`
	Street     *string            `json:"street",bson:"street"`
	City       *string            `json:"city",bson:"city"`
	Pincode    *string            `json:"pincode",bson:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID `json:"_id",bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list",bson:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at",bson:"ordered_at"`
	Price          uint               `json:"total_price",bson:"total_price"`
	Discount       *uint              `json:"discount",bson:"discount"`
	Payment_Method Payment            `json:"payment_method",bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
