package handlers

import "Expenses_REST-API/types"

//go:generate ../autogen -dbtype=MongoDB -op=Delete

type DeletedExpense struct{
	*types.Expense
}