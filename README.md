### Kurs
Dibuat dalam rangka belajar membuat package golang

### Keterangan
Package untuk mendapatkan data kurs mata uang bank di indonesia sementara ini hanya mendukung bank indonesia ("bi") dan bank mandiri ("mandiri")

### Contoh penggunaan

```golang
package main

import (
	"fmt"
	"github.com/iorme/kurs"
)

func main() {
	fmt.Println(kurs.GetCurrency("mandiri"))
}
```