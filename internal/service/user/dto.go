package user

type Filter struct {
	Limit     *int
	Offset    *int
	FirstName *string
}

type Create struct {
	ID       string `json:"id" bun:"id"`
	Email    string `json:"email" bun:"email"`
	Name     string `json:"name" bun:"name"`
	Age      int32 `json:"age" bun:"age"`
	Password string `json:"password" bun:"password"`
	Role     string `json:"role" bun:"role"`
}

type Get struct {
	ID       string `json:"id" bun:"id"`
	Email    string `json:"email" bun:"email"`
	Name     string `json:"name" bun:"name"`
	Age      int32 `json:"age" bun:"age"`
	Password string `json:"password" bun:"password"`
	Role     string `json:"role" bun:"role"`
}