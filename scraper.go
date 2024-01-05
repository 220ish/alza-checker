package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"
)

type AlzaSearchResponse struct {
	ShareURL           string  `json:"shareURL"`
	TotalUnfiltered    int     `json:"totalUnfiltered"`
	Err                int     `json:"err"`
	FavCnt             int     `json:"favCnt"`
	BasketCnt          int     `json:"basket_cnt"`
	BasketTotalCnt     int     `json:"basket_total_cnt"`
	Msg                string  `json:"msg"`
	UserName           string  `json:"user_name"`
	UserID             int     `json:"user_id"`
	Vzt                int     `json:"vzt"`
	CountryID          int     `json:"countryID"`
	CountryPhonePrefix string  `json:"countryPhonePrefix"`
	PremiumBooks       int     `json:"premiumBooks"`
	PremiumDelivery    int     `json:"premiumDelivery"`
	PremiumValidTo     string  `json:"premiumValidTo"`
	PremiumRenew       bool    `json:"premiumRenew"`
	ServerTime         int     `json:"serverTime"`
	DisplayCnt         int     `json:"display_cnt"`
	Data1Cnt           int     `json:"data1_cnt"`
	Data2Cnt           int     `json:"data2_cnt"`
	Data2              []Data2 `json:"data2"`
	Total              int     `json:"total"`
	HasNext            bool    `json:"has_next"`
}

type Data2 struct {
	Self struct {
		Href    string `json:"href"`
		AppLink string `json:"appLink"`
		Enabled bool   `json:"enabled"`
	} `json:"self"`
	PriceInfo struct {
		PriceWithoutVat            string  `json:"priceWithoutVat"`
		PriceWithVat               string  `json:"priceWithVat"`
		ComparePrice               string  `json:"comparePrice"`
		DiscountRate               string  `json:"discountRate"`
		PricePrefix                string  `json:"pricePrefix"`
		PricePostfix               string  `json:"pricePostfix"`
		DiscountReason             string  `json:"discountReason"`
		GaPriceWithoutVat          float64 `json:"gaPriceWithoutVat"`
		PriceNoCurrency            float64 `json:"priceNoCurrency"`
		DelayedPaymentPriceWithVat string  `json:"delayedPaymentPriceWithVat"`
	} `json:"priceInfo"`
	PriceInfoV2 struct {
		PriceWithVat              string  `json:"priceWithVat"`
		ComparePrice              string  `json:"comparePrice"`
		PricePrefix               string  `json:"pricePrefix"`
		PricePostfix              string  `json:"pricePostfix"`
		PriceDescription          string  `json:"priceDescription"`
		PriceType                 int     `json:"priceType"`
		HeaderText                string  `json:"headerText"`
		HeaderIconUrl             string  `json:"headerIconUrl"`
		PriceSave                 string  `json:"priceSave"`
		ExplanationPriceAction    string  `json:"explanationPriceAction"`
		UnitPriceWithVat          string  `json:"unitPriceWithVat"`
		UnitName                  string  `json:"unitName"`
		PriceWithoutVatNoCurrency float64 `json:"priceWithoutVatNoCurrency"`
		PriceNoCurrency           float64 `json:"priceNoCurrency"`
	} `json:"priceInfoV2"`
	CpriceTotal              interface{} `json:"cpriceTotal"`
	IsComparable             bool        `json:"is_comparable"`
	PriceTotal               interface{} `json:"priceTotal"`
	PromosWorth              interface{} `json:"promosWorth"`
	RatingCount              int         `json:"ratingCount"`
	SpecParent               interface{} `json:"specParent"`
	UserOwns                 bool        `json:"userOwns"`
	UserOwnedContentID       interface{} `json:"userOwnedContentId"`
	NavigationURL            interface{} `json:"navigationUrl"`
	CanChangeQuantity        bool        `json:"canChangeQuantity"`
	CanCashBack              bool        `json:"canCashBack"`
	CashBackType             int         `json:"cashBackType"`
	CashBackPriceLabel       interface{} `json:"cashBackPriceLabel"`
	CashBackPrice            interface{} `json:"cashBackPrice"`
	CashBackPercent          interface{} `json:"cashBackPercent"`
	CashBackPriceDescription interface{} `json:"cashBackPriceDescription"`
	CashBackPromoActions     interface{} `json:"cashBackPromoActions"`
	GiftAdvisor              interface{} `json:"giftAdvisor"`
	CatalogNumber            interface{} `json:"catalog_number"`
	ActionCategoryIconsURL   []string    `json:"actionCategoryIconsUrl"`
	Count                    int         `json:"count"`
	ShowUpsellDialog         bool        `json:"showUpsellDialog"`
	OnChangeItemCountClick   interface{} `json:"onChangeItemCountClick"`
	Icons                    []struct {
		Text        string      `json:"text"`
		Image       string      `json:"image"`
		ClickAction interface{} `json:"clickAction"`
	} `json:"icons"`
	ViewerProductInfo struct {
		Href    string `json:"href"`
		AppLink string `json:"appLink"`
		Enabled bool   `json:"enabled"`
	} `json:"viewerProductInfo"`
	DeliveryPromo            interface{} `json:"deliveryPromo"`
	PersonalizedAvailability interface{} `json:"personalizedAvailability"`
	ID                       int         `json:"id"`
	Code                     string      `json:"code"`
	Img                      string      `json:"img"`
	Name                     string      `json:"name"`
	Spec                     string      `json:"spec"`
	Price                    string      `json:"price"`
	Cprice                   interface{} `json:"cprice"`
	PriceWithoutVat          string      `json:"priceWithoutVat"`
	Avail                    string      `json:"avail"`
	AvailPostfix             interface{} `json:"avail_postfix"`
	AvailPostfix2            interface{} `json:"avail_postfix2"`
	AvailLegend              interface{} `json:"availLegend"`
	AvailColor               string      `json:"avail_color"`
	IsAction                 bool        `json:"is_action"`
	ActionName               string      `json:"action_name"`
	Rating                   float64     `json:"rating"`
	PromoCnt                 int         `json:"promo_cnt"`
	Promos                   interface{} `json:"promos"`
	Order                    int         `json:"order"`
	IsSpecialService         bool        `json:"is_special_service"`
	Type                     int         `json:"type"`
	Year                     int         `json:"year"`
	CanBuy                   bool        `json:"can_buy"`
	ItemType                 string      `json:"itemType"`
	OrderItemID              interface{} `json:"orderItemId"`
	IType                    int         `json:"iType"`
	MaxCanBuy                int         `json:"maxCanBuy"`
	EshopType                int         `json:"eshopType"`
	URL                      string      `json:"url"`
	CanStream                bool        `json:"canStream"`
	CanUserStream            bool        `json:"canUserStream"`
	SupplierCode             string      `json:"supplierCode"`
	ParentID                 int         `json:"parentId"`
	MinimumAmount            float64     `json:"minimumAmount"`
	AmountInPack             float64     `json:"amountInPack"`
	StartTime                interface{} `json:"start_time"`
	EndTime                  interface{} `json:"end_time"`
	VariantType              int         `json:"variant_type"`
	PriceNoCurrency          float64     `json:"priceNoCurrency"`
	CategoryName             interface{} `json:"categoryName"`
	InBasket                 float64     `json:"inBasket"`
}

func main() {

	scraper()

	// trigger every 30 seconds
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			scraper()
		}
	}
}

var searchTerm = "iphone 15 pro"

func scraper() {
	webhook := "https://webhooks.slack.com/7io-n16ldfXbBHMEqYb7ag0XoNDgHoG0Xe"
	url := "https://www.alza.sk/Services/RestService.svc/v5/search?density=3.0&country=sk"

	headers := map[string]string{
		"Content-Type": "application/json",
		"Cookie":       "__cf_bm=TQm_dv86wetqlru9zFoEOzsfmaXxEPZ1looQcwjFenk-1704495804-1-AWQhCPbltEx4imWqMdmg7lIL+alS0NA+jDJjXT3DD/S4uV2EmMK3tIZvbT52+3CwHTENq26ZikKDqzfx/AvZIYs=; lb_id=b86049bac5204f1379731b7cb58b6924; VST=7610ed6d-1eac-ee11-843d-0c42a19546e5; VZTX=8769875458; density=3.0; platform=ios",
		"User-Agent":   "Alza/133.0 (cz.juicymo.contracts.ios.Alza-01; build:202312010824; iOS 17.2.0) Alamofire/5.2.1",
	}

	payloadBytes, err := json.Marshal(map[string]interface{}{
		"type":             "CATEGORY",
		"typeId":           0,
		"sendPrices":       false,
		"orderBy":          1,
		"searchTerm":       searchTerm,
		"params":           []interface{}{},
		"selectedBranches": []interface{}{},
		"id":               "0",
		"newsOnly":         false,
		"availabilityType": 0,
		"page":             1,
		"wearType":         0,
	})
	if err != nil {
		fmt.Println("Error encoding payload:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	var newResponse AlzaSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&newResponse)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	existingJSON, err := ioutil.ReadFile("current.json")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading current.json:", err)
		return
	}

	if len(existingJSON) > 0 {
		var existingResponse AlzaSearchResponse
		err := json.Unmarshal(existingJSON, &existingResponse)
		if err != nil {
			fmt.Println("Error unmarshalling existing JSON:", err)
			return
		}

		if product_comparison(existingResponse.Data2, newResponse.Data2) {
			fmt.Printf("%s — alza.sk — %s — changes detected\n", time.Now().Format("15:04:05"), searchTerm)

			err := send_discord_webhook(webhook)
			if err != nil {
				fmt.Println("Discord webhook sent successfully.")
			} else {
				fmt.Println("Discord webhook sent successfully.")
			}

		} else {
			fmt.Printf("%s — alza.sk — %s — checking for inventory changes\n", time.Now().Format("15:04:05"), searchTerm)
		}
	}

	responseBytes, err := json.MarshalIndent(newResponse, "", "  ")
	if err != nil {
		fmt.Println("Error encoding response struct:", err)
		return
	}
	err = ioutil.WriteFile("current.json", responseBytes, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func product_comparison(existingData2, newData2 []Data2) bool {
	if len(existingData2) != len(newData2) {
		return true
	}

	for i, existingItem := range existingData2 {
		newItem := newData2[i]

		if existingItem.Name != newItem.Name || existingItem.Price != newItem.Price {
			return true
		}
	}

	return false
}

func approxEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

type DiscordWebhookEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func send_discord_webhook(webhookURL string) error {
	embed := DiscordWebhookEmbed{
		Title:       "Change Detected — Alza.sk",
		Description: "Change has been detected for `" + searchTerm + "` search term.",
	}

	payload := map[string]interface{}{
		"embeds": []DiscordWebhookEmbed{embed},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}
