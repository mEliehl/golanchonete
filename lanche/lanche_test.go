package lanche

import (
	"testing"

	"github.com/meliehl/golanchonete/ingrediente"
)

func TestTotalXBurger(t *testing.T) {
	lanche := XBurger()
	actual := lanche.Total()
	var expected float32 = 4.5
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestTotalXBacon(t *testing.T) {
	lanche := XBacon()
	actual := lanche.Total()
	var expected float32 = 6.5
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestTotalXEgg(t *testing.T) {
	lanche := XEgg()
	actual := lanche.Total()
	var expected float32 = 5.3
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestTotalXEggBacon(t *testing.T) {
	lanche := XEggBacon()
	actual := lanche.Total()
	var expected float32 = 7.3
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestTotalXBurgerComAlface(t *testing.T) {
	lanche := XBurger()
	lanche.AdicionarIngrediente(ingrediente.Alface())
	actual := lanche.Total()
	var expected float32 = 4.9
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}
func TestTotalXBurgerComDoisHamburgersDeCarne(t *testing.T) {
	lanche := XBurger()
	lanche.AdicionarIngrediente(ingrediente.HamburgerCarne())
	actual := lanche.Total()
	var expected float32 = 7.5
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestXburgerTemQueijo(t *testing.T) {
	lanche := XBurger()
	var actual = lanche.TemIngrediente(ingrediente.Queijo())
	if !actual {
		t.Fatalf("Expected true")
	}
}

func TestXburgerNaoTemBacon(t *testing.T) {
	lanche := XBurger()
	var actual = lanche.TemIngrediente(ingrediente.Bacon())
	if actual {
		t.Fatalf("Expected false")
	}
}

func TestXburgerTemUmaCarne(t *testing.T) {
	lanche := XBurger()
	var actual = lanche.QuantidadeIngrediente(ingrediente.HamburgerCarne())
	var expected = 1
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestXburgerTemZeroBacon(t *testing.T) {
	lanche := XBurger()
	var actual = lanche.QuantidadeIngrediente(ingrediente.Bacon())
	var expected = 0
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}
