package models

// GeoLocation struct represents the geographical coordinates of a location, typically
// used to store the latitude and longitude of the hotel.
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`  // Latitude of the location.
	Longitude float64 `json:"longitude"` // Longitude of the location.
}

// RoomDetails struct provides detailed information about a specific room, including
// the board name (meal plan), price, and other characteristics of the room.
type RoomDetails struct {
	BoardName  string  `json:"board_name"` // The type of board/meal plan offered with the room (e.g., "BB", "HB").
	RoomName   string  `json:"room_name"`  // The name or description of the room.
	Price      float64 `json:"price"`      // Price of the room per night.
	Refundable bool    `json:"refundable"` // Indicates whether the room is refundable or not.
	Adult      int     `json:"adult"`      // Number of adults the room accommodates.
	Child      int     `json:"child"`      // Number of children the room accommodates.
	ChildAge   []int   `json:"child_age"`  // Age range for children, represented as a list of ages.
}

// HotelListing struct represents a complete listing of a hotel, including basic details
// such as the hotel name, rating, geographical location, and available room details.
type HotelListing struct {
	HotelCode    string      `json:"hotel_code"`    // Unique code identifying the hotel.
	Name         string      `json:"name"`          // Name of the hotel.
	Rating       float64     `json:"rating"`        // Rating of the hotel (e.g., 4.5 stars).
	Thumbnail    string      `json:"thumbnail"`     // URL or path to the hotel's thumbnail image.
	PropertyType string      `json:"property_type"` // Type of the property (e.g., "resort", "apartment").
	GeoLocation  GeoLocation `json:"geo_location"`  // Geographical location of the hotel (latitude & longitude).
	Address      string      `json:"address"`       // Address of the hotel.

	GuestRating float64     `json:"guest_rating"` // Guest rating for the hotel, typically a score from customer reviews.
	RoomDetails []RoomGroup `json:"room_details"` // List of room groups available in the hotel.
}

// RoomGroup struct represents a group of rooms (e.g., rooms of the same type or category)
// within a hotel listing, along with the price, currency, and detailed information about each room.
type RoomGroup struct {
	Price       float64       `json:"price"`       // Price of the room group (could be the average or range of prices).
	Currency    string        `json:"currency"`    // Currency in which the price is quoted (e.g., USD, EUR).
	RoomDetails []RoomDetails `json:"roomdetails"` // List of room details within this room group.
}

type HotelResponse struct {
	TrackerID    string         `json:"tracker_id"`
	HotelListing []HotelListing `json:"hotel_listing"`
	ErrorMessage string         `json:"error_message,omitempty"`
}
