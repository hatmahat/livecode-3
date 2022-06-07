package data

import (
	"fmt"
	"strings"
)

type Menu struct {
	AvailableMenu map[string]float64
	MenuShorted   []string
}

func (m *Menu) AppendMenu(name string, price float64) {
	if m.AvailableMenu == nil {
		m.AvailableMenu = make(map[string]float64)
	}
	m.AvailableMenu[name] = price
}

func (m *Menu) DeleteMenu(name string) {
	delete(m.AvailableMenu, name)
}

func (m *Menu) ShowMenu() {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\tMENU WARUNG MAKAN BAHARI")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("No.\tMakanan\t\t\tHarga\n")
	fmt.Println(strings.Repeat("=", 40))
	var count int = 0
	for name, price := range m.AvailableMenu {
		fmt.Printf("%v.\t%v\t\t%.2f\n", count, name, price)
		m.MenuShorted = append(m.MenuShorted, name)
		count++
	}
	fmt.Println(strings.Repeat("=", 40))
}
