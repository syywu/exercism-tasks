package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option2 < option1{
		option1 = option2
	}
	return option1 + " " + "is clearly the better choice."
	

}

// p * x = y

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if int(age) < 3{
		return 80/float64(100) * originalPrice
	} else if int(age) >= 10{
		return 50/float64(100) * originalPrice
	} else{
		return 70/float64(100) * originalPrice
	}
	
}
