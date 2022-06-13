/* comments //
 */
package main

import "fmt"

const A string = "LOL"

func main() {
	//string
	var a string = "Anuj"
	fmt.Println(a)
	fmt.Printf("Vairable is of type: %T \n", a)

	//Boolean
	var b bool = false
	fmt.Println(b)
	fmt.Printf("Vairable is of type: %T \n", b)

	//int -> numbers
	//flaot -> number swith point

	//No Var type
	c := 300 //change 300 to 300.0 wont show error      :=this is called walerus,cant use it outside method/func
	fmt.Println(c)

	fmt.Println(A)
	fmt.Printf("Vairable is of type: %T \n", A)

}
