package merchants

import (
	"net/http"
	"strings"

	"kiplingkelvin/golang-skeleton/internal/common/utils"
	"kiplingkelvin/golang-skeleton/internal/merchants/models"
	merchant_models "kiplingkelvin/golang-skeleton/internal/merchants/models"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

func (p *Payload) MerchantRegistrationHandler(w http.ResponseWriter, r *http.Request) {


	var request merchant_models.MerchantRegisterRequest
	utils.UnmarshallJSONFromRequest(w, r, &request)

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

	phoneNumber := strings.ReplaceAll(request.PhoneNumber, " ", "")
	merchant := models.Merchant{
		BusinessEmail:      strings.ToLower(request.Email),
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		CompanyName:        request.Company,
		MobileNumber:       "+2547" + phoneNumber[len(phoneNumber)-8:],
		Password:           []byte(request.Password),
		IsShopifyActive:    request.IsShopifyActive,
		ShopifyCode:        request.ShopifyCode,
		ShopifyDomain:      request.ShopifyDomain,
		ShopifyAccessToken: request.ShopifyAccessToken,
	}

	_, err = p.DAO.Postgres.Create(r.Context(), merchant)
	if err != nil {
		logrus.WithError(err).Logger.Error("creating merchant db error")
		utils.WriteErrorResponse(w, utils.Response{
			Message: err.Error(),
			Success: false,
			Status:  http.StatusInternalServerError,
		}, utils.ContentTypeJSON, http.StatusInternalServerError)
		return
	}

	response := merchant_models.MerchantRegisterResponse{
		Response: utils.Response{
			Message: "sign up successful",
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

func (p *Payload) ProfileUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func (p *Payload) ProfileGetHandler(w http.ResponseWriter, r *http.Request) {

}
