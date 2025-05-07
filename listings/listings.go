package listings

import (
	"encoding/json"
	"maps"
	"net/http"
	"os"
	"slices"
	"sort"

	"github.com/go-playground/validator/v10"
	"github.com/ptrj96/go-car-storage-api/logging"
)

type Listing struct {
	Id         string `json:"id"`
	LocationId string `json:"location_id"`
	Length     int    `json:"length"`
	Width      int    `json:"width"`
	PriceCents int    `json:"price_in_cents"`
}

type CarRequest struct {
	Length   int `json:"length" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
}

type LocationResponse struct {
	LocationId      string   `json:"location_id"`
	ListingIds      []string `json:"listing_ids"`
	TotalPriceCents int      `json:"total_price_in_cents"`
}

var listings []Listing

func GetListings() ([]Listing, error) {
	logger := logging.GetLogger()
	if len(listings) == 0 {
		logger.Print("listings not loaded yet, loading now")
		listingsFile, err := os.Open("listings.json")
		if err != nil {
			logger.Printf("error opening file: %s", err)
			return []Listing{}, err
		}

		jsonParser := json.NewDecoder(listingsFile)
		if err = jsonParser.Decode(&listings); err != nil {
			logger.Printf("error decoding json from file: %s", err)
		}
	}

	return listings, nil
}

func FindListingsHandler(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()

	logger.Print("hit listings endpoint")
	var cars []CarRequest
	if err := json.NewDecoder(r.Body).Decode(&cars); err != nil {
		logger.Printf("error unmarshalling json: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"error unmarshalling json: ` + string(err.Error()) + `"}`))
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	for _, car := range cars {
		if err := validate.Struct(car); err != nil {
			logger.Printf("error validating payload: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"error with json validation: ` + string(err.Error()) + `"}`))
			return
		}
	}

	listings, err := GetListings()
	if err != nil {
		return
	}

	locations := make(map[string]LocationResponse)

	var testCars []int
	for _, car := range cars {
		for i := 0; i < car.Quantity; i++ {
			testCars = append(testCars, car.Length)
		}
	}

	for _, listing := range listings {
		if CheckListingFit(testCars, listing) {
			if location, ok := locations[listing.LocationId]; ok {
				location.ListingIds = append(location.ListingIds, listing.Id)
				location.TotalPriceCents += listing.PriceCents
			} else {
				locations[listing.LocationId] = LocationResponse{
					LocationId:      listing.LocationId,
					ListingIds:      []string{listing.Id},
					TotalPriceCents: listing.PriceCents,
				}
			}

		}
	}

	jsonLocations, err := json.Marshal(slices.Collect(maps.Values(locations)))
	if err != nil {
		logger.Printf("error marshalling locations: %s", err)
	}

	logger.Printf("found %d valid locations", len(slices.Collect(maps.Values(locations))))
	w.WriteHeader(http.StatusOK)
	w.Write(jsonLocations)
}

func CheckListingFit(cars []int, listing Listing) bool {
	var listingSlice []int
	for i := 0; i < listing.Width/10; i++ {
		listingSlice = append(listingSlice, listing.Length)
	}

	slices.Reverse(sort.IntSlice(cars))
	numCarsFit := 0

	for _, car := range cars {
		for i, spot := range listingSlice {
			if car <= spot {
				listingSlice[i] -= car
				numCarsFit++
				break
			}
		}
	}

	return numCarsFit == len(cars)
}
