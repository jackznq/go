package main
import "fmt"

func main(){

	// fmt.Println(add(5,6))
	region,continent:=localtion("znq")
	fmt.Printf("znq live in %s,%s",region,continent)
	region1,continent1:=localtion1("znq","LA")
	fmt.Printf("znq live in %s,%s",region1,continent1)
}

func add(x,y int) int{
	return x+y
}

func localtion(city string)(string,string){
	var region string
	var continent string
	switch city{
	case "sanming","xiamen":
		region,continent = "fujian","china"
	case "hangzhou","ningbo":
		region,continent = "zhejiang","china"
	default :
		region,continent = "unknow","unknow"
	}
	return region,continent
}
//functions can return multiple “result parameters”, not just a single value.
func localtion1(name, city string) (region, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}
	return
}