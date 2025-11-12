package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, val_exists := units[unit]
	if !val_exists {
		return false
	}
	_, item_already_present := bill[item]
	if item_already_present {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, item_exists := bill[item]
	_, unit_exists := units[unit]
	if !item_exists || !unit_exists {
		return false
	}
	current_amount := bill[item]
	amount_to_remove := units[unit]
	if current_amount-amount_to_remove < 0 {
		return false
	} else if current_amount-amount_to_remove == 0 {
		delete(bill, item)
	} else {
		bill[item] -= amount_to_remove
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amount, exists := bill[item]
	if !exists || amount == 0 {
		return 0, false
	}
	return amount, true
}
