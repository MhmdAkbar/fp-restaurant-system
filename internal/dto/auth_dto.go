package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin waiter cashier super_admin"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UpdateProfileRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateProfileResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TableRequest struct {
	Id        uint   `json:"id"`
	TableCode string `json:"table_code"`
	QrCode    string `json:"qr_code"`
}
type TableUpdateRequest struct {
	Id        uint   `json:"id"`
	TableCode string `json:"table_code"`
	QrCode    string `json:"qr_code"`
}
type TableStatusUpdate struct {
	Id uint `json:"id"`
}

type TableResponse struct {
	Id        uint   `json:"id"`
	TableCode string `json:"table_number"`
	QrCode    string `json:"qr_code"`
	Status    bool   `json:"status"`
}
