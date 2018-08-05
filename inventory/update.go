package inventory

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

// Update changes the items in the inventory to reflect the passing of one day.
func (i *inventory) Update() {
	for index, item := range i.items {
		if item.name == sulfuras {
			continue
		}

		if item.name != agedBrie && item.name != backstagePasses {
			if item.quality > 0 {
				item.quality--
			}
		} else {
			if item.quality < 50 {
				item.quality++

				if item.name == backstagePasses {
					if item.sellIn < 11 {
						if item.quality < 50 {
							item.quality++
						}
					}
					if item.sellIn < 6 {
						if item.quality < 50 {
							item.quality++
						}
					}
				}
			}
		}

		item.sellIn--

		if item.sellIn < 0 {
			if item.name != agedBrie {
				if item.name != backstagePasses {
					if item.quality > 0 {
						item.quality--
					}
				} else {
					item.quality -= item.quality
				}
			} else {
				if item.quality < 50 {
					item.quality++
				}
			}
		}

		i.items[index] = item
	}
}
