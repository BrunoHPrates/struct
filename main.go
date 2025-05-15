package main

import "fmt"

type carro struct {
	modelo string
	ano int
	velocidade int
	potencia int
	aceleracao float32
}

func exibeDados(j carro){
	fmt.Println("Modelo do carro: ", j.modelo)
	fmt.Println("Ano do carro: ", j.ano)
	fmt.Println("Velocidade: ", j.velocidade,"km/h")
	fmt.Println("Potência: ", j.potencia,"cv")
	fmt.Println("0 a 100: ", j.aceleracao,"s")
}

func main() {
	j1 := carro{"NSX Type R", 1992, 270, 274, 4.7}
	exibeDados(j1)
}