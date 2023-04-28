//this is my learning-project for learning go
package main
import("fmt")

const PI float32 = 3.14
func main(){
	
	string student1 := "student1" //string
	var student2 string = "student2" //stirng
	var count int = 2 //int
    courses := [2] string {"math", "physics"} // arr = [length]type{values}
	var grades =[...] string {"pass","fail"} 
	var mySlice = [] string {"slice1"} //slice is like array but length can extend
	mySlice = append(mySlice,"slice2") //append (alicename, param1...)
	fmt.Println("Welcom!")
	ftm.Println("student name is: " student1)
	ftm.Println("student name is: " student2)
	score1 := 9
	acore2 := 13
	//conditions
	if( score1 < 10){
       ftm.Println("you are failed")
	}
	else{
		ftm.Println("you are pass")
	}
}