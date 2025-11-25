package common

import (
	"blogtest/internal/testdata"
	"log"
)

type SelectorGroup struct {
	Name      string
	Selectors []SelectorItem
}

type SelectorItem struct {
	ID       string // uuid v4 をハードコードで設定する
	Selector func(datas *testdata.TestData) *testdata.TestData
}

func (sg SelectorGroup) Select(datas *testdata.TestData) *testdata.TestData {
	for _, selector := range sg.Selectors {
		result := selector.Selector(datas)
		if result != nil {
			log.Printf("matched: %s", selector.ID)
			return result
		}
	}
	return nil
}
