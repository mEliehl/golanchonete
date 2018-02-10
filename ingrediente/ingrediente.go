package ingrediente

import (
	"sync"

	"github.com/meliehl/golanchonete/inflacao"
)

//Ingrediente ingredientes disponiveis para compor um lanche
type Ingrediente struct {
	nome  string
	valor float32
}

//Nome get para o nome
func (ingrediente Ingrediente) Nome() string {
	return ingrediente.nome
}

//Valor get para o valor
func (ingrediente Ingrediente) Valor() float32 {
	var inf = inflacao.GetInflacao()

	var sum = ingrediente.valor
	for _, valor := range inf.Valores() {
		sum += (sum * valor) / 100
	}

	return sum
}

var (
	alface, bacon, hamburgerCarne, ovo, queijo                     *Ingrediente
	onceAlface, onceBacon, onceHamburgerCarne, onceOvo, onceQueijo sync.Once
)

//Alface singleton
func Alface() *Ingrediente {
	onceAlface.Do(func() {
		alface = &Ingrediente{"Alface", .4}
	})
	return alface
}

//Bacon singleton
func Bacon() *Ingrediente {
	onceBacon.Do(func() {
		bacon = &Ingrediente{"Bacon", 2}
	})
	return bacon
}

//HamburgerCarne singleton
func HamburgerCarne() *Ingrediente {
	onceHamburgerCarne.Do(func() {
		hamburgerCarne = &Ingrediente{"Hamburger de Carne", 3}
	})
	return hamburgerCarne
}

//Ovo singleton
func Ovo() *Ingrediente {
	onceOvo.Do(func() {
		ovo = &Ingrediente{"Ovo", .8}
	})
	return ovo
}

//Queijo singleton
func Queijo() *Ingrediente {
	onceQueijo.Do(func() {
		queijo = &Ingrediente{"Queijo", 1.50}
	})
	return queijo
}
