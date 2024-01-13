package main

import "fmt"

func main() {
	names := [...]string{"Hari", "Budi", "Yanto", "Nugra", "Yaya", "Lili"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice1 := days[5:]
	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "New Holiday")
	fmt.Println(days)
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	daySlice2[0] = "Old Saturday"
	fmt.Println(days)
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	/* make slice */
	newSlice := make([]string, 2, 5) // 2->Length, 5->Capacity
	newSlice[0] = "Hari"
	newSlice[1] = "Budi"
	// newSlice[2] = "Yaya" // error, must use append
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Kaka")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Jaja"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	/* copy slice */
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
