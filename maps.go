package main

import "fmt"


func main() {

	x:= make(map[string]string)
	x["key"] = "10"
	x["vale"] = "20"
	x["name"] = "Abebe"

	for _, value := range x {
		fmt.Println(value)
	}
	// fmt.Println(x["name"])



	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen", 
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium", 
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium", 
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium", 
			"state":"solid",
		},
		"B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
		}

		if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
		}


		xx := [6]string{"a","b","c","d","e","f"}
		fmt.Println(xx[2:5])
	   
        // find the smallest values
		xy := []int{
			48,96,86,68,
			57,82,63,70,
			37,34,83,27,
			19,97, 9,17,
			8,4, 5, 55,
			}
	//    fmt.Println(xy[0])
	   var smallest int	= xy[0]	
	   for i :=0; i < len(xy); i++ {
          if smallest > xy[i] {
			 smallest = xy[i]
		  }
	   }
	   fmt.Print(smallest)
      		   
}
