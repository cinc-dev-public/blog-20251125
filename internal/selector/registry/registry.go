package registry

import "blogtest/internal/selector/common"

var selectorGroups [][]common.SelectorGroup

func Register(sg []common.SelectorGroup) { selectorGroups = append(selectorGroups, sg) }
func All() [][]common.SelectorGroup      { return selectorGroups }
