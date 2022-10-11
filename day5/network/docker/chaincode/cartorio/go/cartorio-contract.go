package cartorio

import (
	"fmt"

	"github.com/s7techlab/cckit/router"
)

func imovelMint(c router.Context) (interface{}, error) {
	mintData := c.Param("mintData").(MintImovelPayload)

	imovel := &Imovel{
			Inscricao: mintData.Inscricao,
			Nome:      mintData.Nome,
			Tamanho:   mintData.Tamanho,
	}

	//TODO: se ja existe dar um erro
	///

	err := c.State().Insert(imovel)
	return imovel, err
}


// func imovelComercialMint(c router.Context) (interface{}, error) {
// 	mintData := c.Param("mintData").(MintImovelComercialPayload)

// 	imovel := &ImovelComercial{
// 		Imovel: &Imovel{
// 			Inscricao: mintData.Inscricao,
// 			Nome:      mintData.Nome,
// 			Tamanho:   mintData.Tamanho,
// 		},
// 		Cnpj: mintData.Cnpj,
// 	}

// 	//TODO: se ja existe dar um erro
// 	///

// 	err = c.State().Insert(imovel)
// }

// func imovelResidencialMint(c router.Context) (interface{}, error) {
// 	mintData := c.Param("mintData").(MintImovelResidencialPayload)

// 	imovel := &ImovelResidencial{
// 		Imovel: &Imovel{
// 			Inscricao: mintData.Inscricao,
// 			Nome:      mintData.Nome,
// 			Tamanho:   mintData.Tamanho,
// 		},
// 		Cpf: mintData.Cpf,
// 	}

// 	//TODO: se ja existe dar um erro
// 	///

// 	err = c.State().Insert(imovel)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return imovel, nil
// }

func Get(c router.Context) (interface{}, error) {

	inscricao := c.ParamString("inscricao")
	imovel, err := c.State().Get(&Imovel{Inscricao: inscricao})

	if err != nil {
		return nil, fmt.Errorf("nao consegui pegar o imovel %s! %v", inscricao, err)
	}
	return imovel, nil
}

func SetValue(c router.Context) (interface{}, error) {

	inscricao := c.ParamString("inscricao")
	value := c.ParamInt("value")
	imovel, err := c.State().Get(&Imovel{Inscricao: inscricao})

	if err != nil {
		return nil, fmt.Errorf("nao consegui pegar o imovel %s! %v", inscricao, err)
	}
	im := imovel.(Imovel)
	im.Valor = value
	c.State().Put(im)

	return imovel, nil
}


func Init(c router.Context) (interface{}, error) {
	return true, nil
}
