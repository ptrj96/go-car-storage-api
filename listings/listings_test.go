package listings_test

import (
	"testing"

	"github.com/ptrj96/go-car-storage-api/listings"
)

var testCars [][]int = [][]int{
	{
		10, 20, 20, 25,
	},
	{
		100,
	},
	{
		10,
	},
	{
		15, 20, 25, 5, 10,
	},
	{
		50, 5,
	},
}

var testListings []listings.Listing = []listings.Listing{
	{

		Id:         "2f9266ce-7716-40b1-b27f-c1d77a807551",
		LocationId: "d1c331f1-9ae6-4d8a-9d87-a0cf5cfe1536",
		Length:     40,
		Width:      20,
		PriceCents: 64683,
	},
	{
		Id:         "741a0213-8512-499e-a8f2-8d3aad664d76",
		LocationId: "760ec3ad-5db1-4e81-820b-f15f009d4b5a",
		Length:     20,
		Width:      20,
		PriceCents: 16293,
	},
	{
		Id:         "2213d790-9641-4ac1-b56f-9883ac54ec1a",
		LocationId: "b6893aac-fb26-416e-92c7-1c50d95463ce",
		Length:     30,
		Width:      20,
		PriceCents: 7160,
	},
	{
		Id:         "5bea2c67-7df4-4ebb-ac7c-c2f4e9c25e43",
		LocationId: "d54ac3d1-5ab8-4e4f-97eb-4b1078313042",
		Length:     20,
		Width:      10,
		PriceCents: 91115,
	},
}

var carNumFitResults []int = []int{1, 0, 4, 1, 0}

func TestCheckListingFit(t *testing.T) {
	for i, cars := range testCars {
		numFit := 0

		for _, listing := range testListings {
			if listings.CheckListingFit(cars, listing) {
				numFit++
			}
		}

		if carNumFitResults[i] != numFit {
			t.Errorf("incorrect number of fitments: expected %d, got %d", carNumFitResults[i], numFit)
		}
	}
}
