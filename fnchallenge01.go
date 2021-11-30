/*Create a program that creates a struct called Astro. It should track the name (string), age (int), lastmission (string), and if they are needed on the next mission, isneeded (bool). Create 3 structs, and populate them with fictitious data. When you are done, print your records to the screen*/

package main
import "fmt"

type astro struct {
	name string
	age int
	mission string
	isneeded bool
}

func main() {
	ast1 := astro{"Lavanya", 20, "ISS", false}
	ast2 := astro{"Anshu", 25, "Boeing Crew Flight Test (CFT)", true}
	ast3 := astro{"Srini", 30, "SpaceX Crew-3", true}
	fmt.Println(ast1)
	fmt.Println(ast2)
	fmt.Println(ast3)
}
