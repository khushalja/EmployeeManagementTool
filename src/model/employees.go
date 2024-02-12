package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	PhoneNo int64  `json:"phoneno" bson:"phoneno"`
	Address string `json:"address" bson:"address"`
	EmailId string `json:"emailid" bson:"emailid"`
}
type Employee struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	EmployeeId     int                `json:"employeeid"`
	EmployeeName   string             `json:"employeename" bson:"employeename"`
	Contact        Contact            `json:"contact" bson:"contact"`
	JobTitle       string             `json:"jobtitle" bson:"jobtitle"`
	Department     string             `json:"department" bson:"department"`
	Salary         float64            `json:"salary" bson:"salary"`
	EmployeeType   int                `json:"employeetype" bson:"employeetype"`
	Token          string             `json:"token" bson:"token"`
	RefreshedToken string             `json:"RefreshedToken" bson:"RefreshedToken"`
}
