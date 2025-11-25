package testb

import (
	"blogtest/internal/selector/common"
	"blogtest/internal/testdata"
)

var TypeCheckSelector = common.SelectorGroup{
	Name: "testb type check selector",
	Selectors: []common.SelectorItem{
		{
			ID: "355e6f03-00e5-41a7-8f26-94152fd1eabc",
			Selector: func(data *testdata.TestData) *testdata.TestData {
				if data.Content == "b1" {
					return data
				}
				return nil
			},
		},
	},
}
