package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Get token from the environment variable
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}

	// Get payment provider token from the environment variable
	paymentToken := os.Getenv("PAY_TOKEN")
	if paymentToken == "" {
		panic("PAY_TOKEN environment variable is empty")
	}

	var repo Repo = NewInMemoryRepo()

	// Setup new HTTP server mux to handle different paths
	mux := httprouter.New()

	// Handle OPTIONS (preflight) requests
	mux.HandleOPTIONS = true
	mux.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r)

		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Headers", "Content-Type, Ngrok-Skip-Browser-Warning, Init-Data")
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	// This serves /goods endpoint - list of shop items
	mux.HandlerFunc(http.MethodGet, "/goods", goods(repo))
	// This serves /order endpoit - place order
	mux.HandlerFunc(http.MethodPost, "/order", order(repo, token, paymentToken))

	server := http.Server{
		Addr:    ":4001",
		Handler: mux,
	}

	// Start the webserver
	if err := server.ListenAndServe(); err != nil {
		panic("failed to listen and serve: " + err.Error())
	}
}

// goods loads list of shop items from the repository and returns it as json payload
func goods(repo Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)

		// Set CORS header
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// load items via repo
		g := repo.LoadItems()

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// order parses order from json body, stores it via repository and creates invoice link via Bot API
func order(repo Repo, token string, paymentToken string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)

		// Set CORS header
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var o []OrderItem

		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(o)

		// persist order via repo
		orderNum := repo.StoreOrder(o)

		// create invoice link in Telegram (using createInvoiceLink from Bot API)
		// it's kinda strange to format the client UI in here, also to duplicate the bot functionality
		var i struct {
			Title               string         `json:"title"`
			Description         string         `json:"description"`
			Payload             string         `json:"payload"`
			ProviderToken       string         `json:"provider_token"`
			Currency            string         `json:"currency"`
			Prices              []LabeledPrice `json:"prices"`
			NeedPhoneNumber     bool           `json:"need_phone_number"`
			NeedShippingAddress bool           `json:"need_shipping_address"`
		}

		i.Title = fmt.Sprintf("Order #%d", orderNum)
		i.Description = "PPAP"
		i.Payload = "payload"
		i.ProviderToken = paymentToken
		i.Currency = "USD"
		i.Prices = make([]LabeledPrice, 0, len(o))
		for _, it := range o {
			item := repo.FindItem(it.Id)
			if item == nil {
				i.Prices = append(i.Prices, LabeledPrice{fmt.Sprintf("#%d", it.Id), 0})
				continue
			}
			i.Prices = append(i.Prices, LabeledPrice{fmt.Sprintf("%s %dx", item.Name, it.Qty), it.Qty * item.Price * 100})
		}
		i.NeedPhoneNumber = true
		i.NeedShippingAddress = true

		ie, err := json.Marshal(i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reqBody := bytes.NewReader(ie)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/createInvoiceLink", token), reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// use Bot API response as ours
		w.Header().Set("Content-Type", "application/json")

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// refs miniapp.Item
type Item struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Pic    string `json:"pic"`
	PicAlt string `json:"picAlt"`
}

// refs miniapp.OrderItem
type OrderItem struct {
	Id  int `json:"id"`
	Qty int `json:"qty"`
}

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// Repo desribes data repository
type Repo interface {
	LoadItems() []Item
	FindItem(id int) *Item
	StoreOrder(items []OrderItem) int
}
