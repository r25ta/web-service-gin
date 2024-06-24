package model

// album represents data about a record album.
/*Struct tags such as json:"artist" specify what a field’s name should be when the struct’s
contents are serialized into JSON.
Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.
*/
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
