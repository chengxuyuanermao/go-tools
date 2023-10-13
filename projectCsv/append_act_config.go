package projectCsv

import "fmt"

func Generate() {
	var a, b, c int64
	a = 1698508800
	j := 52
	for i := 1; i <= 1000; i++ {
		if i == 1 {
			fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s\n", "id", "start_time", "end_time", "draw_time", "status", "total_reward", "ticket_number_reward", "is_reward")
		}
		c = a + 60*60*24*3
		b = c - 1
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d\n", j, a, b, c, 0, 10000000, 3000, 0)
		j++
		a = c
	}
}
