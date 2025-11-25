package main

import (
	_ "blogtest/internal/selector/autoload"
	"blogtest/internal/selector/registry"
	"log"
)

func main() {
	// registry から全ての登録情報を取得
	uuidMap := make(map[string]struct{})
	for _, gs := range registry.All() {
		for _, g := range gs {
			// Selectorについて中身の取得
			for _, s := range g.Selectors {
				uuid := s.ID
				// 重複チェック
				if _, exists := uuidMap[uuid]; exists {
					log.Fatalf("duplicate UUID found: %s (group: %s)", uuid, g.Name)
				}
			}
		}
	}
}
