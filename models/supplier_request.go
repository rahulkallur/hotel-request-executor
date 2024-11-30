package models

import "time"

// AvailabilityRequest represents the request for hotel availability search.
// It includes details about the stay, occupancy, hotels, filters, etc.
type AvailabilityRequest struct {
	Stay         Stay            `json:"stay"`         // The stay details (check-in and check-out dates).
	Occupancies  []Occupancy     `json:"occupancies"`  // The list of occupancy details (rooms, adults, children).
	Destination  Destination     `json:"destination"`  // The destination code (e.g., city or region).
	Hotels       HotelBedsHotels `json:"hotels"`       // The list of hotel IDs for availability search.
	Filter       Filter          `json:"filter"`       // Filter parameters for the availability search (e.g., max rooms).
	DailyRate    bool            `json:"dailyRate"`    // Flag to indicate if daily rates should be included.
	Platform     string          `json:"platform"`     // The platform where the request is made from.
	Packaging    bool            `json:"packaging"`    // Indicates if the request is for a packaged hotel.
	SourceMarket string          `json:"sourceMarket"` // The source market (e.g., country of the user).
	HotelPackage string          `json:"hotelPackage"` // The specific hotel package requested (if any).
	Review       []Review        `json:"review"`       // The review criteria for filtering hotels.
	Boards       Boards          `json:"boards"`       // Meal board options for the hotel search.
}

// Boards represents meal board options and whether they are included.
type Boards struct {
	Board    []string `json:"board"`    // List of meal boards (e.g., breakfast, half-board).
	Included bool     `json:"included"` // Indicates if the meal boards are included in the request.
}

// Destination represents the destination for the availability request, identified by a code.
type Destination struct {
	Code string `json:"Code"` // The destination code (e.g., city code).
}

// Filter represents the filtering criteria for the availability request.
type Filter struct {
	MaxRooms        int    `json:"maxRooms"`        // The maximum number of rooms to search for.
	MaxRatesPerRoom int    `json:"maxRatesPerRoom"` // The maximum number of rates per room to consider.
	PaymentType     string `json:"paymentType"`     // The type of payment method (e.g., prepayment, pay on arrival).
}

// HotelBedsHotels represents the list of hotel IDs to search for in the availability request.
type HotelBedsHotels struct {
	Hotel []int `json:"hotel"` // A list of hotel IDs to search for.
}

// Occupancy represents the occupancy details for a hotel room request, including the number of people.
type Occupancy struct {
	Rooms    int   `json:"rooms"`    // The number of rooms requested.
	Adults   int   `json:"adults"`   // The number of adults in the room.
	Children int   `json:"children"` // The number of children in the room.
	Paxes    []Pax `json:"paxes"`    // The list of Pax (guests) for the occupancy.
}

// Stay represents the stay details, including check-in and check-out dates.
type Stay struct {
	CheckinDate  time.Time `json:"checkIn"`  // The check-in date.
	CheckoutDate time.Time `json:"checkOut"` // The check-out date.
}

// Pax represents an individual guest (adult or child) in the room.
type Pax struct {
	RoomID  int    `json:"roomId"`  // The room ID associated with the guest.
	Type    string `json:"type"`    // The type of guest (e.g., "AD" for adult, "CH" for child).
	Age     int    `json:"age"`     // The age of the child (if applicable).
	Name    string `json:"name"`    // The first name of the guest.
	Surname string `json:"surname"` // The surname of the guest.
}

// Review represents the review criteria for a hotel, such as minimum rate and review count.
type Review struct {
	MinRate        int    `json:"minRate"`        // Minimum rate to filter the hotels.
	MaxRate        int    `json:"maxRate"`        // Maximum rate to filter the hotels.
	MinReviewCount int    `json:"minReviewCount"` // Minimum review count required for the hotel.
	Type           string `json:"type"`           // Type of review (e.g., "positive", "all").
}

// Resp represents the common response structure from the hotel supplier.
type Resp struct {
	SupplierResp string `json:"supplier_resp"` // Supplier-specific response (raw data or message).
	RoomInfo     string `json:"roominfo"`      // Room information in the supplier's response.
}

// CommonResp represents the common response structure for the availability request,
// which includes both the supplier request and room information.
type CommonResp struct {
	SupplierRequest string `json:"supplier_request"` // The request sent to the supplier (raw data).
	RoomInfo        string `json:"room_info"`        // Room information for the availability request.
}

// RoomInfo represents the information about a room, including details about adults, children, and the Pax.
type RoomInfo struct {
	Adult     int    `json:"adult"`      // Number of adults in the room.
	Child     int    `json:"child"`      // Number of children in the room.
	PaxKey    string `json:"paxkey"`     // A unique key to identify the Pax.
	TrackerID string `json:"tracker_id"` // Tracker ID associated with the booking or request.
	ChildAge  []int  `json:"childAge"`   // A list of children's ages (if applicable).
}
