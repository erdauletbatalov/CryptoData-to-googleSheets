package main

import (
	service "github.com/erdauletbatalov/WEB3-Space---Terms-of-reference/src"
)

func main() {
	service.FillSpreadsheet(service.GetAPIFromCryptoRank())

}
