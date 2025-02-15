package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"john": {
		AuthToken: "145ABC",
		Username:  "john",
	},
	"shayla": {
		AuthToken: "156ABC",
		Username:  "shayla",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    200,
		Username: "alex",
	},
	"john": {
		Coins:    100,
		Username: "john",
	},
	"shayla": {
		Coins:    900,
		Username: "shayla",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
