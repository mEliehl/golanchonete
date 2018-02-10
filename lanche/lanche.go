package lanche

import "github.com/meliehl/golanchonete/ingrediente"

//Lanche Um lanche é composto por muitos ingredientes
type Lanche struct {
	ingredientes []*ingrediente.Ingrediente
}

//AdicionarIngrediente set de Ingrediente
func (lanche *Lanche) AdicionarIngrediente(ingrediente *ingrediente.Ingrediente) {

	lanche.ingredientes = append(lanche.ingredientes, ingrediente)
}

//Total Calcula o total do lanche sem desconto
func (lanche Lanche) Total() float32 {
	var sum float32
	for _, ingrediente := range lanche.ingredientes {
		sum += ingrediente.Valor()
	}
	return sum
}

//XBurger factory de um xburger
func XBurger() Lanche {
	var ingredientes = make([]*ingrediente.Ingrediente, 2)
	ingredientes[0] = ingrediente.HamburgerCarne()
	ingredientes[1] = ingrediente.Queijo()
	hamburger := Lanche{ingredientes: ingredientes}
	return hamburger
}

//XBacon factory de um xBacon
func XBacon() Lanche {
	var ingredientes = make([]*ingrediente.Ingrediente, 3)
	ingredientes[0] = ingrediente.HamburgerCarne()
	ingredientes[1] = ingrediente.Queijo()
	ingredientes[2] = ingrediente.Bacon()
	hamburger := Lanche{ingredientes: ingredientes}
	return hamburger
}

//XEgg factory de um xEgg
func XEgg() Lanche {
	var ingredientes = make([]*ingrediente.Ingrediente, 3)
	ingredientes[0] = ingrediente.HamburgerCarne()
	ingredientes[1] = ingrediente.Queijo()
	ingredientes[2] = ingrediente.Ovo()
	hamburger := Lanche{ingredientes: ingredientes}
	return hamburger
}

//XEggBacon factory de um xEggBacon
func XEggBacon() Lanche {
	var ingredientes = make([]*ingrediente.Ingrediente, 4)
	ingredientes[0] = ingrediente.HamburgerCarne()
	ingredientes[1] = ingrediente.Queijo()
	ingredientes[2] = ingrediente.Ovo()
	ingredientes[3] = ingrediente.Bacon()
	hamburger := Lanche{ingredientes: ingredientes}
	return hamburger
}

//TemIngrediente Verifica se o lanche possui o ingrediente passado
func (lanche Lanche) TemIngrediente(ingrediente *ingrediente.Ingrediente) bool {
	for _, _ingrediente := range lanche.ingredientes {
		if _ingrediente == ingrediente {
			return true
		}
	}
	return false
}

//QuantidadeIngrediente retorna a Quantidade do ingrediente passado
func (lanche Lanche) QuantidadeIngrediente(ingrediente *ingrediente.Ingrediente) int {
	var sum int
	for _, _ingrediente := range lanche.ingredientes {
		if _ingrediente == ingrediente {
			sum++
		}
	}
	return sum
}
