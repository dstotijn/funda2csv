package funda

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFundaObjectsFromSearchResult(t *testing.T) {
	file, err := os.Open("test_data/funda_search_response.json")
	require.NoError(t, err)

	objects, pageCount, err := ObjectsFromSearchResult(file)
	require.NoError(t, err)

	assert.Equal(t, objects, Objects{
		&Object{
			ID:            "d113f0dd-4c05-4984-92ca-f7c739623dec",
			Address:       "Hoofdweg 99 - C",
			Price:         int64(300000),
			URL:           parseURL("http://www.funda.nl/koop/amsterdam/appartement-49397570-hoofdweg-99-c/"),
			ImageURL:      parseURL("http://cloud.funda.nl/valentina_media/085/371/511_grotere.jpg"),
			SurfaceArea:   65,
			NumberOfRooms: 3,
		},
		&Object{
			ID:            "27097d93-3547-47a3-ac9f-b16e1112e9ad",
			Address:       "Geuzenstraat 77 III",
			Price:         int64(250000),
			URL:           parseURL("http://www.funda.nl/koop/amsterdam/appartement-49397476-geuzenstraat-77-iii/"),
			ImageURL:      parseURL("http://cloud.funda.nl/valentina_media/085/368/478_grotere.jpg"),
			SurfaceArea:   48,
			NumberOfRooms: 3,
		},
	})

	assert.Equal(t, pageCount, 4)
}

func TestFundaSearchURL(t *testing.T) {
	exp := parseURL("http://partnerapi.funda.nl/feeds/Aanbod.svc/search/json/foobar/?page=1&pagesize=10&type=koop&website=funda&zo=%2Famsterdam%2F1-dag")

	got, err := searchURL("foobar", "/amsterdam/1-dag", 1, 10)
	require.NoError(t, err)

	assert.Equal(t, exp, *got)
}

func parseURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	return *u
}
