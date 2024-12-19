package defined

import "fmt"

// Struct in the Not equal datatype saved
type Datatype struct {
	Hostname  string
	Platform  string
	Username  string
	Password  string
	StrictKey bool
}

// Datatype struct for Routers variable bind
type Inventory struct {
	Routers []Datatype
}

func BindStructField() {
	var r1 Datatype
	r1.Hostname = "router1.com"

	inv := Inventory{
		Routers: []Datatype{r1},
	}

	fmt.Println("Inventory : %+v", inv)
}
