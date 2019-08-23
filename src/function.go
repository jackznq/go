package main
import "fmt"

func main(){

	// fmt.Println(add(5,6))
	region,continent:=localtion("sanming")
	fmt.Println("znq live in %s,%s",region,continent)
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