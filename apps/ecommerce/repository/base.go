package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

const (
	ProductsfilePath = "products.json"
)

type ProductRepository struct {
	Products []ProductDAO
}

func LoadDatabase() (ProductRepository, error) {
	ret := ProductRepository{}

	content, err := ioutil.ReadFile(ProductsfilePath)
	if err != nil {
		return ret, errors.New("error opening file " + ProductsfilePath + ": " + err.Error())
	}

	err = json.Unmarshal(content, &ret.Products)
	if err != nil {
		return ret, errors.New("error unmarshalling json: " + err.Error())
	}

	return ret, nil
}
