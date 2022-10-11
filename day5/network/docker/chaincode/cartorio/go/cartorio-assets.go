package cartorio

type Imovel struct {
	Inscricao string `json:"inscricao"`
	Nome      string `json:"nome"`
	Valor     int    `json:"valor"`
	Tamanho   int    `json:"tamanho"`
}

type ImovelComercial struct {
	Imovel
	Cnpj string `json:"cnpj"`
}

type ImovelResidencial struct {
	Imovel
	Cpf string `json:"cpf"`
}

//------------------------------------

type MintImovelPayload struct {
	Inscricao string `json:"inscricao"`
	Nome      string `json:"nome"`
	Tamanho   int    `json:"tamanho"`
}

type MintImovelComercialPayload struct {
	Inscricao string `json:"inscricao"`
	Nome      string `json:"nome"`
	Tamanho   int    `json:"tamanho"`
	Cnpj      string `json:"cnpj"`
}

type MintImovelResidencialPayload struct {
	Inscricao string `json:"inscricao"`
	Nome      string `json:"nome"`
	Tamanho   int    `json:"tamanho"`
	Cpf       string `json:"cpf"`
}
