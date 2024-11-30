package models

import "time"

// AutoGenerated contains the response data along with audit information.
type AutoGenerated struct {
	AuditData AuditData `json:"auditData"` // Metadata about the request (e.g., timestamps, environment).
	Hotels    *Hotels   `json:"hotels"`    // Hotels information that includes hotel details and room availability.
}

// AuditData holds metadata about the request, like timestamps and the server details.
type AuditData struct {
	ProcessTime string `json:"processTime"` // Time it took to process the request.
	Timestamp   string `json:"timestamp"`   // Timestamp of when the request was made.
	RequestHost string `json:"requestHost"` // The host that received the request.
	ServerID    string `json:"serverId"`    // ID of the server handling the request.
	Environment string `json:"environment"` // The environment where the request was processed (e.g., production).
	Release     string `json:"release"`     // The version/release of the system processing the request.
	Token       string `json:"token"`       // Authentication token used for the request.
	Internal    string `json:"internal"`    // Internal flags or information used for request processing.
}

// Hotels represents a collection of hotels that match the search criteria.
type Hotels struct {
	Hotels   []Hotel `json:"hotels"`   // List of hotels returned from the search.
	CheckIn  string  `json:"checkIn"`  // The check-in date for the hotels.
	Total    int     `json:"total"`    // Total number of hotels returned in the search.
	CheckOut string  `json:"checkOut"` // The check-out date for the hotels.
}

// Hotel represents the details of a specific hotel.
type Hotel struct {
	Code            int    `json:"code"`            // Hotel's unique identifier.
	Name            string `json:"name"`            // Name of the hotel.
	CategoryCode    string `json:"categoryCode"`    // Category code for the hotel (e.g., 5-star, 3-star).
	CategoryName    string `json:"categoryName"`    // Category name of the hotel (e.g., Luxury, Budget).
	DestinationCode string `json:"destinationCode"` // Code for the destination (e.g., city or region).
	DestinationName string `json:"destinationName"` // Name of the destination (e.g., Paris, New York).
	ZoneCode        int    `json:"zoneCode"`        // Code for the geographical zone.
	ZoneName        string `json:"zoneName"`        // Name of the zone (e.g., Downtown, Airport).
	Latitude        string `json:"latitude"`        // Latitude coordinate of the hotel.
	Longitude       string `json:"longitude"`       // Longitude coordinate of the hotel.
	Rooms           []Room `json:"rooms"`           // Rooms available in the hotel.
	MinRate         string `json:"minRate"`         // Minimum rate for the rooms.
	MaxRate         string `json:"maxRate"`         // Maximum rate for the rooms.
	Currency        string `json:"currency"`        // Currency used for pricing.
}

// Room represents a specific room in a hotel.
type Room struct {
	Code  string `json:"code"`  // Room's unique code.
	Name  string `json:"name"`  // Name or type of the room (e.g., Deluxe Room).
	Rates []Rate `json:"rates"` // List of rates available for the room.
}

// Rate represents a specific pricing option for a room.
type Rate struct {
	RateKey              string               `json:"rateKey"`                  // Unique key for the rate.
	RateClass            string               `json:"rateClass"`                // Class of the rate (e.g., standard, promotional).
	RateType             string               `json:"rateType"`                 // Type of the rate (e.g., non-refundable, refundable).
	Net                  string               `json:"net"`                      // Net price for the rate.
	Allotment            int                  `json:"allotment"`                // Number of rooms available at this rate.
	PaymentType          string               `json:"paymentType"`              // Payment type (e.g., prepaid, pay-at-hotel).
	Packaging            bool                 `json:"packaging"`                // Whether this rate includes packaging (e.g., flight + hotel).
	BoardCode            string               `json:"boardCode"`                // Code for the meal board (e.g., BB for Bed and Breakfast).
	BoardName            string               `json:"boardName"`                // Name of the meal board (e.g., Full Board, All Inclusive).
	CancellationPolicies []CancellationPolicy `json:"cancellationPolicies"`     // List of cancellation policies for the rate.
	Rooms                int                  `json:"rooms"`                    // Number of rooms available at this rate.
	Adults               int                  `json:"adults"`                   // Number of adults for this rate.
	Children             int                  `json:"children"`                 // Number of children for this rate.
	ChildrenAges         string               `json:"childrenages"`             // Ages of children for this rate.
	DailyRates           []DailyRate          `json:"dailyRates"`               // List of daily rates for the room.
	SellingRate          string               `json:"sellingRate,omitempty"`    // Selling price of the rate (optional).
	HotelMandatory       bool                 `json:"hotelMandatory,omitempty"` // Whether the rate is mandatory for the hotel (optional).
	Offers               []Offer              `json:"offers,omitempty"`         // List of offers available with the rate.
}

// CancellationPolicy represents the cancellation policy for a rate.
type CancellationPolicy struct {
	Amount string    `json:"amount"` // Amount charged for cancellation.
	From   time.Time `json:"from"`   // Time from which cancellation applies.
}

// DailyRate represents a daily rate for a room.
type DailyRate struct {
	Offset   int    `json:"offset"`   // Offset from check-in date.
	DailyNet string `json:"dailyNet"` // Net price for the specific day.
}

// Offer represents a special offer available for a rate.
type Offer struct {
	Code   string `json:"code"`   // Unique code for the offer.
	Name   string `json:"name"`   // Name of the offer.
	Amount string `json:"amount"` // Discount or amount of the offer.
}
