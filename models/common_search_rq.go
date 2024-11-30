package models

// ExecutorRoom represents the room details for an executor request,
// including the number of adults, children, and the children's ages.
type ExecutorRoom struct {
	Room     int   `json:"Room"`     // The room number or identifier.
	Adult    int   `json:"adult"`    // The number of adults in the room.
	Child    int   `json:"child"`    // The number of children in the room.
	ChildAge []int `json:"childAge"` // A list of children's ages, if applicable.
}

// HotelExecutorRequest represents the structure of a request to execute a hotel search or action,
// which includes room details, dates, and other relevant data.
type HotelExecutorRequest struct {
	Country       string         `json:"Country"`       // The country where the hotel is located.
	CheckinDate   string         `json:"checkinDate"`   // The check-in date for the booking.
	CheckoutDate  string         `json:"checkoutDate"`  // The checkout date for the booking.
	ExecutorRooms []ExecutorRoom `json:"ExecutorRooms"` // A list of rooms for the request, including the number of people and their details.
	HotelCode     string         `json:"hotelcode"`     // The hotel code or identifier.
	Nationality   string         `json:"nationality"`   // The nationality of the requester.
	MealOptions   string         `json:"MealOptions"`   // The meal options chosen for the rooms.
}

// Test represents a basic test structure, with a single value field used for testing or validation purposes.
type Test struct {
	Value string `json:"value"` // The value associated with this test request.
}
