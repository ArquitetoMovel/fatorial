# fatorial
Um simples teste de como distribuir modulos em GO

## como usar

```
package main

import (
  "fmt"
  "github.com/arquitetomovel/fatorial"
)

func main() {

  nFatorial := &fatorial.Fatorial{
    Numero: 42,
  }

  nFatorial.Calcular()

  fmt.Println("foi calculado o fatorial", nFatorial.Nfatorial, "para o numero", nFatorial.Numero)

}

```
