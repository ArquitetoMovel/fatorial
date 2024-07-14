package fatorial

type Fatorial struct {
  Numero    int
  Nfatorial int
}

func (f *Fatorial) Calcular() {
  f.Nfatorial = f.Numero  
  for i := f.Nfatorial - 1; i >= 1; i-- {
      f.Nfatorial *= i
  }
}
