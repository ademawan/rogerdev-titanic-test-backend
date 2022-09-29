package person

// "gorm.io/gorm"

type PersonCreateResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// =================== Create User Request =======================
type CreatePersonRequestFormat struct {
	Name string `json:"name" form:"name" validate:"required,min=3,max=25"`
	Age  int    `json:"age" form:"age" validate:"required"`
}

// type CreatePersonsRequestFormat struct {
// 	Persons []CreatePersonRequestFormat `json:"persons"`
// }
