package data

import "fmt"

type Table struct {
	TableID   int
	Customers []Consumer
	Available bool
}

type Tables struct {
	Tables []Table
}

func (t *Tables) GenerateAvailableTable(num int) {
	// TableID mulai dari 0
	for i := 0; i < num; i++ {
		t.Tables = append(t.Tables, Table{
			TableID:   i,
			Available: true,
		})
	}
}

func (t *Tables) ServeCustomerAtTableID(tableID int, consumer Consumer) {
	// hanya bisa append ketika t.Tables[table.ID].Available == true dan custumer aktif
	// seteleah di-append maka t.Tables[table.ID].Available = false dan custumer tetao aktif
	// jika kostumer sudah tidak aktif maka t.Tables[table.ID].Available = true

	if t.Tables[tableID].Available == true && consumer.Active == true {
		t.Tables[tableID].Customers = append(t.Tables[tableID].Customers, consumer)
		// consumer.TableNum = tableID         // tambah pada struct Consumer
		t.Tables[tableID].Available = false // karena udah diisi
	} else if consumer.Active == false {
		fmt.Println("Customer", consumer, "tidak aktif")
	} else {
		fmt.Println("Meja terisi", consumer, "tidak bisa di meja", tableID)
	}
}

func (t *Tables) DoneCustomerAtTableID(tableID int, consumer Consumer) { // kosongkan meja
	if consumer.Active == false && t.Tables[tableID].Available == false {
		t.Tables[tableID].Available = true
	}
}

func (t *Tables) AvailableTablesNums() []int {
	var availableTableNums []int
	for num, table := range t.Tables {
		// fmt.Println("AVA", table.Available) // debug code
		if table.Available {
			availableTableNums = append(availableTableNums, num)
		}
	}
	return availableTableNums
}
