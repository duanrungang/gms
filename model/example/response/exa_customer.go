package response

import (
	"gms/model/example"
)

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
