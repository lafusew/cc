package controllers

import (
	"github.com/lafusew/cc/data"
)

type Controller struct {
	Orm data.ORM
	TransactionController TransactionsHandlers
}