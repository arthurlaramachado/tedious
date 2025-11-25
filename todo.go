package main

import (
	"fmt"
)

func msain() {
	fmt.Println("hello, go!")
}

// Creates a new item on the todo list
func createItem() {

}

/*
Change item state, should handle:
-> state == unmarked THEN mark in progress and start time counter, show task in yellow to indicate in progress
-> state == in progress THEN mark done and finish time counter, show time to task completeness and line in green
-> state == marked THEN unmark and clean state data, show line in terminal normal color
*/
func updateItemState() {

}

/*
Should completely delete an item
The item is only deleted when text.length == 0
*/
func deleteItem() {

}

/*
Updates item text
*/
func updateItem() {

}

/*
Shows all items from a todo list
*/
func getItems() {

}
