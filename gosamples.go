package main
import (
        "fmt"
        "math/rand"
        "time"
        "strings"
        )

func randomic(){
   rand.Seed(time.Now().UnixNano())
   ranGo := rand.Intn(32) 
   joint := []string{}
   charset := "aeioucdfgkl mnqstv"
   for i := 0; i < ranGo; i++{  
       c := charset[rand.Intn(len(charset))]
    strinGo := string(c)
        joint = append(joint, strinGo)
        }
        fmt.Println(strings.Join(joint, ""))
}

func main(){
    var tries int
    fmt.Print("ENTER NUMBER OF STRINGS, PLEASE: ")
    fmt.Scanln(&tries)
    for j := 0; j < tries; j++ {
        randomic()
        time.Sleep(1 * time.Second)
        }
}
