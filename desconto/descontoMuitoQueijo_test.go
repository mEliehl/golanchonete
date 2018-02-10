package desconto

import (
	"testing"

	"github.com/meliehl/golanchonete/ingrediente"
	"github.com/meliehl/golanchonete/lanche"
)

func TestDescontoLancheTresQueijos(t *testing.T) {
	lanche := lanche.XBurger()
	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())
	var desconto Desconto = MuitoQueijo{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 1.5
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheSemQueijos(t *testing.T) {
	lanche := lanche.XBurger()
	var desconto Desconto = MuitoQueijo{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheUmQueijos(t *testing.T) {
	lanche := lanche.Lanche{}
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	var desconto Desconto = MuitoQueijo{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheSeisQueijos(t *testing.T) {
	lanche := lanche.XBurger()
	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())

	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())

	var desconto Desconto = MuitoQueijo{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 3
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestDescontoLancheOitoQueijos(t *testing.T) {
	lanche := lanche.XBurger()
	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())

	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())

	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())

	var desconto Desconto = MuitoQueijo{}
	actual := desconto.ValorDesconto(lanche)
	var expected float32 = 3
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}
