package main

import (
	"encoding/json"
	"net/http"
)

type TransactionType int64

const (
	Home TransactionType = iota
	Food
	Services
	Clothing
	Health
	Transports
	Entertainment
	Miscellaneous
)

func (t TransactionType) String() string {
	switch t {
	case Home:
		return "home"
	case Food:
		return "food"
	case Services:
		return "services"
	case Clothing:
		return "cloathing"
	case Health:
		return "health"
	case Transports:
		return "transports"
	case Entertainment:
		return "entertainment"
	case Miscellaneous:
		return "miscellaneous"
	}
	return "unknown"
}

type Transaction struct {
	Id           string          `json:"id"`
	Type         TransactionType `json:"type"`
	Descritpion  string          `json:"description"`
	Value        float32         `json:"value"`
	Date         string          `json:"date"`
	CreatedAt    string          `json:"created_at"`
	LastModified string          `json:"last_modified"`
}

type transactionHandler struct {
	store map[string]Transaction
}

func (h *transactionHandler) get(w http.ResponseWriter, r *http.Request) {
	transactions := make([]Transaction, len(h.store))

	i := 0
	for _, transaction := range h.store {
		transactions[i] = transaction
		i++
	}

	jsonBytes, err := json.Marshal(transactions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newTransactionHandler() *transactionHandler {
	return &transactionHandler{
		store: map[string]Transaction{
			"1": {
				Id:           "0",
				Type:         Home,
				Descritpion:  "tachipirina",
				Value:        -10,
				Date:         "01-01-2022",
				CreatedAt:    "01-01-2022",
				LastModified: "01-01-2022",
			},
		},
	}
}

func main() {
	transactionHandler := newTransactionHandler()
	http.HandleFunc("/transactions", transactionHandler.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
