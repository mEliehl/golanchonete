package desconto

import (
	"testing"

	"github.com/meliehl/golanchonete/ingrediente"
	"github.com/meliehl/golanchonete/lanche"
)

func TestDescontoLancheDuasCarnesUmBacon(t *testing.T) {
	lanche := lanche.Lanche{}
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	var desconto Desconto = DobroDeCarneGanhaBacon{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 2
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheDuasCarnesNenhumBacon(t *testing.T) {
	lanche := lanche.Lanche{}
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	var desconto Desconto = DobroDeCarneGanhaBacon{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheQuatroCarnesUmBacon(t *testing.T) {
	lanche := lanche.Lanche{}
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.Bacon())

	var desconto Desconto = DobroDeCarneGanhaBacon{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 2
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheDuasCarnesTresBacon(t *testing.T) {
	lanche := lanche.Lanche{}
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	var desconto Desconto = DobroDeCarneGanhaBacon{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 2
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheTresCarnesTresBacon(t *testing.T) {
	lanche := lanche.Lanche{}

	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())

	lanche.AdicionarIngrediente(ingrediente.Bacon())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	lanche.AdicionarIngrediente(ingrediente.Bacon())

	var desconto Desconto = DobroDeCarneGanhaBacon{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 2
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheQuatroCarnesTresBacon(t *testing.T) {
	lanche := lanche.Lanche{}
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())

	lanche.AdicionarIngrediente(ingrediente.Bacon())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	lanche.AdicionarIngrediente(ingrediente.Bacon())
	var desconto Desconto = DobroDeCarneGanhaBacon{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 4
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}
