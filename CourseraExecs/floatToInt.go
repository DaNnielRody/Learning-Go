package CourseraExecs

import "fmt"

func main() {

	var floatToInteger float64

	fmt.Println("Which float to integer?")

	_, err := fmt.Scan(&floatToInteger)

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	turnIntoInt := int64(floatToInteger)

	fmt.Println("Converted integer:", turnIntoInt)
}
