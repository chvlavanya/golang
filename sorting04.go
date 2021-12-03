/* Alta3 Research | RZFeeser
   Custom Sorting - Sorting structs with custom functions   */

package main

import (
    "fmt"
    "sort"
)

// record a "Person" Name, Age
type Person struct {
    Name   string
    Age    int
}

// Custom format to show - use String(acts as tostring)
//anything lower case starting letter in code - visible to main packages
//anything uppercase starting letter in code - Visuable to other packages.
/* In Go, if you define "String()" for a struct
   the result is a custom "ToString". That is to say,
   we can control the "view" that is returned when someone tries
   to print our struct */
func (p Person) String() string {
    return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int {
    return len(a)
}
func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
    return a[i].Age < a[j].Age
}

func main() {

    // Name, Age
    people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    fmt.Println(people)

    // We defined the interface for sort.Sort
    // a set of methods for the slice type, as with ByAge, and
    // call sort.Sort.
    sort.Sort(ByAge(people))
    fmt.Println(people)

}

