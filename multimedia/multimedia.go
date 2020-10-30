package multimedia

import "fmt"

type ContenidoWeb struct {
	Multimedia []Multimedia
}

func (mm *ContenidoWeb) Mostrar() {
	fmt.Println("Mostrar contenido web")
	for _, m := range mm.Multimedia {
		m.Mostrar()
		fmt.Println()
	}
}

type Multimedia interface {
	Mostrar()
}
type Imagen struct {
	T string
	F string
	C uint
}
type Audio struct {
	T string
	F string
	S uint
}
type Video struct {
	T  string
	F  string
	Fr uint
}

func (i *Imagen) Mostrar() {
	fmt.Println("titulo : " + i.T)
	fmt.Println("formato: " + i.F)
	fmt.Println("canales: " + fmt.Sprint(i.C))
}

func (a *Audio) Mostrar() {
	fmt.Println("titulo : " + a.T)
	fmt.Println("formato: " + a.F)
	fmt.Println("canales: " + fmt.Sprint(a.S))
}

func (v *Video) Mostrar() {
	fmt.Println("titulo : " + v.T)
	fmt.Println("formato: " + v.F)
	fmt.Println("frames: " + fmt.Sprint(v.Fr))
}
