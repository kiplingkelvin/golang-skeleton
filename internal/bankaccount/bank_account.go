package bankaccount

import (
	bank_accounts_models "kiplingkelvin/golang-skeleton/internal/bankaccount/models"
	"kiplingkelvin/golang-skeleton/internal/common/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

func (p *Payload) BankAccountCreateHandler(w http.ResponseWriter, r *http.Request) {

	var request bank_accounts_models.BankAccount
	utils.UnmarshallJSONFromRequest(w, r, &request)
	logrus.Info("Merchant registration request received. ", request)



	// Validate struct fields
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		logrus.Error("Validation error. " + err.Error())
		utils.WriteErrorResponse(w, utils.Response{
			Message: "Invalid payload ",
			Success: false,
			Status:  http.StatusBadRequest,
		}, utils.ContentTypeJSON, http.StatusBadRequest)
		return
	}



	_, err = p.DAO.Postgres.Create(r.Context(), request)
	if err != nil {
		logrus.WithError(err).Logger.Error("creating merchant db error")
		utils.WriteErrorResponse(w, utils.Response{
			Message: err.Error(),
			Success: false,
			Status:  http.StatusInternalServerError,
		}, utils.ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	response := bank_accounts_models.BankAccountCreateResponse{
		Response: utils.Response{
			Message: "bank account created successfully",
			Status:  http.StatusCreated,
			Success: true,
		},
	}

	err = utils.WriteHTTPResponse(w, response, utils.ContentTypeJSON, http.StatusCreated)
	if err != nil {
		logrus.WithError(err).Logger.Error("writeHTTPResponse error")

		utils.WriteErrorResponse(w, utils.Response{
			Message: "error writing http response. " + err.Error(),
			Success: false,
			Status:  http.StatusInternalServerError,
		}, utils.ContentTypeJSON, http.StatusInternalServerError)
		return
	}

}

func (p *Payload) BankAccountUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func (p *Payload) BankAccountGetHandler(w http.ResponseWriter, r *http.Request) {

}