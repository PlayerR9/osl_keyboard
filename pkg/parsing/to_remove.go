package parsing

import (
	gr "github.com/PlayerR9/SlParser/grammar"
	prx "github.com/PlayerR9/SlParser/parser"
)

type ItemSet[T gr.TokenTyper] struct {
	rules      []*prx.Rule[T]
	item_table map[T][]*prx.Item[T]
}

func NewItemSet[T gr.TokenTyper]() *ItemSet[T] {
	return &ItemSet[T]{
		item_table: make(map[T][]*prx.Item[T]),
	}
}

func (is *ItemSet[T]) AddRule(lhs T, rhss ...T) (*prx.Rule[T], error) {
	rule, err := prx.NewRule(lhs, rhss...)
	if err != nil {
		return nil, err
	}

	is.rules = append(is.rules, rule)

	return rule, nil
}
