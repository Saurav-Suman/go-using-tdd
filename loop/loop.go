package loop

import "fmt"

func loop(data string, count int) string {
	returnstr := ""
	for i := 0; i < count; i++ {
		returnstr += data
	}
	fmt.Println(returnstr)
	return returnstr

}
