package desconto

import (
	"testing"

	"github.com/meliehl/golanchonete/ingrediente"
	"github.com/meliehl/golanchonete/lanche"
)

func TestDescontoSejaLightHambugerAlface(t *testing.T) {
	lanche := lanche.XBurger()
	lanche.AdicionarIngrediente(ingrediente.Alface())
	var desconto Desconto = SejaLight{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = .49
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoSejaLightHambugerBaconAlface(t *testing.T) {
	lanche := lanche.XBurger()
	lanche.AdicionarIngrediente(ingrediente.Alface())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	var desconto Desconto = SejaLight{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoSejaLightHambugerSemBaconSemAlface(t *testing.T) {
	lanche := lanche.XBurger()
	var desconto Desconto = SejaLight{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}
