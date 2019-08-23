package main
import "fmt"

type Artist struct{
	Name,Genre string
	Songs int
}
func newRelease(a Artist) int{
	a.Songs++
	return a.Songs
}
func main(){

	me:=Artist{Name:"znq",Songs:34,Genre:"Electro"}
	fmt.Printf("%s has a total  %d Songs\n",me.Name,me.Songs)
	fmt.Printf("%s has released  %dth Songs\n",me.Name,newRelease(me))
}

