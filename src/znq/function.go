package znq

func Localtion(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "sanming", "xiamen":
		region, continent = "fujian", "china"
	case "hangzhou", "ningbo":
		region, continent = "zhejiang", "china"
	default:
		region, continent = "unknow", "unknow"
	}
	return region, continent
}

//functions can return multiple “result parameters”, not just a single value.
//return multiple param
func Localtion1(name, city string) (region, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		region, continent = "California", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}
