package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	var data = student{
		name:        "wick",
		height:      182.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}

	fmt.Printf("Format numerik to string berbasis 2(biner): %b\n", data.age)  // 11010
	fmt.Printf("Format numerik (unicode) to string unicodenya: %c\n", 1400)   // n
	fmt.Printf("Format numerik (unicode) to string unicodenya: %c\n", 1235)   // ӓ
	fmt.Printf("Format numerik to numerik berbasis 10: %d\n", data.age)       // 26
	fmt.Printf("Format numerik to notasi numerik standar: %e\n", data.height) // 1.825000e+02
	fmt.Printf("Format numerik to notasi numerik standar: %E\n", data.height) // 1.825000E+02
	fmt.Println()

	fmt.Printf("Format numerik dengan lebar desimal (default 6 digit): %f\n", data.height) // 182.500000
	fmt.Printf("Format numerik dengan lebar desimal : %.9f\n", data.height)                // 182.500000000
	fmt.Printf("Format numerik dengan lebar desimal : %.2f\n", data.height)                // 182.50
	fmt.Printf("Format numerik dengan lebar desimal : %.f\n", data.height)                 // 182
	fmt.Println()

	fmt.Println("===============Perbedaan %e %f %f")
	fmt.Printf("Format numerik dengan lebar desimal e: %e\n", 0.123123123123)            // 1.231231e-01
	fmt.Printf("Format numerik dengan lebar desimal f: %f\n", 0.123123123123)            // 0.123123
	fmt.Printf("Format numerik dengan lebar desimal yang besar g: %g\n", 0.123123123123) // 0.123123123123

	fmt.Println()
	fmt.Printf("Format numerik to string numerik (8 oktal): %o\n", data.age)                                 // 32
	fmt.Printf("Format data pointer to alamat pointer referensi variabel (basis 16 + 0x): %p\n", &data.name) // 0xc00007e040
	fmt.Printf("Format untuk escape string: %q\n", `"name \ height"`)                                        // "\"name \\ height\""
	fmt.Printf("Format data to string: %s\n", data.name)                                                     // wick
	fmt.Printf("Formmat dat boolean: %t\n", data.isGraduated)                                                // false

	fmt.Println()
	fmt.Printf("Format tipe variabel: %T\n", data.name)        // string
	fmt.Printf("Format tipe variabel: %T\n", data.height)      // float64
	fmt.Printf("Format tipe variabel: %T\n", data.age)         // int32
	fmt.Printf("Format tipe variabel: %T\n", data.isGraduated) // bool
	fmt.Printf("Format tipe variabel: %T\n", data.hobbies)     // []string

	fmt.Println()
	fmt.Printf("Format tipe data apa saja: %v\n", data)  // {wick 182.5 26 false [eating sleeping]}
	fmt.Printf("Format tipe data apa saja: %+v\n", data) //{name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}
	fmt.Printf("Format tipe data apa saja: %#v\n", data)
	// main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:[]string{"eating", "sleeping"}}

	fmt.Println()
	fmt.Println("====Format %x atau %X")
	fmt.Printf("Format numerik to string numerik (basis 16): %x\n", data.age) //1a

	var d = data.name
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3]) // 7769636b
	fmt.Printf("%x\n", d)                            // 7769636b

	fmt.Println()
	fmt.Printf("Menulis karakter: %%\n") // %
}
