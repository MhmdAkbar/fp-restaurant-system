package dto

type AddMenuRequest struct {
	Name        string  `json:"name"`                                                       // nama menu
	Category    string  `json:"category" validate:"required,oneof=makanan minuman dessert"` // kategori menu (string dulu, nanti di-convert)
	Price       float64 `json:"price"`                                                      // harga
	Description string  `json:"description"`                                                // deskripsi
	ImageURL    string  `json:"image_url"`                                                  // url gambar
	IsAvailable bool    `json:"is_available"`                                               // ketersediaan
}

type AddMenuResponse struct { // menjelaskan output data menu
	ID          uint    `json:"id"`           // id menu yang baru dibuat
	Name        string  `json:"name"`         // nama menu
	Category    string  `json:"category"`     // kategori menu (string)
	Price       float64 `json:"price"`        // harga
	Description string  `json:"description"`  // deskripsi
	ImageURL    string  `json:"image_url"`    // url gambar
	IsAvailable bool    `json:"is_available"` // status
}


type UpdateMenuRequest struct {
	Name        string  `json:"name"`
	Category    string  `json:"category" validate:"required,oneof=makanan minuman dessert"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	IsAvailable bool    `json:"is_available"`
}
