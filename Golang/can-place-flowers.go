package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	for index := 0; index < len(flowerbed); index++ {

		if flowerbed[index] == 0 {
			var previousFlower int
			var nextFlower int

			if index == 0 {
				previousFlower = 0
			} else {
				previousFlower = flowerbed[index-1]
			}

			if index == len(flowerbed)-1 {
				nextFlower = 0
			} else {
				nextFlower = flowerbed[index+1]
			}

			if previousFlower == 0 && nextFlower == 0 {
				flowerbed[index] = 1
				n -= 1
			}

		}

	}

	return n <= 0
}

func main() {
	print(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
}
