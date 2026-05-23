package amember

import (
	"net/url"
	"strconv"
)

type productRow struct {
	Id     int              `json:"product_id"`
	Nested productNestedVal `json:"nested"`
}

type productNestedVal struct {
	Map []productToCategoryNested `json:"product-product-category,flow"`
}

type productToCategoryNested struct {
	CategoryId string `json:"product_category_id"`
}

func (c *Client) GetProductIdsForCategory(categoryId int) ([]int, error) {
	query := url.Values{}
	query.Set("_nested[]", "product-product-category")
	products, err := allPages[productRow](c, "/api/products", query)
	if err != nil {
		return nil, err
	}

	ids := make([]int, 0)
	for _, p := range products {
		for _, c := range p.Nested.Map {
			cint, err := strconv.Atoi(c.CategoryId)
			if err != nil {
				return nil, err
			}
			if cint == categoryId {
				ids = append(ids, p.Id)
				break
			}
		}
	}

	return ids, nil
}
