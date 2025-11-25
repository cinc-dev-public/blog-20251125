package testa

import (
	"blogtest/internal/selector/common"
	"blogtest/internal/testdata"
)

var TypeCheckSelector = common.SelectorGroup{
	Name: "testa type check selector",
	Selectors: []common.SelectorItem{
		{
			ID: "fc1af446-1402-46c1-a694-1579fca68731",
			Selector: func(data *testdata.TestData) *testdata.TestData {
				if data.Content == "a1" {
					return data
				}
				return nil
			},
		},
		{
			ID: "f35ca47fb-bdde-4fd7-b125-9a59b6611b5f",
			Selector: func(data *testdata.TestData) *testdata.TestData {
				if data.Content == "a2" {
					return data
				}
				return nil
			},
		},
	},
}
