package main

import utils "example/project/mypackage"

type customer struct {
	name    string
	address string
	balance float64
}

func getCustomerInfo(c customer) {
	utils.F("%s owes us %.2f\n", c.name, c.balance)
}

func newCustomerAddress(c *customer, address string) {
	c.address = address
}

type rectangle struct {
	lenght, height float64
}

func (r rectangle) Area() float64 {
	return r.lenght * r.height
}

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	utils.F("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.phone)
}

func main() {
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 Main Street"
	tS.balance = 234.54

	getCustomerInfo(tS)
	newCustomerAddress(&tS, "123 South Street")
	utils.Pl("Address: ", tS.address)

	sS := customer{"Sally Smith", "123 Main", 0.0}
	getCustomerInfo(sS)

	rect1 := rectangle{10.0, 15.0}
	utils.Pl("Rect Area: ", rect1.Area())

	contact1 := contact{"James", "Wang", "555-1223"}
	business1 := business{"ABC Plumbing", "223 North Street", contact1}
	business1.info()

}
