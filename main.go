package main

import (
	"fmt"

	"./multimedia"
)

func main() {
	opc := 1
	var t, f string
	var a uint

	for opc > 0 {
		fmt.Println("1.- Nuevo contenido web")
		fmt.Println("0.- Salir")
		fmt.Scanln(&opc)
		if opc == 1 {
			cWeb := multimedia.ContenidoWeb{Multimedia: []multimedia.Multimedia{}}
			fmt.Println("1.-Agregar imagen")
			fmt.Println("2.-Agregar audio")
			fmt.Println("3.-Agregar video")
			fmt.Println("4.-Mostrar")
			fmt.Println("5.-Nuevo contenido web")
			fmt.Println("0.-Salir")
			fmt.Scanln(&opc)
			for opc < 5 && opc > 0 {
				if opc < 4 {
					fmt.Println("Titulo: ")
					fmt.Scanln(&t)
					fmt.Println("Formato: ")
					fmt.Scanln(&f)
					if opc == 1 {
						fmt.Println("Canales: ")
						fmt.Scanln(&a)
						cWeb.Multimedia = append(cWeb.Multimedia, &multimedia.Imagen{T: t, F: f, C: a})
					}
					if opc == 2 {
						fmt.Println("Duracion en segundos: ")
						fmt.Scanln(&a)
						cWeb.Multimedia = append(cWeb.Multimedia, &multimedia.Audio{T: t, F: f, S: a})
					}
					if opc == 3 {
						fmt.Println("Frames: ")
						fmt.Scanln(&a)
						cWeb.Multimedia = append(cWeb.Multimedia, &multimedia.Video{T: t, F: f, Fr: a})
					}
				}
				if opc == 4 {
					cWeb.Mostrar()
				}
				fmt.Println("1.-Agregar imagen")
				fmt.Println("2.-Agregar audio")
				fmt.Println("3.-Agregar video")
				fmt.Println("4.-Mostrar")
				fmt.Println("5.-Nuevo contenido web")
				fmt.Println("0.-Salir")
				fmt.Scanln(&opc)
			}
			if opc == 5 {
				opc = 1
			}
		}
	}
}
