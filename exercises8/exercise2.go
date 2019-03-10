package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of people", people)
	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for idx, val := range v.Sayings {
			fmt.Println("\t\t", idx, val)
		}

	}

}

