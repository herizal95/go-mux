package barang

import (
	"net/http"

	"github.com/herizal95/gomux/helper"
)

func GetAllBarang(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{

		{
			"id":          1,
			"nama_barang": "kemeja",
			"stok":        100,
		},
		{
			"id":          2,
			"nama_barang": "celana",
			"stok":        200,
		},
		{
			"id":          3,
			"nama_barang": "sepatu",
			"stok":        100,
		},
	}
	helper.ResponseJson(w, 200, data)

}
