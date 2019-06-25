package types

type Expense struct{
	Id 				int 		`json:"id" bson:"id"`
	Description 	string 		`json:"description" bson:"description"`
	Type 			string 		`json:"type" json:"type"`
	Amount 			float64 	`json:"amount" bson:"amount"`
	CreatedOn 		string 	`json:"created_on" bson:"created_on"`
	UpdatedOn 		string 	`json:"updated_on" bson:"updated_on"`
}

type Expenses []Expense




