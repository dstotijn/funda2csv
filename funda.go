package funda

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Object represents a building object (e.g. house, apartment) at Funda.
type Object struct {
	ID            string
	Address       string
	Price         int64
	URL           url.URL
	ImageURL      url.URL
	SurfaceArea   int
	NumberOfRooms int
}

// Objects is an array of Funda objects.
type Objects []*Object

// SearchResult represents search result data from Funda.
type SearchResult struct {
	// AccountStatus     int         `json:"AccountStatus"`
	// EmailNotConfirmed bool        `json:"EmailNotConfirmed"`
	// ValidationFailed  bool        `json:"ValidationFailed"`
	// ValidationReport  interface{} `json:"ValidationReport"`
	// Website           int         `json:"Website"`
	// Metadata struct {
	// ObjectType   string `json:"ObjectType"`
	// Omschrijving string `json:"Omschrijving"`
	// Titel        string `json:"Titel"`
	// } `json:"Metadata"`
	Objects []struct {
		// AangebodenSindsTekst        string        `json:"AangebodenSindsTekst"`
		// AanmeldDatum                string        `json:"AanmeldDatum"`
		// AantalBeschikbaar           interface{}   `json:"AantalBeschikbaar"`
		AantalKamers int `json:"AantalKamers"`
		// AantalKavels                interface{}   `json:"AantalKavels"`
		// Aanvaarding                 string        `json:"Aanvaarding"`
		Adres string `json:"Adres"`
		// Afstand                     int           `json:"Afstand"`
		// BronCode                    string        `json:"BronCode"`
		// ChildrenObjects             []interface{} `json:"ChildrenObjects"`
		// DatumAanvaarding            interface{}   `json:"DatumAanvaarding"`
		// DatumOndertekeningAkte      interface{}   `json:"DatumOndertekeningAkte"`
		// Foto                        string        `json:"Foto"`
		// FotoLarge string `json:"FotoLarge"`
		FotoLargest string `json:"FotoLargest"`
		// FotoMedium string `json:"FotoMedium"`
		// FotoSecure                  string        `json:"FotoSecure"`
		// GewijzigdDatum              interface{}   `json:"GewijzigdDatum"`
		// GlobalID                    int           `json:"GlobalId"`
		// GroupByObjectType           string        `json:"GroupByObjectType"`
		// Heeft360GradenFoto          bool          `json:"Heeft360GradenFoto"`
		// HeeftBrochure               bool          `json:"HeeftBrochure"`
		// HeeftOpenhuizenTopper       bool          `json:"HeeftOpenhuizenTopper"`
		// HeeftOverbruggingsgrarantie bool          `json:"HeeftOverbruggingsgrarantie"`
		// HeeftPlattegrond            bool          `json:"HeeftPlattegrond"`
		// HeeftTophuis                bool          `json:"HeeftTophuis"`
		// HeeftVeiling                bool          `json:"HeeftVeiling"`
		// HeeftVideo                  bool          `json:"HeeftVideo"`
		// HuurPrijsTot                interface{}   `json:"HuurPrijsTot"`
		// Huurprijs                   interface{}   `json:"Huurprijs"`
		// HuurprijsFormaat            interface{}   `json:"HuurprijsFormaat"`
		ID string `json:"Id"`
		// InUnitsVanaf                interface{}   `json:"InUnitsVanaf"`
		// IndProjectObjectType        bool          `json:"IndProjectObjectType"`
		// IndTransactieMakelaarTonen  interface{}   `json:"IndTransactieMakelaarTonen"`
		// IsSearchable                bool          `json:"IsSearchable"`
		// IsVerhuurd                  bool          `json:"IsVerhuurd"`
		// IsVerkocht                  bool          `json:"IsVerkocht"`
		// IsVerkochtOfVerhuurd        bool          `json:"IsVerkochtOfVerhuurd"`
		// Koopprijs                   int           `json:"Koopprijs"`
		// KoopprijsFormaat            string        `json:"KoopprijsFormaat"`
		// KoopprijsTot                int           `json:"KoopprijsTot"`
		// MakelaarID                  int           `json:"MakelaarId"`
		// MakelaarNaam                string        `json:"MakelaarNaam"`
		// MobileURL                   string        `json:"MobileURL"`
		// Note                        interface{}   `json:"Note"`
		// OpenHuis                    []string      `json:"OpenHuis"`
		// Oppervlakte int `json:"Oppervlakte"`
		// Perceeloppervlakte          interface{}   `json:"Perceeloppervlakte"`
		// Postcode string `json:"Postcode"`
		Prijs struct {
			// GeenExtraKosten     bool        `json:"GeenExtraKosten"`
			// HuurAbbreviation    string      `json:"HuurAbbreviation"`
			// Huurprijs           interface{} `json:"Huurprijs"`
			// HuurprijsOpAanvraag string      `json:"HuurprijsOpAanvraag"`
			// HuurprijsTot        interface{} `json:"HuurprijsTot"`
			KoopAbbreviation string `json:"KoopAbbreviation"`
			Koopprijs        int64  `json:"Koopprijs"`
			// KoopprijsOpAanvraag string      `json:"KoopprijsOpAanvraag"`
			// KoopprijsTot        int         `json:"KoopprijsTot"`
			// OriginelePrijs      interface{} `json:"OriginelePrijs"`
			// VeilingText         string      `json:"VeilingText"`
		} `json:"Prijs"`
		// PrijsGeformatteerdHTML     string   `json:"PrijsGeformatteerdHtml"`
		// PrijsGeformatteerdTextHuur string   `json:"PrijsGeformatteerdTextHuur"`
		// PrijsGeformatteerdTextKoop string   `json:"PrijsGeformatteerdTextKoop"`
		// Producten                  []string `json:"Producten"`
		// Project                    struct {
		// 	AantalKamersTotEnMet interface{}   `json:"AantalKamersTotEnMet"`
		// 	AantalKamersVan      interface{}   `json:"AantalKamersVan"`
		// 	AantalKavels         interface{}   `json:"AantalKavels"`
		// 	Adres                interface{}   `json:"Adres"`
		// 	FriendlyURL          interface{}   `json:"FriendlyUrl"`
		// 	GewijzigdDatum       interface{}   `json:"GewijzigdDatum"`
		// 	GlobalID             interface{}   `json:"GlobalId"`
		// 	HoofdFoto            string        `json:"HoofdFoto"`
		// 	IndIpix              bool          `json:"IndIpix"`
		// 	IndPDF               bool          `json:"IndPDF"`
		// 	IndPlattegrond       bool          `json:"IndPlattegrond"`
		// 	IndTop               bool          `json:"IndTop"`
		// 	IndVideo             bool          `json:"IndVideo"`
		// 	InternalID           string        `json:"InternalId"`
		// 	MaxWoonoppervlakte   interface{}   `json:"MaxWoonoppervlakte"`
		// 	MinWoonoppervlakte   interface{}   `json:"MinWoonoppervlakte"`
		// 	Naam                 interface{}   `json:"Naam"`
		// 	Omschrijving         interface{}   `json:"Omschrijving"`
		// 	OpenHuizen           []interface{} `json:"OpenHuizen"`
		// 	Plaats               interface{}   `json:"Plaats"`
		// 	Prijs                interface{}   `json:"Prijs"`
		// 	PrijsGeformatteerd   interface{}   `json:"PrijsGeformatteerd"`
		// 	PublicatieDatum      interface{}   `json:"PublicatieDatum"`
		// 	Type                 int           `json:"Type"`
		// 	Woningtypen          interface{}   `json:"Woningtypen"`
		// } `json:"Project"`
		// ProjectNaam interface{} `json:"ProjectNaam"`
		// PromoLabel  struct {
		// 	HasPromotionLabel     bool          `json:"HasPromotionLabel"`
		// 	PromotionPhotos       []interface{} `json:"PromotionPhotos"`
		// 	PromotionPhotosSecure interface{}   `json:"PromotionPhotosSecure"`
		// 	PromotionType         int           `json:"PromotionType"`
		// 	RibbonColor           int           `json:"RibbonColor"`
		// 	RibbonText            interface{}   `json:"RibbonText"`
		// 	Tagline               interface{}   `json:"Tagline"`
		// } `json:"PromoLabel"`
		// PublicatieDatum string `json:"PublicatieDatum"`
		// PublicatieStatus       int         `json:"PublicatieStatus"`
		// SavedDate              interface{} `json:"SavedDate"`
		// SoortAanbod            string      `json:"Soort-aanbod"`
		// SoortAanbod            int         `json:"SoortAanbod"`
		// StartOplevering        interface{} `json:"StartOplevering"`
		// TimeAgoText            interface{} `json:"TimeAgoText"`
		// TransactieAfmeldDatum  interface{} `json:"TransactieAfmeldDatum"`
		// TransactieMakelaarID   interface{} `json:"TransactieMakelaarId"`
		// TransactieMakelaarNaam interface{} `json:"TransactieMakelaarNaam"`
		// TypeProject            int         `json:"TypeProject"`
		URL string `json:"URL"`
		// VerkoopStatus string `json:"VerkoopStatus"`
		// WGS84X                 float64     `json:"WGS84_X"`
		// WGS84Y                 float64     `json:"WGS84_Y"`
		// WoonOppervlakteTot     int         `json:"WoonOppervlakteTot"`
		Woonoppervlakte int `json:"Woonoppervlakte"`
		// Woonplaats      string `json:"Woonplaats"`
		// ZoekType        []int  `json:"ZoekType"`
	} `json:"Objects"`
	Paging struct {
		AantalPaginas int `json:"AantalPaginas"`
		// HuidigePagina int         `json:"HuidigePagina"`
		// VolgendeURL   interface{} `json:"VolgendeUrl"`
		// VorigeURL     interface{} `json:"VorigeUrl"`
	} `json:"Paging"`
	TotaalAantalObjecten int `json:"TotaalAantalObjecten"`
}

// Search performs a search request at Funda.
func Search(token, opts string, page, pageSize int) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Funda/0 CFNetwork/887 Darwin/17.0.0")

	u, err := searchURL(token, opts, page, pageSize)
	if err != nil {
		return nil, err
	}
	req.URL = u

	log.Printf("Searching Funda for objects: %v", u.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected HTTP response code `%d` received",
			resp.StatusCode,
		)
	}

	return resp.Body, nil
}

// ObjectsFromSearchResult parses raw search result data into an Objects value.
func ObjectsFromSearchResult(r io.Reader) (objects Objects, pageCount int, err error) {
	var result SearchResult
	if err = json.NewDecoder(r).Decode(&result); err != nil {
		return
	}

	pageCount = result.Paging.AantalPaginas

	for _, o := range result.Objects {
		var houseURL, imageURL *url.URL
		object := &Object{}

		object.ID = o.ID
		object.Address = o.Adres
		object.Price = o.Prijs.Koopprijs
		object.SurfaceArea = o.Woonoppervlakte
		object.NumberOfRooms = o.AantalKamers

		houseURL, err = url.Parse(o.URL)
		if err != nil {
			log.Printf("funda: error parsing house URL: %s", err)
			return
		}
		object.URL = *houseURL

		imageURL, err = url.Parse(o.FotoLargest)
		if err != nil {
			log.Printf("funda: error parsing image URL: %s", err)
			return
		}
		object.ImageURL = *imageURL

		objects = append(objects, object)
	}

	return
}

func searchURL(token, searchOptions string, page, pageSize int) (*url.URL, error) {
	u, err := url.Parse("http://partnerapi.funda.nl/feeds/Aanbod.svc/search/json/" + token + "/")
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Set("website", "funda")
	q.Set("type", "koop")
	q.Set("zo", searchOptions)
	q.Set("page", strconv.Itoa(page))
	q.Set("pagesize", strconv.Itoa(pageSize))

	u.RawQuery = q.Encode()

	return u, nil
}
