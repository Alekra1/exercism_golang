package partyrobot

import "fmt"

func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + fmt.Sprintf("%d", age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return Welcome(name) + "\nYou have been assigned to table " + fmt.Sprintf("%03d", table) +
		". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) +
		" meters from here.\n" + "You will be sitting next to " + neighbor + "."
}
