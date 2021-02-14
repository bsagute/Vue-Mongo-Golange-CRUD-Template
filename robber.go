package robber

//GetMaxRobAmount - To calculate max amount robber can rob withoud caught
func GetMaxRobAmount(houseOfAmounts []int) int {
	houseLength := len(houseOfAmounts)
	//If house length is 0 or 2 , cant rob , if len ==2 both are adjacent to each other
	if houseLength == 0 || houseLength == 2 {
		return 0
	}
	//No one is adjacent ,so robber can directly Rob house
	if houseLength == 1 {
		return houseOfAmounts[0]
	}

	totalEvenHouses := 0
	totalOddHouses := 0
	for i := 0; i < houseLength; i++ {
		if i%2 == 0 {
			totalEvenHouses = totalEvenHouses + houseOfAmounts[i]
		} else {
			totalOddHouses = totalOddHouses + houseOfAmounts[i]
		}

	}
	// inputHouseList := []int{2, 3, 8, 4,2}
	// inputHouseList := []int{2, 3, 8, 4}

	if houseLength%2 == 1 {
		totalEvenHouses = totalEvenHouses - GetMinimumAmout(houseOfAmounts[0], houseOfAmounts[houseLength-1])

	}

	return GetMaximumAmout(totalEvenHouses, totalOddHouses)
}

//HANDEL if house condition
//Get the Maximum fo two houses
func GetMaximumAmout(firstHouse, secondHouse int) int {
	if firstHouse > secondHouse {
		return firstHouse

	}
	return secondHouse

}

//Min - Function to return min value
func GetMinimumAmout(firstHouse, secondHouse int) int {
	if firstHouse > secondHouse {
		return secondHouse

	}
	return firstHouse

}
