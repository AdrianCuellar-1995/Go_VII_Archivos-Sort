package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var opc uint64
	for opc != 4 {
		fmt.Println("1.- Ingresar cadenas")
		fmt.Println("2.- Mostrar de manera Ascendente")
		fmt.Println("3.- Mostrar de manera Descendente")
		fmt.Println("4.- Salir")
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			var n int
			fmt.Println(" Cantidad de  cadenas : ")
			fmt.Scanln(&n)
			datos := make([]string, n)
			for i := 0; i < n; i++ {
				fmt.Println(i+1, ")")
				fmt.Scanln(&datos[i])
			}
			sort.Strings(datos)
			EscribirArchivo("ascendente.txt", datos)
			sort.Sort(sort.Reverse(sort.StringSlice(datos)))
			EscribirArchivo("descendente.txt", datos)
		case 2:
			LeerArchivo("ascendente.txt")
		case 3:
			LeerArchivo("descendente.txt")
		case 4:
		default:
			fmt.Println("no existe la opcion")
		}
	}
}

func EscribirArchivo(nombre string, datos []string) {
	file, err := os.Create(nombre)
	if err != nil {
		fmt.Println("Err")
		return
	}
	defer file.Close()
	for _, cadena := range datos {
		file.WriteString(cadena + "\n")
	}
}

func LeerArchivo(nombre string) {
	file, err := os.Open(nombre)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	total := stat.Size()
	bs := make([]byte, total)
	count, err := file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bs)
	fmt.Println(str, "\n Bytes: ", count)
}
