package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	bookingListURL   = "https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList"
	jenisKonsumsiURL = "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"
)

type ListConsumptionItem struct {
	Name string `json:"name"`
}
type Booking struct {
	ID              string                `json:"id"`
	RoomName        string                `json:"roomName"`
	OfficeName      string                `json:"officeName"`
	BookingDate     string                `json:"bookingDate"`
	EndTime         string                `json:"endTime"`
	Participants    int                   `json:"participants"`
	ListConsumption []ListConsumptionItem `json:"listConsumption"`
}
type JenisKonsumsi struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	MaxPrice int    `json:"maxPrice"`
}
type DetailConsumptionItem struct {
	SnackSiang int `json:"snackSiang"`
	MakanSiang int `json:"makanSiang"`
	SnackSore  int `json:"snackSore"`
}
type Summary struct {
	Participants        int                   `json:"participants"`
	ConsumptionCount    int                   `json:"consumptionCount"`
	TotalBookingRoom    int                   `json:"totalBookingRoom"`
	TotalPercentageRoom float64               `json:"totalPercentageRoom"`
	Total               int                   `json:"total"`
	DetailConsumption   DetailConsumptionItem `json:"detailConsumption"`
}
func getBookings() ([]Booking, error) {
	resp, err := http.Get(bookingListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var bookings []Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func getJenisKonsumsi() ([]JenisKonsumsi, error) {
	resp, err := http.Get(jenisKonsumsiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var jenisKonsumsi []JenisKonsumsi
	if err := json.NewDecoder(resp.Body).Decode(&jenisKonsumsi); err != nil {
		return nil, err
	}
	return jenisKonsumsi, nil
}
func filterPrice(name string, jenisKonsumsi []JenisKonsumsi) int {
	for _, konsumsi := range jenisKonsumsi {
		if name == konsumsi.Name {
			return konsumsi.MaxPrice
		}
	}
	return 0
}

// filter officename
func summarizeBookings(bookings []Booking, jenisKonsumsi []JenisKonsumsi) map[string]map[string]Summary {
	summary := make(map[string]map[string]Summary)
	for _, booking := range bookings {
		office := booking.OfficeName
		room := booking.RoomName
		if _, ok := summary[office]; !ok {
			summary[office] = make(map[string]Summary)
		}
		roomSummary, exists := summary[office][room]
		if !exists {
			roomSummary = Summary{
				Participants:     0,
				ConsumptionCount: 0,
			}
		}
		roomSummary.Participants += booking.Participants
		roomSummary.ConsumptionCount += len(booking.ListConsumption)
		roomSummary.TotalBookingRoom += 1
		percentage := float64(roomSummary.TotalBookingRoom) / float64(len(bookings))
		if percentage > 0 {
			roomSummary.TotalPercentageRoom = math.Round(percentage * 100)
		}
		// hitung price masing2 room
		for _, price := range booking.ListConsumption {
			result := filterPrice(price.Name, jenisKonsumsi)
			roomSummary.Total += result
			if result > 0 {
				if price.Name == "Makan Siang" {
					roomSummary.DetailConsumption.MakanSiang += booking.Participants
				} else if price.Name == "Snack Siang" {
					roomSummary.DetailConsumption.SnackSiang += booking.Participants
				} else if price.Name == "Snack Sore" {
					roomSummary.DetailConsumption.SnackSore += booking.Participants
				}
			}
		}
		summary[office][room] = roomSummary
	}
	return summary
}
func main() {
soal1()
soal2()
soal3()
soal4()
soal5()
	r := chi.NewRouter()
	r.Get("/summary", func(w http.ResponseWriter, r *http.Request) {
		bookings, err := getBookings()
		jenisKonsumsi, _ := getJenisKonsumsi()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bookingGroup := summarizeBookings(bookings, jenisKonsumsi)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookingGroup)
	})

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}