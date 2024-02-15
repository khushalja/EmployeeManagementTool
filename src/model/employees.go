package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	PhoneNo int64  `json:"phoneno" bson:"phoneno"`
	Address string `json:"address" bson:"address"`
	EmailId string `json:"emailid" bson:"emailid"`
}
type Employee struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	//min and max value works only with strintg or array lengths and not for numeric data types.
	EmployeeId int64 `json:"employeeid" validate:"required"`
	// we need to use custom validations for that.
	EmployeeName   string  `json:"employeename" validate:"required,min=4,max=20"`
	Contact        Contact `json:"contact" bson:"contact" validate:"required"`
	JobTitle       string  `json:"jobtitle" bson:"jobtitle" validate:"required"`
	Department     string  `json:"department" bson:"department" validate:"required"`
	Salary         float64 `json:"salary" bson:"salary" validate:"required"`
	EmployeeType   int     `json:"employeetype" bson:"employeetype" validate:"required"`
	Token          string  `json:"token" bson:"token"`
	RefreshedToken string  `json:"refreshedtoken" bson:"refreshedtoken"`
	Password       string  `json:"password" validate:"required"`
}
