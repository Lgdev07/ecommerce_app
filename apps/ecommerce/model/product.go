package model

import (
	"github.com/asaskevich/govalidator"
)

type Product struct {
	Id            int  `json:"id" valid:"required,int"`
	Quantity      int  `json:"quantity" valid:"required,int"`
	UnitAmount    int  `json:"unit_amount" valid:"int,optional"`
	TotalAmount   int  `json:"total_amount" valid:"int,optional"`
	DiscountGiven int  `json:"discount" valid:"int,optional"`
	IsGift        bool `json:"is_gift" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (product *Product) IsValid() error {
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return err
	}
	return nil
}
