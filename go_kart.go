package main

import "fmt"

var productList []string

type Order struct {
	product_name string
	qty          int
	user         Users
}
type Users struct {
	User_name    string
	User_mobile  string
	User_address string
	order        []Order
}
type Go_kart interface {
	authUser()
	showProduct()
	placeOrder()
	modifyOder()
	deleteOrder()
}

var user []Users
var orderlist []Order
var prodarr [6]string = [6]string{"iron", "fridge", "cooler", "iphone", "tv", "washing machine"}

func deleteOrder(curUser Users) {
	var prodname string
	fmt.Println("enter the delete product name")
	fmt.Scanf("%s\n", &prodname)

	for j := 0; j < len(orderlist); j++ {
		if curUser.User_name == orderlist[j].user.User_name && orderlist[j].product_name == prodname {
			orderlist[j] = orderlist[len(orderlist)-1]
			orderlist = orderlist[:len(orderlist)-1]
			fmt.Printf("%v", orderlist[j])
			fmt.Printf("username:%s\nproductname:%s\nqty:%d\n", orderlist[j].user.User_name,
				orderlist[j].product_name, orderlist[j].qty)
			break
		}
	}
}

func modifyOder(curUser Users) {
	var qty int
	var modname string
	fmt.Printf("\nEnter the Product name to modify")
	fmt.Scanf("%s\n", &modname)
	var flag bool = false
	for i := 0; i < 6; i++ {
		if modname == prodarr[i] {
			fmt.Printf("\nEnter the quantity to modify\n")
			fmt.Scanf("%d\n", &qty)
			flag = true
		}
	}
	if flag {
		for j := 0; j < len(orderlist); j++ {
			if curUser.User_name == orderlist[j].user.User_name && orderlist[j].product_name == modname {
				orderlist[j].qty = qty
				fmt.Printf("username:%s\nproductname:%s\nqty:%d\n", orderlist[j].user.User_name,
					orderlist[j].product_name, orderlist[j].qty)
				break
			}
		}

	}
}

func placeOrder(curUser Users) []Order {
	var prodName string
	var qty int
	fmt.Println("\nEnter the Product Name")
	fmt.Scanf("%s\n", &prodName)
	fmt.Println("\nEnter the Quantity")
	fmt.Scanf("%d\n", &qty)
	for i := 0; i < 6; i++ {
		if prodarr[i] == prodName {
			var orderlist1 []Order
			var order1 Order = Order{product_name: prodName, qty: qty, user: curUser}
			orderlist1 = append(orderlist1, order1)

			return orderlist1
		}

	}
	return Order{}
}
func showProduct() {
	fmt.Printf("\nName of products in Go_kart\n")
	for _, v := range prodarr {
		fmt.Printf("%s\n", v)
	}
}

func authUser(mobile string) Users {
	for i := 0; i < len(user); i++ {
		if user[i].User_mobile == mobile {
			return user[i]
		}
	}
	var newUser Users = createUser(mobile)
	user = append(user, newUser)
	return newUser
}
func createUser(mobile string) Users {
	var name, address string
	fmt.Println("Enter the name")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Enter the address")
	fmt.Scanf("%s\n", &address)
	newUser := Users{User_name: name, User_mobile: mobile, User_address: address}
	return newUser
}

/*func printOrder(curUser Users) {
	for i := 0; i < len(orderlist); i++ {
		if orderlist[i] == curUser.mobile {
			fmt.Printf("Print order:%v\n", curUser)
		}
	}

}*/
func main() {

	var mobile string
	fmt.Println("\nEnter the mobile number")
	fmt.Scanf("%s\n", &mobile)
	firstUser := Users{User_name: "harsh", User_mobile: "9693880525", User_address: "E/12"}

	fmt.Printf("%v", firstUser)

	user = append(user, firstUser)
	var curUser Users = authUser(mobile)
	fmt.Printf("%v", user)
	//showProduct()

	/*var orderplace Order = placeOrder(curUser)

	orderlist = append(orderlist, orderplace)

	fmt.Printf("Username:%s\nProductName:%s\nProductQty:%d\n",
		orderplace.user.User_name, orderplace.product_name, orderplace.qty)*/
	//modifyOder(curUser)
	//deleteOrder(curUser)

	var op int

	for true {
		fmt.Printf("\n")
		fmt.Println("Enter 1 to Get Products")
		fmt.Println("Enter 2 to Place Orders")
		fmt.Println("Enter 3 to Modify Products")
		fmt.Println("Enter 4 to Cancel Products")
		fmt.Println("Enter 5 to Print Products")
		fmt.Println("Enter 6 to Print Products ordered from me")
		fmt.Println("Enter 7 to Login to another account")
		fmt.Scanln(&op)
		if op == 1 {
			showProduct()
		} else if op == 2 {
			var orderplace []Order = placeOrder(curUser)
               
			orderlist = append(orderlist, orderplace)

			fmt.Printf("Username:%s\nProductName:%s\nProductQty:%d\n",
				orderplace.user.User_name, orderplace.product_name, orderplace.qty)
		} else if op == 3 {
			modifyOder(curUser)
		} else if op == 4 {
			deleteOrder(curUser)
		} else if op == 6 {
			break
		}
	}

}
