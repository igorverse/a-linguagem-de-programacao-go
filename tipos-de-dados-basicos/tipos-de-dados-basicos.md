# Tipos de Dados Básicos

Os tipos de Go classificam-se em quatro grupos: tipos básicos, tipos agregados, tipos referência e tipos interface. Os tipos básicos incluem números, strings e booleanos

## Inteiros

- Go permite operações numéricas com inteiros com sinal e sem sinal
- Há quatro tamanhos para inteiros com sinal: `int8`, `int16`, `int32` e `int64`
- Também há quatro tamanhos para inteiros sem sinal: `uint8`, `uint16`, `uint32` e `uint64`
- Há os inteiros com tamanho natural: `int` e `uint`. Eles possuem o tamanho de 32 ou 64 bits, ficando a cargo do compilador a escolha de qual tamanho será usado
- `rune` é sinônimo para `int32`. Por convenção, indica que um valor é um código Unicode
- `byte` é sinônimo para `uint8`, usado para enfatizar que um valor é um dado bruto
- int é diferente de int32, assim como é diferente de int64

## Números de ponto flutuante

- Go oferece dois tamanhos de números de ponto flutuante: `float32` e `float64`
- Deve-se dar preferência a `float64` na maioria dos casos
