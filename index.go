package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"wrb/data"
)

var ConsumerID int = 0

func main() {

	mainMenu()
	mainScanner := bufio.NewScanner(os.Stdin)

	totalMejaMakan := 30

	table := new(data.Tables)                    // bikin meja
	table.GenerateAvailableTable(totalMejaMakan) // bikin 10 meja, TableID 0-9

	mainScanner.Scan()
	condition := mainScanner.Text()

	menu := new(data.Menu)

	var customer []data.Consumer
	var customerCount int = 0

	for condition != "" {
		if condition == "1" {

			fmt.Print("Nama Makanan: ")
			mainScanner.Scan()
			namaMakanan := mainScanner.Text()

			fmt.Print("Harga Makanan: ")
			mainScanner.Scan()
			harga := mainScanner.Text()
			hargaInt, _ := strconv.Atoi(harga)

			menu.AppendMenu(namaMakanan, float64(hargaInt))

			for namaMakanan != "" {
				fmt.Print("Nama Makanan (tekan Enter saja untuk berhenti menambah): ")
				mainScanner.Scan()
				namaMakanan = mainScanner.Text()

				if namaMakanan == "" {
					continue
				}

				fmt.Print("Harga Makanan: ")
				mainScanner.Scan()
				harga := mainScanner.Text()
				hargaInt, _ := strconv.Atoi(harga)
				menu.AppendMenu(namaMakanan, float64(hargaInt))
			}

			fmt.Println()
			fmt.Println(strings.Repeat("=", 40))
			fmt.Println("* PEMBERITAHUAN: Menu telah ditambah")
			fmt.Println(strings.Repeat("=", 40))
			fmt.Println()

			mainMenu()
			mainScanner.Scan()
			condition = mainScanner.Text()

		} else if condition == "2" {

			menu.ShowMenu()
			mainMenu()

			mainScanner.Scan()
			condition = mainScanner.Text()

		} else if condition == "3" {

			menu.ShowMenu()
			appendCustomer(&customer, screenInsertConsumer())

			if len(table.AvailableTablesNums()) != 0 {

				MejadID := table.AvailableTablesNums()[0]
				table.ServeCustomerAtTableID(MejadID, customer[customerCount])
				customer[customerCount].AddTableNum(MejadID)

				fmt.Println()
				fmt.Println(strings.Repeat("=", 40))
				fmt.Println("* PEMBERITAHUAN:", customer[customerCount].Name, "silahkan menuju ke meja No", MejadID)
				fmt.Println(strings.Repeat("=", 40))
				fmt.Println()

			} else {
				fmt.Println("Maaf Pak/Bu", customer[customerCount].Name, "meja penuh")
			}
			customerCount++

			// tujukin meja berapa
			mainMenu()
			mainScanner.Scan()
			condition = mainScanner.Text()
		} else if condition == "4" {

			fmt.Print("Nama pelanggan yang bayar: ")
			mainScanner.Scan()
			pelanggan := mainScanner.Text()

			customerOut := customerDone(customer, pelanggan, menu.MenuShorted, menu.AvailableMenu) // masukan nama yg udah selesai/mau bayar

			// gimana dapatin object customer dan id mejanya
			table.DoneCustomerAtTableID(customerOut.TableNum, customerOut) // meja 0 dikosongkan
			fmt.Println()
			// rincian pesanan
			// fmt.Println(strings.Repeat("=", 40))
			// fmt.Println(customerOut)
			customerOut.ShowBillList(menu.MenuShorted, menu.AvailableMenu)
			fmt.Println(strings.Repeat("=", 40))
			fmt.Println("* Total yang dibayar oleh Bapak/Ibu", customerOut.Name, "adalah", customerOut.Bill, "dari meja No", customerOut.TableNum)
			fmt.Println(strings.Repeat("=", 40))
			fmt.Println()
			mainMenu()
			mainScanner.Scan()
			condition = mainScanner.Text()
		} else if condition == "5" {
			fmt.Println(strings.Repeat("=", 50))
			fmt.Printf("Nama\t\tNo Meja\t\tBill\tPAID/NOT PAID\n")
			fmt.Println(strings.Repeat("=", 50))

			var paidMessage string

			for _, cus := range customer {
				switch cus.Active {
				case true:
					paidMessage = "NOT PAID"
				case false:
					paidMessage = "PAID"
				}
				fmt.Println(cus.Name, "\t\t", cus.TableNum, "\t\t", cus.Bill, "\t", paidMessage)
			}
			fmt.Println(strings.Repeat("=", 50))

			mainMenu()
			mainScanner.Scan()
			condition = mainScanner.Text()
		} else if condition == "6" {
			fmt.Println(strings.Repeat("=", 50))
			fmt.Println("No Meja\t\tStatus")
			fmt.Println(strings.Repeat("=", 50))
			for _, table := range table.Tables {
				fmt.Println(table.TableID, "\t\t", table.Available)
			}

			mainMenu()
			mainScanner.Scan()
			condition = mainScanner.Text()
		}
	}
}

func addCustomer(id int, name string, order []string) data.Consumer {
	// inisiasi customer
	return data.Consumer{
		ID:     id,
		Name:   name,
		Order:  order,
		Active: true,
	}
}

func appendCustomer(arrOfCustomer *[]data.Consumer, customer data.Consumer) {
	*arrOfCustomer = append(*arrOfCustomer, customer)
}

func customerDone(arrOfCustomer []data.Consumer, name string, menuShorted []string, availableMenu map[string]float64) data.Consumer {
	var customerOut data.Consumer
	for idx, c := range arrOfCustomer {
		if c.Name == name {
			c.Inactive()                            // sudah tidak di warung jadi active false
			c.CountBill(menuShorted, availableMenu) // hitung bill
			arrOfCustomer[idx] = c                  // ditambahkan ke index itu lagi
			customerOut = arrOfCustomer[idx]
		}
	}
	return customerOut
}

func mainMenu() {
	fmt.Println(strings.Repeat("*", 40))
	fmt.Println("\t\tMenu Kasir")
	fmt.Println("\tPilih nomor berikut untuk:")
	fmt.Println(strings.Repeat("*", 40))
	fmt.Println("1. Masukkan menu hari ini")
	fmt.Println("2. Tampilkan semua menu yang telah dimasukkan")
	fmt.Println("3. Pelanggan masuk/mesan")
	fmt.Println("4. Pelanggan keluar/bayar")
	fmt.Println("5. Tampilkan semua status pelanggan terdaftar")
	fmt.Println("6. Tampilkan meja tersedia")
	fmt.Println()
	fmt.Println("\tTekan Enter untuk keluar")

}

func screenInsertConsumer() data.Consumer {
	var (
		name        string
		orderString string
		order       []string
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Nama Pelanggan: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Nomor yang ingin dipesanan pada menu: ")
	scanner.Scan()

	orderString = scanner.Text()
	order = append(order, orderString)
	for orderString != "" { // perlu penamahan kalau di luar index dari menu
		fmt.Print("Nomor (tekan Enter saja untuk selesai memesan): ")
		scanner.Scan()
		orderString = scanner.Text()
		order = append(order, orderString)
	}

	customer := data.Consumer{
		ID:     ConsumerID,
		Name:   name,
		Order:  order[:len(order)-1],
		Active: true,
	}
	ConsumerID += 1
	return customer
}
