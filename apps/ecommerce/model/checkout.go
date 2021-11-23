package model

import (
	"github.com/asaskevich/govalidator"
)

type Checkout struct {
	TotalAmount             int       `json:"total_amount" valid:"int,optional"`
	TotalAmountWithDiscount int       `json:"total_amount_with_discount" valid:"int,optional"`
	TotalDiscount           int       `json:"total_discount" valid:"int,optional"`
	Products                []Product `json:"products" valid:"required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (checkout *Checkout) IsValid() error {
	_, err := govalidator.ValidateStruct(checkout)
	if err != nil {
		return err
	}
	return nil
}
