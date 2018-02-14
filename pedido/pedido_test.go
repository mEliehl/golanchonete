package pedido

import (
	"testing"

	"github.com/meliehl/golanchonete/inflacao"

	"github.com/meliehl/golanchonete/ingrediente"
	"github.com/meliehl/golanchonete/lanche"
)

func TestPedirUmXBurger(t *testing.T) {
	var lanche = lanche.XBurger()

	var pedido Pedido
	pedido.AdicionarLanche(&lanche)

	var actual = pedido.Total()
	var expected float32 = 4.5
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestPedirUmXBurgerComInflacao(t *testing.T) {
	var inf = inflacao.GetInflacao()
	inf.AdicionarInflacao(10)
	defer inf.LimparInflacao()
	var lanche = lanche.XBurger()

	var pedido Pedido
	pedido.AdicionarLanche(&lanche)

	var actual = pedido.Total()
	var expected float32 = 4.95
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestPedirUmXBurgerComInflacaoSobreInflacao(t *testing.T) {
	var inf = inflacao.GetInflacao()
	inf.AdicionarInflacao(10)
	inf.AdicionarInflacao(20)
	defer inf.LimparInflacao()
	var lanche = lanche.XBurger()

	var pedido Pedido
	pedido.AdicionarLanche(&lanche)

	var actual = pedido.Total()
	var expected float32 = 5.94
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestPedirUmXBurgerEUmXBacon(t *testing.T) {
	var lanche1 = lanche.XBurger()
	var lanche2 = lanche.XBacon()

	var pedido Pedido
	pedido.AdicionarLanche(&lanche1)
	pedido.AdicionarLanche(&lanche2)

	var actual = pedido.Total()
	var expected float32 = 11
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestPedirUmXBurgerEUmXBaconUmSejaLight(t *testing.T) {
	var lanche1 = lanche.XBurger()
	var lanche2 = lanche.XBacon()
	var lanche3 = lanche.XBurger()
	lanche3.AdicionarIngrediente(ingrediente.Alface())

	var pedido Pedido
	pedido.AdicionarLanche(&lanche1)
	pedido.AdicionarLanche(&lanche2)
	pedido.AdicionarLanche(&lanche3)

	var actual = pedido.Total()
	var expected float32 = 15.41
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestPedirUmXBurgerEUmXBaconUmSejaLightUmMuitoQueijo(t *testing.T) {
	var lanche1 = lanche.XBurger()
	var lanche2 = lanche.XBacon()
	var lanche3 = lanche.XBurger()
	lanche3.AdicionarIngrediente(ingrediente.Alface())
	var lanche4 = lanche.XBurger()
	lanche4.AdicionarIngrediente(ingrediente.Queijo())
	lanche4.AdicionarIngrediente(ingrediente.Queijo())

	var pedido Pedido
	pedido.AdicionarLanche(&lanche1)
	pedido.AdicionarLanche(&lanche2)
	pedido.AdicionarLanche(&lanche3)
	pedido.AdicionarLanche(&lanche4)

	var actual = pedido.Total()
	var expected float32 = 21.41
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestPedirUmXBurgerEUmXBaconUmSejaLightUmMuitoQueijoEUmDobroCarneGanhaQueijo(t *testing.T) {
	var lanche1 = lanche.XBurger()
	var lanche2 = lanche.XBacon()
	var lanche3 = lanche.XBurger()
	lanche3.AdicionarIngrediente(ingrediente.Alface())
	var lanche4 = lanche.XBurger()
	lanche4.AdicionarIngrediente(ingrediente.Queijo())
	lanche4.AdicionarIngrediente(ingrediente.Queijo())

	var lanche5 = lanche.XBacon()
	lanche5.AdicionarIngrediente(ingrediente.HamburgerCarne())

	var pedido Pedido
	pedido.AdicionarLanche(&lanche1)
	pedido.AdicionarLanche(&lanche2)
	pedido.AdicionarLanche(&lanche3)
	pedido.AdicionarLanche(&lanche4)
	pedido.AdicionarLanche(&lanche5)

	var actual = pedido.Total()
	var expected float32 = 28.91
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}

func TestPedirUmMuitoQueijoESejaLight(t *testing.T) {
	var lanche = lanche.XBurger()
	lanche.AdicionarIngrediente(ingrediente.Alface())
	lanche.AdicionarIngrediente(ingrediente.Queijo())
	lanche.AdicionarIngrediente(ingrediente.Queijo())

	var pedido Pedido
	pedido.AdicionarLanche(&lanche)

	var actual = pedido.Total()
	var expected float32 = 5.61
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}
}
