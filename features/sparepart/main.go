package sparepart

import (
	"bengkelin/common"
	sparepart "bengkelin/features/sparepart/functions"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func Main(spareparts *sStruct.ArrSparepart) {
	common.ResetConsole()
	
	var input string

	// Menampilkan main menu
	sparepart.ShowSparepartMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)

	for input != "0" {

		// Mengecek ketersediaan menu
		if input == "1" {
			sparepart.AddSparepart(spareparts)
			sparepart.ShowSparepartMenu()
		} else if input == "2" {
			common.ResetConsole()
			sparepart.ShowSparepart(*spareparts)
			common.ShowEndAction(1)
			sparepart.ShowSparepartMenu()
		} else if input == "3" {
			sparepart.EditSparepart(spareparts)
			sparepart.ShowSparepartMenu()
		} else if input == "4" {
			sparepart.DeleteSparepart(spareparts)
			sparepart.ShowSparepartMenu()
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}
	
	common.ResetConsole()
}
