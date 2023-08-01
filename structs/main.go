package main

import "fmt"

// declare struct here
type person struct {
    firstName string
    lastName string
    contact contactInfo
}

//embedded structs
type contactInfo struct {
    email string
    zipCode int
}

func main() {
    
    //alex := person{firstName: "Alex", lastName: "Anderson"}         

    // kind of the worst
    //alex := person{"Alex","Anderson"}         
   
    // struct is initialized to the types "Zero Value"
    //var jim person 
    
    // updating struct fields with dot syntax
    jim := person {
    firstName : "Alex",
    lastName : "Anderson",
    contact: contactInfo {
            email: "jim@gmail.com",
            zipCode: 91827,
        },
    }

    // doesn't update "jim" immediately (Go is a pass by copy lang) 
    // (hint : you need to use pointers)
    // jim.updateFirstName("Jimmy")
    
    // Verbose 
    //jim_p := &jim // assign jimPointer to the address of jim
    //jim_p.updateFirstName("Jimmy") //pass the address to updateFirstName
   
    jim.updateFirstName("Jimmy")
    jim.print()
}

func (p person) print() {
    
    // just prints values
    //fmt.Println(p)

    // prints the fields as well
    fmt.Printf("%+v", p)
}

// changing the type to a pointer -> * is a type of pointer
// * in front of type is diff than * in front of value

// * in front of type is a type description

// * in front of a value is an operator that gives us access to value
// at the pointer address
func (p *person) updateFirstName(newFirstName string) {
        // passes a copy

    //p.firstName = newFirstName
    (*p).firstName = newFirstName
}
