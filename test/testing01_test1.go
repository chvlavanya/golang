/* Alta3 Research | RZFeeser
   CHALLENGE 01 - testing an empty string 
   go test -v
=== RUN   TestHello
--- PASS: TestHello (0.00s)
PASS
ok      github.com/alta3/yamlgo/test    0.002s

*/

package main

import (
    "testing"   // used for testing
    "regexp"    // regular expression library
)

func TestHello(t *testing.T) {
    testName := "RZFeeser"
    desiredResult := regexp.MustCompile("Hello there "+testName)

    result, err := Hello("RZFeeser")
    
    if !desiredResult.MatchString(result) || err != nil {
        t.Fatalf("Wanted %v, but got, %v, or got %v", desiredResult, result, err)
    }
    
    result, err = Hello("")
    
    // we are testing the ability to produce an error when entering "ENTER" instead of a name
    // this should return an error for err
    if result != "" || err == nil {
        t.Fatalf("Test failed to produce correct failure for empty string. Got, %v", result)
    }
}

