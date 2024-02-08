package main

import (
	"fmt"
	"math"
)

func main() {
	//task1()
	//task2()
	task3()
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Order struct {
	customerName string
	quantity     int
	item
	total int
}

type item struct {
	productName string
	price       int
}

func task1() {
	ord := Order{customerName: "Иван", quantity: 12, item: item{productName: "Книга", price: 300}}
	ord.calc()
	fmt.Println(formatedOrder(ord))
}

func (o *Order) calc() {
	o.total = o.price * o.quantity
}

func formatedOrder(o Order) string {
	return fmt.Sprintf("Заказ от %s: %dx %s (Итого: $%d)\n", o.customerName, o.quantity, o.productName, o.total)
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func task2() {
	circle := Circle{Shape{"Круг"}, 5}
	rectangle := Rectangle{Shape{"Прямоугольник"}, 4, 6}

	fmt.Printf("%s: Площадь = %.2f\n", circle.GetName(), circle.Area())
	fmt.Printf("%s: Площадь = %.2f\n", rectangle.GetName(), rectangle.Area())

}

type Shape struct {
	Name string
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	return 0.0
}

type Circle struct {
	Shape
	radius float64
}

func (c Circle) Area() float64 {

	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	Shape
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func task3() {
	account := &BankAccount{
		owner:   "Иван",
		balance: 1000,
	}

	account.Deposit(500)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())

	account.Withdraw(200)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())

	account.Withdraw(1500)
	fmt.Printf("%s: Баланс - $%.2f\n", account.owner, account.Balance())

}

type BankAccount struct {
	owner   string
	balance float32
}

func (a *BankAccount) Balance() float32 {
	return a.balance
}

func (a *BankAccount) Deposit(i float32) {
	a.balance += i
}

func (a *BankAccount) Withdraw(i float32) {
	a.balance -= i
}
