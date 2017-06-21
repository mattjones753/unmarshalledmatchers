package ExpandedUnmarshalledMatchers

import (
	"github.com/onsi/gomega/types"
	"fmt"
)


func MatchUnorderedJSON(json interface{}, keys ...KeyExclusions) types.GomegaMatcher {
	deepMatcher := UnmarshalledDeepMatcher{
		Ordered: false,
		Subset:  false,
	}

	if len(keys) > 0{
		if len(keys) > 1 {
			fmt.Errorf("Only 1 key exclusion set is currently supported")
		} else if keys[0].IsOrdered(){
			deepMatcher.InvertOrderingKeys = keys[0].GetMap()
		} else {
			fmt.Errorf("You are trying to set unordered list keys for unordered JSON")
		}
	}


	return &ExpandedJsonMatcher{
		JSONToMatch: json,
		DeepMatcher: deepMatcher,
	}
}

func MatchOrderedJSON(json interface{}, keys ...KeyExclusions) types.GomegaMatcher {
	deepMatcher := UnmarshalledDeepMatcher{
		Ordered: true,
		Subset:  false,
	}

	if len(keys) > 0{
		if len(keys) > 1 {
			fmt.Errorf("Only 1 key exclusion set is currently supported")
		} else if keys[0].IsOrdered(){
			fmt.Errorf("You are trying to set ordered list keys for ordered JSON")
		} else {
			deepMatcher.InvertOrderingKeys = keys[0].GetMap()
		}
	}


	return &ExpandedJsonMatcher{
		JSONToMatch: json,
		DeepMatcher: deepMatcher,
	}
}

func ContainUnorderedJSON(json interface{}, keys ...KeyExclusions) types.GomegaMatcher {
	deepMatcher := UnmarshalledDeepMatcher{
		Ordered: false,
		Subset:  true,
	}

	if len(keys) > 0{
		if len(keys) > 1 {
			fmt.Errorf("Only 1 key exclusion set is currently supported")
		} else if keys[0].IsOrdered(){
			deepMatcher.InvertOrderingKeys = keys[0].GetMap()
		} else {
			fmt.Errorf("You are trying to set unordered list keys for unordered JSON")
		}
	}


	return &ExpandedJsonMatcher{
		JSONToMatch: json,
		DeepMatcher: deepMatcher,
	}
}

func ContainOrderedJSON(json interface{}, keys ...KeyExclusions) types.GomegaMatcher {
	deepMatcher := UnmarshalledDeepMatcher{
		Ordered: true,
		Subset:  true,
	}

	if len(keys) > 0{
		if len(keys) > 1 {
			fmt.Errorf("Only 1 key exclusion set is currently supported")
		} else if keys[0].IsOrdered(){
			fmt.Errorf("You are trying to set ordered list keys for ordered JSON")
		} else {
			deepMatcher.InvertOrderingKeys = keys[0].GetMap()
		}
	}


	return &ExpandedJsonMatcher{
		JSONToMatch: json,
		DeepMatcher: deepMatcher,
	}
}

type OrderedKeys struct {
	val map[string]bool
}

func NewOrderedKeys() OrderedKeys {
	return OrderedKeys{
		val: make(map[string]bool),
	}
}

func (k OrderedKeys) IsOrdered() bool {
	return true;
}

func (k OrderedKeys) GetMap() map[string]bool {
	return k.val;
}

type UnorderedKeys struct {
	val map[string]bool
}

func NewUnorderedKeys() UnorderedKeys {
	return UnorderedKeys{
		val: make(map[string]bool),
	}
}

func (k UnorderedKeys) IsOrdered() bool {
	return false;
}

func (k UnorderedKeys) GetMap() map[string]bool {
	return k.val;
}

type KeyExclusions interface {
	IsOrdered() bool
	GetMap() map[string]bool
}

func WithOrderedListKeys(keys ...string) OrderedKeys{
	ok := NewOrderedKeys()

	for _, v := range keys {
		ok.val[v] = true
	}

	return ok
}

func WithUnorderedListKeys(keys ...string) UnorderedKeys{
	uk := NewUnorderedKeys()

	for _, v := range keys {
		uk.val[v] = true
	}

	return uk
}