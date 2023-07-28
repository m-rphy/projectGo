// There are executable files and reusable file

// (Always at the top)
package main // Package == Projeect == Workspace

// package name = "main" is used for an executable file

// Any other name other than main
// means we are making a reusable package

//(Always below package name)
import "fmt" 
// -> give my package main all the functionality of package fmt

// func main is needed in package main 

//(Always below the import statements)
func main() {
    fmt.Println("Hello World!")
}
