package resource

import (
	"os"

	"github.com/chronojam/aws-pricing-api/types/schema"
	"github.com/olekukonko/tablewriter"
)

func GetManageBlockChain() {
	mgmtblockchain := &schema.AmazonManagedBlockchain{}
	err := mgmtblockchain.Refresh()
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Description", "USD", "Unit"})
	table.SetRowLine(true)
	data := []*schema.AmazonManagedBlockchain_Product{}
	for _, price := range mgmtblockchain.Products {
		data = append(data, price)
	}

	for _, p := range data {
		for _, term := range mgmtblockchain.Terms {
			if term.Sku == p.Sku {
				for _, priceData := range term.PriceDimensions {
					x := []string{}
					v := append(x, priceData.Description, priceData.PricePerUnit.USD, priceData.Unit)
					table.Append(v)
				}
			}
		}
	}

	table.Render()
}
