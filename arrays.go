package main

//import other packages

import(
    "fmt"
)

func main(){
    
    //Arrays
    var firstNames = [3] string {"John", "April", "Wes"}  
    fmt.Println(firstNames[1])  
    
    // Infer length
    var nummies = [] int {1,2,3,4,5}
    fmt.Println(nummies[0])
    
    //shorthand
    lastNames := [2] string {"Elder", "Smith"}
    fmt.Println(lastNames)
    
    //shorthand with inferred
    ourNames := []string {"Elder","Browne"}
    fmt.Println(ourNames)
    
    // Change item in array
    ourNames[1] = "Smith"
    fmt.Println(ourNames)
    
    // default Assignment
    var ourNummies = [5] int {1,2,3}
    fmt.Println(ourNummies)
    
    //Assign in certain positions
    var moreNummies = [10] int {0:41, 4:99, 9:69}
    fmt.Println(moreNummies)
    
    // Slices
    var toppings = [5] string {"Pepperoni","Onion", "Cheese", "Supreme", "Bacon"}
    toppingSlice := toppings[0:2]
    fmt.Println(toppingSlice)
    
    toppingSlice[1] = "Peppers"
    fmt.Println(toppingSlice)
    
    //add item to slice
    toppingSlice = append(toppingSlice, "Mozzarela")
    fmt.Println(toppingSlice)
    
     //two slices together
     otherSlice := toppings[3:4]
     fmt.Println(otherSlice)
     
     otherSlice = append(otherSlice, toppingSlice...)
     fmt.Println(otherSlice)
    
}
