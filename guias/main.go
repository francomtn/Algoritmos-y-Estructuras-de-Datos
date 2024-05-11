/*package main

import (
	"fmt"
	"guias/diccionarios"
)

func main() {
	arreglo := [5]int{3, 4, 3, 2, 1}
	fmt.Println((diccionarios.ParElementosSumanK(arreglo[:], 7)))
}*/

// arreglo ordenado := [12]int{1,2,3,5,8,11,17,21,26,30,35,56}
// arreglo desordenado := [12]int{2, 6, 5, 8, 4, 9, 11, -56, 42, 0, 7, 5}
// arreglo casi ordenado := [12]int{1,2,3,0,8,11,17,21,26,30,35,56}        [10]int{1,2,3,4,5,6,7,8,90,10}
// 1s y 0s := [9]int{1,1,1,0,0,1,0,1,0}
// pico := [12]int{0,1,2,3,4,5,7,5,4,3,2,1} >> 6      [9]int{2,3,4,5,6,7,6,5,4} >> 5      [9]int{2,3,4,5,4,3,2,1,0} >> 3
