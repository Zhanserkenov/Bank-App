package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"bank-app/helpers"
	"bank-app/transactions"
	"bank-app/useraccounts"
	"bank-app/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type TransactionBody struct {
	UserId   uint
	From     uint
	To       uint
	Amount   int
	Currency string
}

type ConversionRequest struct {
	UserID       uint
	FromCurrency string
	ToCurrency   string
	Amount       uint
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "all is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	login := users.Login(formattedBody.Username, formattedBody.Password)
	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	apiResponse(register, w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

func getMyTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]
	auth := r.Header.Get("Authorization")

	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sortBy := requestBody["sortBy"]
	currency := requestBody["currency"]
	pageStr := requestBody["page"]
	pageSizeStr := requestBody["pageSize"]

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10 // Default page size
	}

	offset := (page - 1) * pageSize

	transactions := transactions.GetMyTransactions(userId, auth, sortBy, currency, offset, pageSize)
	apiResponse(transactions, w)
}

func transaction(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)
	auth := r.Header.Get("Authorization")

	var formattedBody TransactionBody
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	if formattedBody.UserId == 0 || formattedBody.From == 0 || formattedBody.To == 0 || formattedBody.Amount == 0 || formattedBody.Currency == "" {
		apiResponse(map[string]interface{}{"message": "Missing required fields"}, w)
		return
	}

	transaction := useraccounts.Transaction(formattedBody.UserId, formattedBody.From, formattedBody.To, formattedBody.Amount, formattedBody.Currency, auth)
	apiResponse(transaction, w)
}

func convert(w http.ResponseWriter, r *http.Request) {

	body := readBody(r)
	auth := r.Header.Get("Authorization")

	var formattedBody ConversionRequest
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	if !helpers.ValidateToken(fmt.Sprint(formattedBody.UserID), auth) {
		apiResponse(map[string]interface{}{"message": "Not valid token"}, w)
		return
	}

	convertedAmount := useraccounts.ConvertCurrency(formattedBody.UserID, formattedBody.FromCurrency, formattedBody.ToCurrency, formattedBody.Amount)

	response := map[string]interface{}{
		"message": "Conversion successful",
		"data": map[string]interface{}{
			"amount":           formattedBody.Amount,
			"from_currency":    formattedBody.FromCurrency,
			"to_currency":      formattedBody.ToCurrency,
			"converted_amount": convertedAmount,
		},
	}
	apiResponse(response, w)
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	router.HandleFunc("/transactions/{userID}", getMyTransactions).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/convert", convert).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
