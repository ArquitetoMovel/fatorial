package fatorial

type Fatorial struct {
  Numero int
}

func (f *Fatorial) Calcular() {
  for f.Numero > 1 {
    f.Numero *= f.Numero - 1
  }
  f.Numero *= 1
}
