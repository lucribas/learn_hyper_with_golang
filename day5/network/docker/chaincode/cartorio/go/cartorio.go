package cartorio

import (
	"github.com/s7techlab/cckit/extensions/owner"
	"github.com/s7techlab/cckit/router"
	"github.com/s7techlab/cckit/router/param"
)

//-------------------------------------------------------

func NewCC() *router.Chaincode {
	r := router.New(`cartorio`)
	r.Init(invokeInit)

	r.Group(`imovel`)

	// Invoke
	r.Invoke(`imovelMint`, imovelMint, param.Struct("mintData", &MintImovelPayload{}))
	// r.Invoke(`imovelComercialMint`, imovelComercialMint, param.Struct("mintData", &MintImovelComercialPayload{}))
	// r.Invoke(`imovelResidencialMint`, imovelResidencialMint, param.Struct("mintData", &MintImovelResidencialPayload{}))

	// Get method
	r.Query(`imovelGet`, Get, param.String("inscricao"))

	// Set method
	r.Query(`imovelSetValue`, SetValue, param.String("inscricao"), param.String("value"))

	// Query(`imovelSetCpnj`, SetCnpj, param.String("inscricao"), param.String("cnpj")).
	// Query(`imovelSetCpf`, SetCpf, param.String("inscricao"), param.String("cpf")).

	return router.NewChaincode(r)
}

// ======= Init ==================
func invokeInit(c router.Context) (interface{}, error) {
	Init(c)
	return owner.SetFromCreator(c)
}
