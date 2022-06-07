package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Consumer struct {
	ID       int
	Name     string
	TableNum int
	Order    []string
	Bill     float64
	Active   bool
}

func (c *Consumer) ShowBillList(menuShorted []string, availableMenu map[string]float64) {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\tNOT PEMBAYARAN")
	fmt.Printf("No.\tMakanan\t\t\tHarga\n")
	fmt.Println(strings.Repeat("=", 40))
	var count int = 0
	for _, menu := range c.Order {
		menuInt, _ := strconv.Atoi(menu)
		fmt.Printf("%v.\t%v\t\t%.2f\n", count, menuShorted[menuInt], availableMenu[menuShorted[menuInt]])
		// fmt.Println(availableMenu[menuShorted[menuInt]])
		count++
	}
}

func (c *Consumer) CountBill(menuShorted []string, availableMenu map[string]float64) {
	for _, menu := range c.Order {
		menuInt, _ := strconv.Atoi(menu)
		fmt.Println(availableMenu[menuShorted[menuInt]])
		c.Bill += availableMenu[menuShorted[menuInt]]
	}
}

func (c *Consumer) Inactive() {
	c.Active = false
}

func (c *Consumer) AddTableNum(tableNum int) {
	c.TableNum = tableNum
}
