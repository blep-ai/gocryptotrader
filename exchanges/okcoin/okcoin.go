package okcoin

import (
<<<<<<< HEAD
	"time"

	"github.com/idoall/gocryptotrader/common"
	exchange "github.com/idoall/gocryptotrader/exchanges"
	"github.com/idoall/gocryptotrader/exchanges/okgroup"
	"github.com/idoall/gocryptotrader/exchanges/request"
	"github.com/idoall/gocryptotrader/exchanges/ticker"
	"github.com/idoall/gocryptotrader/exchanges/websocket/wshandler"
=======
	"github.com/thrasher-corp/gocryptotrader/exchanges/okgroup"
>>>>>>> upstrem/master
)

const (
	okCoinAuthRate     = 600
	okCoinUnauthRate   = 600
	okCoinAPIPath      = "api/"
	okCoinAPIURL       = "https://www.okcoin.com/" + okCoinAPIPath
	okCoinAPIVersion   = "/v3/"
	okCoinExchangeName = "OKCOIN International"
	okCoinWebsocketURL = "wss://real.okcoin.com:10442/ws/v3"
)

// OKCoin bases all methods off okgroup implementation
type OKCoin struct {
	okgroup.OKGroup
}
