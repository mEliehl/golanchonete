package pedido

import (
	"github.com/meliehl/golanchonete/desconto"
	"github.com/meliehl/golanchonete/lanche"
)

//Pedido Responsavel por calcular o preço de cada lanche pedido
type Pedido struct {
	lanches []*pedidoLanche
}

type pedidoLanche struct {
	lanche   *lanche.Lanche
	desconto float32
}

var descontos = []desconto.Desconto{
	desconto.MuitoQueijo{},
	desconto.SejaLight{},
	desconto.DobroDeCarneGanhaBacon{},
}

//AdicionarIngrediente adiciona um lanche ao pedido
func (pedido *Pedido) AdicionarIngrediente(lanche *lanche.Lanche) {
	var sumDesconto float32
	for _, desconto := range descontos {
		sumDesconto += desconto.ValorDesconto(*lanche)
	}
	pedido.lanches = append(pedido.lanches, &pedidoLanche{lanche, sumDesconto})
}

//Total calcula o valor total do pedido
func (pedido Pedido) Total() float32 {
	var sum float32
	for _, lanche := range pedido.lanches {
		sum += lanche.lanche.Total() - lanche.desconto
	}
	return sum
}
