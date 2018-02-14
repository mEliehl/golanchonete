package ingrediente

import (
	"testing"

	"github.com/meliehl/golanchonete/inflacao"
)

func TestAlfaceIsSingleton(t *testing.T) {
	alface := Alface()

	if alface != Alface() {
		t.Fatalf("Esperado um singleton")
	}
}

func TestBaconIsSingleton(t *testing.T) {
	bacon := Bacon()

	if bacon != Bacon() {
		t.Fatalf("Esperado um singleton")
	}
}

func TestHamburgerCarneIsSingleton(t *testing.T) {
	bacon := HamburgerCarne()

	if bacon != HamburgerCarne() {
		t.Fatalf("Esperado um singleton")
	}
}

func TestOvoIsSingleton(t *testing.T) {
	ovo := Ovo()

	if ovo != Ovo() {
		t.Fatalf("Esperado um singleton")
	}
}

func TestQueijoIsSingleton(t *testing.T) {
	queijo := Queijo()

	if queijo != Queijo() {
		t.Fatalf("Esperado um singleton")
	}
}

func TestAlfacePropriedades(t *testing.T) {
	alface := Alface()
	actualNome := alface.Nome()
	expectedNome := "Alface"
	if actualNome != expectedNome {
		t.Fatalf("Esperado %v mas temos %v", expectedNome, actualNome)
	}

	actualValor := alface.Valor()
	var expectedValor float32 = .4
	if actualValor != expectedValor {
		t.Fatalf("Esperado %v mas temos %v", expectedNome, actualNome)
	}
}

func TestAlfaceValorComInflacao(t *testing.T) {
	alface := Alface()
	var inf = inflacao.GetInflacao()
	inf.AdicionarInflacao(10)

	actual := alface.Valor()
	var expected float32 = .44
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", actual, expected)
	}
}
