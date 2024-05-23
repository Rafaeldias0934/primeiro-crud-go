package rest_err

import (
	"net/http"
)

// Criação de um objeto do tipo Struct com os campos desejaveis:  OBS: Precisamos que nossas objeto seja exportavel para JSON "json:"message"`"
type RestErr struct {
	Message string   `json:"message"`          //Este campo servirá para infomar ao usuário qual foi o erro que deu detro da aplicação
	Err     string   `json:"error"`            // "Stack trace" rastreamento do erro que deu para colpeta descrição oa usuário
	Code    int      `json:"code"`             // Irá guardar qual o código da requisição que vamos estar retornando ao cliente
	Causes  []Causes `json:"causes,omitempty"` // Outro objeto Causes que basicamente vai criar quais as causas de erros que deu dentro da nossa aplicação
}

/*
	abaixo criamos o objeto Causes do tipo struct para padronizar uma Array de um objeto,

que servirá para que quando usuário mandar uma senha errada, terá uma lista informando "Campo - Campo senha está incorreto"
*/
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Função criada que será implementada a partir de *RestErr, que terá o nome de Error para satisfazer a interface de erro do Go
func (r *RestErr) Error() string {
	return r.Message
}

// Criação de um contrutor de objetos de erros para receber nosso compos
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// método criado para facilitar a programação de erros como um todo
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

/*
	Função NewBadRequestValidationError criada para caso o usuário chame o nosso objeto/ função com o valor errado

dentro desta função inserimos o Causes o array de causas como parâmetro e como return
*/
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// Está função é para padronizar o internalServerError que acontecer dentro da nossa aplicaçao
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal_Server_error",
		Code:    http.StatusInternalServerError,
	}
}

// Está função é para padronizar o not_found que acontecer dentro da nossa aplicaçao
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

// Está função é para padronizar o forbidden que acontecer dentro da nossa aplicaçao
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
