package inflacao

import (
	"sync"
)

type Inflacao struct {
	valores []float32
}

var (
	inflacao *Inflacao
	once     sync.Once
)

//Inflacao singleton
func GetInflacao() *Inflacao {
	once.Do(func() {
		inflacao = &Inflacao{}
	})
	return inflacao
}

//AdicionarInflacao adicionar uma inflação globalmente
func (inflacao *Inflacao) AdicionarInflacao(valor float32) {
	inflacao.valores = append(inflacao.valores, valor)
}

//Valores get para valores
func (inflacao Inflacao) Valores() []float32 {
	return inflacao.valores
}

//LimparInflacao Basicamente para Testes
func (inflacao *Inflacao) LimparInflacao() {
	inflacao.valores = []float32{}
}
