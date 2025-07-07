package main

import "fmt"

type forma interface {
	area() float64
}

type formaGeometrica struct {
	tipo string
}

type quadrado struct {
	formaGeometrica
	lado float64
}

func novoQuadrado(lado float64) quadrado {
	return quadrado{formaGeometrica{"Quadrado"}, lado}
}

/*
declara igual o metodo, mas o nome do método tem que ser o mesmo
da interface
*/
func (q quadrado) area() float64 {
	return q.lado * 2
}

type triangulo struct {
	formaGeometrica
	largura float64
	altura  float64
}

func novoTriangulo(largura, altura float64) triangulo {
	return triangulo{formaGeometrica{"Triangulo"}, largura, altura}
}
func (t triangulo) area() float64 {
	return (t.largura * t.altura) / 2
}

func escreveArea(f forma) {
	fmt.Printf("A área é: %.2f\n", f.area())
}

func (h formaGeometrica) teseHeranca() {
	fmt.Printf("A forma é: %s\n", h.tipo)
}

func main() {
	quadrado1 := novoQuadrado(7.7)
	triangulo1 := novoTriangulo(10.9, 5.1)
	escreveArea(quadrado1)
	escreveArea(triangulo1)

	quadrado1.teseHeranca()
	triangulo1.teseHeranca()

}
