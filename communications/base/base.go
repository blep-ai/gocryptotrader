package base

import (
	"time"
<<<<<<< HEAD

	"github.com/idoall/gocryptotrader/common"
	"github.com/idoall/gocryptotrader/exchanges/ticker"
=======
>>>>>>> upstrem/master
)

// global vars contain staged update data that will be sent to the communication
// mediums
var (
<<<<<<< HEAD
	// map[exchangeName]
	TickerStaged    map[string]map[string]map[string]ticker.Price
	OrderbookStaged map[string]map[string]map[string]Orderbook
	PortfolioStaged Portfolio
	SettingsStaged  Settings
	ServiceStarted  time.Time
	m               sync.Mutex
)

// Orderbook 订单信息
// medium
type Orderbook struct {
	CurrencyPair string
	AssetType    string
	TotalAsks    float64
	TotalBids    float64
	LastUpdated  string
}

// Ticker 交易数据信息
// medium
type Ticker struct {
	CurrencyPair string
	LastUpdated  string
}

// Portfolio holds the minimal portfolio details to be sent to a communication
// medium
type Portfolio struct {
	ProfitLoss string
}

// Settings holds the minimal setting details to be sent to a communication
// medium
type Settings struct {
	EnabledExchanges      string
	EnabledCommunications string
}

=======
	ServiceStarted time.Time
)

>>>>>>> upstrem/master
// Base enforces standard variables across communication packages
type Base struct {
	Name      string
	Enabled   bool
	Verbose   bool
	Connected bool
}

// Event is a generalise event type
type Event struct {
	Type    string
	Message string
}

// CommsStatus stores the status of a comms relayer
type CommsStatus struct {
	Enabled   bool `json:"enabled"`
	Connected bool `json:"connected"`
}

// IsEnabled returns if the comms package has been enabled in the configuration
func (b *Base) IsEnabled() bool {
	return b.Enabled
}

// IsConnected returns if the package is connected to a server and/or ready to
// send
func (b *Base) IsConnected() bool {
	return b.Connected
}

// GetName returns a package name
func (b *Base) GetName() string {
	return b.Name
}

// GetStatus returns status data
func (b *Base) GetStatus() string {
	return `
	GoCryptoTrader Service: Online
	Service Started: ` + ServiceStarted.String()
}
