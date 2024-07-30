package container

import (
	"github.com/zeppelinmc/zeppelin/net/slot"
	"github.com/zeppelinmc/zeppelin/server/item"
	"github.com/zeppelinmc/zeppelin/server/registry"
)

// A container that holds items
type Container []item.Item

// encodes the container to network slots
func (c Container) Network(size int) []slot.Slot {
	s := make([]slot.Slot, size)
	for _, item := range c {
		id, ok := registry.Item.Lookup(item.Id)
		if !ok {
			continue
		}
		s[item.Slot.Network()] = slot.Slot{
			ItemCount: item.Count,
			ItemId:    id,
		}
	}

	return s
}

// adds the item to the container and replaces the existing one if found, and returns if the operation was successful
func (c *Container) SetSlot(item item.Item) {
	for i := range *c {
		if (*c)[i].Slot == item.Slot {
			(*c)[i] = item
			return
		}
	}
	*c = append(*c, item)
}

// finds the item at the specified data slot
func (c Container) Slot(slot item.DataSlot) (item.Item, bool) {
	for _, item := range c {
		if item.Slot == slot {
			return item, true
		}
	}
	return item.Item{}, false
}