package validations

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {

}
