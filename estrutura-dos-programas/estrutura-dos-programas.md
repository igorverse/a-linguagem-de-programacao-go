# Estrutura dos Programas

## Nomes

- Go tem 25 palavras reservadas
- Há distinções entre letras maiúsculas e minúsculas: **exemplo** e **Exemplo** são nomes diferentes
- Se um nome começa com letra maiúscula, ele é exportado. Ou seja, ele é visível e acessível fora do seu próprio pacote, podendo ser referenciado em outras partes do programa
- Nomes de pacotes são sempre minúsculos
- Para variáveis locais e escopos pequenos, a tendência é usar nomes curtos em Go
- Go costuma ser escrito usando o estilo [camel case](https://pt.wikipedia.org/wiki/CamelCase#:~:text=CamelCase%20%C3%A9%20a%20denomina%C3%A7%C3%A3o%20em,defini%C3%A7%C3%B5es%20de%20classes%20e%20objetos.)
- Para siglas e acrônimos é adotado algo do tipo: htmlExample, HTMLExample, exampleHTML. Mas **nunca** exampleHtml

## Declarações

- Há quatro tipos principais de declaração: var, const, type e func

## Variáveis

- Toda declaração de variável tem o formato geral: var nome tipo = expressão
- O *tipo* ou a parte *= expressão* pode ser omitido, mas não ambos
- Há uma *short syntax*: nome := expressão. Nesse modelo de declaração ocorre a inferência de tipo
- A declaração curta de variável nem sempre *declara* todas as variáveis do lado esquerdo. Contudo, uma declaração curta de variável deve declarar pelo menos uma nova variável

### Ponteiros

- Um ponteiro é um endereço de uma variável. Portanto, é o local em que um valor é armazenado. Nem todo valor tem um endereço, mas toda variável tem. É possível, por meio de um ponteiro, atualizar o valor de uma variável indiretamente, sem usar ou sequer saber o nome da variável

### Tempo de vida das variáveis

- O tempo de vida de uma variável de nível de pacote é toda a execução de um programa
- As variáveis locais têm tempo de vida dinâmicos: uma nova instância é criada sempre que a instrução de declaração é executada

## Atribuições

- O valor mantido por uma variável é atualizado por uma instrução de atribuição que, em sua forma mais simples, tem uma variável à esquerda do sinal **=** e uma expressão à direita

```
x = 1 // variável nomeada
*p = true // variável indireta
person.name = "bob" // campo de estrutura
count[x] = count[x] * scale // elemento de array, slice ou map
```

### Atribuição de tupla

- Permite que diversas variáveis recebam valores de uma só vez
- Toda expressão do lado direito é avaliada antes da variável ser atualizada
- Determinadas expressões podem resultar na geração de diversos valores. Quando é o caso, o lado esquerdo deve ter tantas variáveis quantos os resultados

```
f, err := os.Open("foo.txt") // a chamada da função retorna dois valores
```

- É possível atribuir valores que não queremos ao identificado vazio **_**

```
_, err = io.Copy(dst, src) // descarta o número de bytes
```

## Declaração de Tipo

- Um declaração **type** define um novo tipo nomeado que tem o mesmo tipo subjacente de um tipo existente: `type nome tipo-subjacente`
- Para todo tipo T há uma operação de conversão T(x) correspondente, que converte o valor x para o tipo T
- Qualquer que seja o caso, uma conversão jamais falha em tempo de execução

## Pacotes e Arquivos

- Pacotes em Go têm o mesmo propósito das bibliotecas ou módulos em outras linguagens; eles suportam modularidade, encapsulamento, compilação separada e reutilização
- Cada pacote serve como um name space (espaço de nomes) separado paras suas declarações
- Para referenciar uma função fora de seu pacote é necessário qualificar o identificador para deixar explícito se queremos dizer, por exemplo: `image.Decode` ou `utf16.Decode`

## Importações

- Todo pacote é identificado por uma string única chamada path de importação
- Um path de importação representa um diretório contendo um ou mais arquivos-fonte Go que, juntos, compõe o pacote
- Cada pacote tem um `package name` que, por convenção, coincide com o último segmento do seu `path de importação`

```
gopl.io/ch2/tempconv -> path de importação
tempconv -> package name
```