package models

type Location struct {
	Latitude  float64 `json:"lat" xml:"lat" form:"lat" query:"lat"`
	Longitude float64 `json:"lng" xml:"lng" form:"lng" query:"lng"`
}
