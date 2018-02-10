package desconto

import (
	"github.com/meliehl/golanchonete/ingrediente"
	"github.com/meliehl/golanchonete/lanche"
)

//Desconto Interface padrão para implementação de descontos
type Desconto interface {
	ValorDesconto(lanche lanche.Lanche) float32
}

//SejaLight se tiver alface e não tiver bacon, ganhe 10% de desconto.
type SejaLight struct {
}

//ValorDesconto implementação da promoção seja light
func (desconto SejaLight) ValorDesconto(lanche lanche.Lanche) float32 {
	var retorno float32
	var temBacon = lanche.TemIngrediente(ingrediente.Bacon())
	var temAlface = lanche.TemIngrediente(ingrediente.Alface())
	if !temBacon && temAlface {
		retorno = (lanche.Total() * .1)
	}
	return retorno
}

//MuitoQueijo se tiver duas porções de queijo, ganhe a terceira.
type MuitoQueijo struct {
}

//ValorDesconto implementação da promoção muito queijo
func (desconto MuitoQueijo) ValorDesconto(lanche lanche.Lanche) float32 {
	var qtdeQueijo = lanche.QuantidadeIngrediente(ingrediente.Queijo())
	var multiplicador = qtdeQueijo / 3
	return float32(multiplicador) * ingrediente.Queijo().Valor()
}

//DobroDeCarneGanhaBacon se tiver duas carnes ganha um bacon
type DobroDeCarneGanhaBacon struct {
}

//ValorDesconto implementação da promoção muito queijo
func (desconto DobroDeCarneGanhaBacon) ValorDesconto(lanche lanche.Lanche) float32 {
	var qtdeCarne = lanche.QuantidadeIngrediente(ingrediente.HamburgerCarne())
	var qtdeBacon = lanche.QuantidadeIngrediente(ingrediente.Bacon())
	if qtdeBacon == 0 {
		return 0
	}
	var qtdeMaximoDesconto = int(qtdeCarne / 2)

	if qtdeBacon >= qtdeMaximoDesconto {
		return float32(qtdeMaximoDesconto) * ingrediente.Bacon().Valor()
	}
	return float32(qtdeBacon) * ingrediente.Bacon().Valor()
}
