package main

import (
	"Ayoconnect/Server/btree"
	"Ayoconnect/Server/robber"
	"fmt"
)

//Run below code to get the Output of a both questions

func main() {
	//Prepared the list from input list
	inputHouseList := []int{2, 3, 2, -2, 0, 4}
	houseWithNonPositiveBalance := []int{}
	//Remove nonzero and negative numbers
	for _, house := range inputHouseList {
		if house > 0 {
			houseWithNonPositiveBalance = append(houseWithNonPositiveBalance, house)
		}
	}
	stmt := `Maximum amount of money professional robber can rob tonight
	        without alerting the police: `
	fmt.Println(stmt, robber.GetMaxRobAmount(houseWithNonPositiveBalance))
	var treePointer btree.Tree
	keyList := []int{10, 20, 30, 40, 60, 50}
	fmt.Println("", keyList)
	treePointer.InsertKey(10)
	treePointer.InsertKey(20)
	treePointer.InsertKey(30)
	treePointer.InsertKey(40)
	treePointer.InsertKey(50)
	treePointer.InsertKey(60)
	btree.PrintInOrder(&treePointer)
	btree.PrintPreOrder(&treePointer)
	btree.PrintPostOrder(&treePointer)
}
