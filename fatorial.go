package fatorial

import (
  "fmt"
)

type Fatorial struct {
  Numero   int
  Fator int
}

func (f Fatorial) Calcular() {
  fmt.Println("Iniciando o calculo para", f.Numero)
  for f.Numero > 1 {
    num := (f.Numero - 1)
    f.Fator = f.Numero * num
    f.Numero = num
    fmt.Println("fatorial = ", f.Numero)
  }
  fmt.Println("fim", f.Numero)

  if f.Numero == 1 { 
    fmt.Println("fator end") 
  }
}
