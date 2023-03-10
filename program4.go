package main

import "fmt"

type Address struct {
	Name        string
	LastName    string
	city        string
	phonenumber int
	pincode     int
	state       string
}

type AddressSlice struct {
	slice []Address
}
type AddressBok struct {
	AddressMap map[string]Address
}

func (MyObj *AddressSlice) addressdetails(address Address) {
	MyObj.slice = append(MyObj.slice, address)
}

func (MyMapObj *AddressBok) addToMap(address Address) {
	MyMapObj.AddressMap[address.Name] = address
}

func main() {
	addr := Address{}
	fmt.Println("Enter  the FirstName : ")
	fmt.Scan(&addr.Name)
	fmt.Println("enter the LastName")
	fmt.Scan(&addr.LastName)
	fmt.Println("Enter the city Name : ")
	fmt.Scan(&addr.city)
	fmt.Println("Enter the phonenumber : ")
	fmt.Scan(&addr.phonenumber)
	fmt.Println("Enter the Pincode : ")
	fmt.Scan(&addr.pincode)
	fmt.Println("Enter the state Name : ")
	fmt.Scan(&addr.state)
	address1 := Address{"Goutham", "Raj", "Banglore", 9876543211, 560020, "Karnataka"}
	address2 := Address{"Goutham", "Ram", "Banglore", 9876544211, 560040, "Karnataka"}
	address3 := Address{"Goutham", "Rao", "Banglore", 9876643211, 560050, "Karnataka"}

	AddressSlice := AddressSlice{}
	AddressSlice.addressdetails(address1)
	AddressSlice.addressdetails(address2)
	AddressSlice.addressdetails(address3)

	AddressSlice.addressdetails(addr)
	for key, value := range AddressSlice.slice {
		fmt.Println(key)
		fmt.Println(value.Name, value.LastName, value.city, value.phonenumber, value.pincode, value.state)
	}

}
