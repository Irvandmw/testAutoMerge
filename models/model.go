package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	UserType int    `json:"user_type"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Products `json:"data"`
}

type ProductsResponses struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Products `json:"data"`
}

type Transactions struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type TransactionResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    Transactions `json:"data"`
}

type TransactionsResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []Transactions `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Address struct {
	ID     int    `json:"id"`
	Street string `json:"street"`
	UserID int    `json:"user_id"`
}

type UserAdresses struct {
	User    User      `json:"user"`
	Address []Address `json:"address"`
}

type UserAdress struct {
	User    User    `json:"user"`
	Address Address `json:"address"`
}

type UserAdressesResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    UserAdresses `json:"data"`
}

type DetailTransactions struct {
	User         User           `json:"user"`
	Product      Products       `json:"product"`
	Transactions []Transactions `json:"transactions"`
}

type DetailTransaction struct {
	ID       int      `json:"id"`
	User     User     `json:"user"`
	Product  Products `json:"product"`
	Quantity int      `json:"quantity"`
}

type DetailTransactionsResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    []DetailTransaction `json:"data"`
}
