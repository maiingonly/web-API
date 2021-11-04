package entities

type TopUPInput struct {
	// validation JSON required or not
	//if you want learn more about tag go https://pkg.go.dev/github.com/go-playground/validator/v10@v10.9.0#section-readme
	Name     string `json:"Name" binding:"required"`
	WalletID string `json:"WalletID" binding:"required"`
	Amount   int    `json:"Amount" binding:"required,number,max=1000000,min=10000"`
}
