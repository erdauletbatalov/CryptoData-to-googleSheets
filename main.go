package main

import (
	service "github.com/erdauletbatalov/WEB3-Space---Terms-of-reference/src"
)

func main() {
	service.FillSpreadsheet('A', service.GetAPIFromCryptoRank(3))
	service.FillSpreadsheet('E', service.GetAPIFromCoinGecko(65))
}
