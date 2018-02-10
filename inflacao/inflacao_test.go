package inflacao

import (
	"testing"
)

func TestAdicinarUmaInflacao(t *testing.T) {
	var inf = GetInflacao()
	inf.AdicionarInflacao(10)
	defer inf.LimparInflacao()
	if len(inf.Valores()) != 1 {
		t.Fatalf("Esperado array com apenas um valor")
	}
	var actual = inf.Valores()[0]
	var expected float32 = 10
	if actual != expected {
		t.Fatalf("Esperado %v mas temos %v", expected, actual)
	}

}

func TestAdicinarInflacaoSobreInflacao(t *testing.T) {
	var inf = GetInflacao()
	inf.AdicionarInflacao(10)
	inf.AdicionarInflacao(20)
	defer inf.LimparInflacao()

	if len(inf.Valores()) != 2 {
		t.Fatalf("Esperado array com dois valores")
	}

}

func TestAdicinarInflacaoSobreInflacaoELimpar(t *testing.T) {
	var inf = GetInflacao()
	inf.AdicionarInflacao(10)
	inf.AdicionarInflacao(20)
	inf.LimparInflacao()

	if len(inf.Valores()) != 0 {
		t.Fatalf("Experado array vazio")
	}

}
