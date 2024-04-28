# Tutorial

Algumas anotações relevantes do capítulo 1: Tutorial.

## Hello, World

- go **run** example.go -> compila o código-fonte de um ou mais arquivos-fonte cujos nomes terminam em .go, faz a ligação com as bibliotecas e, então, roda o arquivo executável resultante.
- go **build** example.go -> compila e salva o resultado compilado
- go **get** github.com/example/example -> busca o código-fonte e coloca no diretório correspondente

Todo código em Go está organizado em pacotes. Um pacote é constituído de um ou mais arquivos-fontes .go em um único diretório que define o que o pacote faz.
Cada arquivo-fonte começa com uma declaração **package**. O pacote **main** é especial por definir um programa executável independente, e não uma biblioteca.d
No pacote **main**, a função **main** é especial também, pois é onde a execução do programa começa.

## Argumentos de linha de comando

A variável os.Args é uma fatia (slice) de strings. 

## Encontrando linhas duplicadas

`fmt.Printf` gera uma saída formatada a partir de uma lista de expressões. Alguns recursos disponíveis são:

- %d -> inteiro em formato decimal
- %x, %o, %b -> inteiro em formato hexadecimal, octal, binário
- %f, %g, %e -> número de ponto flutuante
- %t -> booleano: *true* ou *false*
- %c -> runa
- %s -> string
- %q -> string "abc" ou runa 'c' entre aspas
- %v -> qualquer valor em um formato natural
- %T -> tipo de qualquer valor
- %% -> sinal de porcentagem literal