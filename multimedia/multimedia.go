package multimedia

import "fmt"

type ContenidoWeb struct {
	multimedia []Multimedia
}
type Multimedia interface {
	Mostrar()
}
type Imagen struct {
	t string
	f string
	c uint
}
type Audio struct {
	t string
	f string
	s uint
}
type Video struct {
	t  string
	f  string
	fr uint
}

func (i *Imagen) Mostrar() {
	fmt.Println("titulo : " + i.t)
	fmt.Println("formato: " + i.f)
	fmt.Println("canales: " + fmt.Sprint(i.c))
}

func (a *Audio) Mostrar() {
	fmt.Println("titulo : " + a.t)
	fmt.Println("formato: " + a.f)
	fmt.Println("canales: " + fmt.Sprint(a.s))
}

func (v *Video) Mostrar() {
	fmt.Println("titulo : " + v.t)
	fmt.Println("formato: " + v.f)
	fmt.Println("frames: " + fmt.Sprint(v.fr))
}
