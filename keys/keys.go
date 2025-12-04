package main

import "fmt"

func main() {
	p := Player{}
	_ = p.FoundKey(Gold)
	_ = p.FoundKey(Silver)
	_ = p.FoundKey(Bronze)
	fmt.Println("Player's keys:", p.Keys)
}

type Key byte

type Player struct {
	Keys []Key
}

const (
	Gold Key = iota + 1
	Silver
	Bronze
)

func (p *Player) FoundKey(k Key) error {
	if k != Gold && k != Silver && k != Bronze {
		return fmt.Errorf("invalid key")
	}
	for _, existing := range p.Keys {
		if existing == k {
			return nil
		}
	}

	p.Keys = append(p.Keys, k)
	return nil

}
