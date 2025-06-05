package entity

type PassageirosManager interface {
	GetTotalPassagens(destination string) (int, error)
	GetPeriodo(period string) (int, error)
	CustoPorDestino(destination string) (float64, error)
}

type Passageiros struct {
	Id      int     `json:"id"`
	Nome    string  `json:"nome"`
	Email   string  `json:"email"`
	Pais    string  `json:"pais"`
	Horario string  `json:"horario"`
	Preco   float64 `json:"preco"`
}

type PassageirosService struct {
	List []Passageiros
}

func (ps PassageirosService) GetTotalPassagens() (int, error) {
	//tamanho := len(ps.List)

	// if tamanho == 0 {
	// 	return 0, errors.New("quantidade de passageiros é zero")
	// }
	return len(ps.List), nil
}

func (ps PassageirosService) CustoPorDestino(destination string) (float64, error) {
	total := 0.0
	count := 0
	for _, p := range ps.List {
		if p.Pais == destination {
			total += p.Preco
			count++
		}
	}
	if count == 0 {
		return 0, nil
	}
	return total / float64(count), nil
}
