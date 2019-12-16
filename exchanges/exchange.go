package exchange

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

<<<<<<< HEAD
	"github.com/idoall/gocryptotrader/common"
	"github.com/idoall/gocryptotrader/config"
	"github.com/idoall/gocryptotrader/currency"
	"github.com/idoall/gocryptotrader/exchanges/kline"
	"github.com/idoall/gocryptotrader/exchanges/orderbook"
	"github.com/idoall/gocryptotrader/exchanges/request"
	"github.com/idoall/gocryptotrader/exchanges/ticker"
	"github.com/idoall/gocryptotrader/exchanges/websocket/wshandler"
	log "github.com/idoall/gocryptotrader/logger"
=======
	"github.com/thrasher-corp/gocryptotrader/common/crypto"
	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/protocol"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	log "github.com/thrasher-corp/gocryptotrader/logger"
>>>>>>> upstrem/master
)

const (
	warningBase64DecryptSecretKeyFailed = "exchange %s unable to base64 decode secret key.. Disabling Authenticated API support" // nolint:gosec
	// WarningAuthenticatedRequestWithoutCredentialsSet error message for authenticated request without credentials set
	WarningAuthenticatedRequestWithoutCredentialsSet = "exchange %s authenticated HTTP request called but not supported due to unset/default API keys"
	// DefaultHTTPTimeout is the default HTTP/HTTPS Timeout for exchange requests
	DefaultHTTPTimeout = time.Second * 15
	// DefaultWebsocketResponseCheckTimeout websocket 响应时的默认延迟
	DefaultWebsocketResponseCheckTimeout = time.Millisecond * 30
	// DefaultWebsocketResponseMaxLimit websocket 响应的默认最大等待时间
	DefaultWebsocketResponseMaxLimit = time.Second * 7
	// DefaultWebsocketOrderbookBufferLimit is the maximum number of orderbook updates that get stored before being applied
	DefaultWebsocketOrderbookBufferLimit = 5
)

<<<<<<< HEAD
// FeeType custom type for calculating fees based on method
type FeeType uint8

// Const declarations for fee types
const (
	BankFee FeeType = iota
	InternationalBankDepositFee
	InternationalBankWithdrawalFee
	CryptocurrencyTradeFee
	CyptocurrencyDepositFee
	CryptocurrencyWithdrawalFee
	OfflineTradeFee
)

// InternationalBankTransactionType custom type for calculating fees based on fiat transaction types
type InternationalBankTransactionType uint8

// Const declarations for international transaction types
const (
	WireTransfer InternationalBankTransactionType = iota
	PerfectMoney
	Neteller
	AdvCash
	Payeer
	Skrill
	Simplex
	SEPA
	Swift
	RapidTransfer
	MisterTangoSEPA
	Qiwi
	VisaMastercard
	WebMoney
	Capitalist
	WesternUnion
	MoneyGram
	Contact
)

// SubmitOrderResponse is what is returned after submitting an order to an
// exchange
type SubmitOrderResponse struct {
	IsOrderPlaced bool
	OrderID       string
}

// FeeBuilder is the type which holds all parameters required to calculate a fee
// for an exchange
type FeeBuilder struct {
	IsMaker             bool
	PurchasePrice       float64
	Amount              float64
	FeeType             FeeType
	FiatCurrency        currency.Code
	BankTransactionType InternationalBankTransactionType
	Pair                currency.Pair
}

// OrderCancellation type required when requesting to cancel an order
type OrderCancellation struct {
	AccountID     string
	OrderID       string
	WalletAddress string
	Side          OrderSide
	CurrencyPair  currency.Pair
}

// WithdrawRequest used for wrapper crypto and FIAT withdraw methods
type WithdrawRequest struct {
	// General withdraw information
	Description     string
	OneTimePassword int64
	AccountID       string
	PIN             int64
	TradePassword   string
	Amount          float64
	Currency        currency.Code
	// Crypto related information
	Address    string
	AddressTag string
	FeeAmount  float64
	// FIAT related information
	BankAccountName   string
	BankAccountNumber float64
	BankName          string
	BankAddress       string
	BankCity          string
	BankCountry       string
	BankPostalCode    string
	SwiftCode         string
	IBAN              string
	BankCode          float64
	IsExpressWire     bool
	// Intermediary bank information
	RequiresIntermediaryBank      bool
	IntermediaryBankAccountNumber float64
	IntermediaryBankName          string
	IntermediaryBankAddress       string
	IntermediaryBankCity          string
	IntermediaryBankCountry       string
	IntermediaryBankPostalCode    string
	IntermediarySwiftCode         string
	IntermediaryBankCode          float64
	IntermediaryIBAN              string
	WireCurrency                  string
}

// Definitions for each type of withdrawal method for a given exchange
const (
	// No withdraw
	NoAPIWithdrawalMethods                  uint32 = 0
	NoAPIWithdrawalMethodsText              string = "NONE, WEBSITE ONLY"
	AutoWithdrawCrypto                      uint32 = (1 << 0)
	AutoWithdrawCryptoWithAPIPermission     uint32 = (1 << 1)
	AutoWithdrawCryptoWithSetup             uint32 = (1 << 2)
	AutoWithdrawCryptoText                  string = "AUTO WITHDRAW CRYPTO"
	AutoWithdrawCryptoWithAPIPermissionText string = "AUTO WITHDRAW CRYPTO WITH API PERMISSION"
	AutoWithdrawCryptoWithSetupText         string = "AUTO WITHDRAW CRYPTO WITH SETUP"
	WithdrawCryptoWith2FA                   uint32 = (1 << 3)
	WithdrawCryptoWithSMS                   uint32 = (1 << 4)
	WithdrawCryptoWithEmail                 uint32 = (1 << 5)
	WithdrawCryptoWithWebsiteApproval       uint32 = (1 << 6)
	WithdrawCryptoWithAPIPermission         uint32 = (1 << 7)
	WithdrawCryptoWith2FAText               string = "WITHDRAW CRYPTO WITH 2FA"
	WithdrawCryptoWithSMSText               string = "WITHDRAW CRYPTO WITH SMS"
	WithdrawCryptoWithEmailText             string = "WITHDRAW CRYPTO WITH EMAIL"
	WithdrawCryptoWithWebsiteApprovalText   string = "WITHDRAW CRYPTO WITH WEBSITE APPROVAL"
	WithdrawCryptoWithAPIPermissionText     string = "WITHDRAW CRYPTO WITH API PERMISSION"
	AutoWithdrawFiat                        uint32 = (1 << 8)
	AutoWithdrawFiatWithAPIPermission       uint32 = (1 << 9)
	AutoWithdrawFiatWithSetup               uint32 = (1 << 10)
	AutoWithdrawFiatText                    string = "AUTO WITHDRAW FIAT"
	AutoWithdrawFiatWithAPIPermissionText   string = "AUTO WITHDRAW FIAT WITH API PERMISSION"
	AutoWithdrawFiatWithSetupText           string = "AUTO WITHDRAW FIAT WITH SETUP"
	WithdrawFiatWith2FA                     uint32 = (1 << 11)
	WithdrawFiatWithSMS                     uint32 = (1 << 12)
	WithdrawFiatWithEmail                   uint32 = (1 << 13)
	WithdrawFiatWithWebsiteApproval         uint32 = (1 << 14)
	WithdrawFiatWithAPIPermission           uint32 = (1 << 15)
	WithdrawFiatWith2FAText                 string = "WITHDRAW FIAT WITH 2FA"
	WithdrawFiatWithSMSText                 string = "WITHDRAW FIAT WITH SMS"
	WithdrawFiatWithEmailText               string = "WITHDRAW FIAT WITH EMAIL"
	WithdrawFiatWithWebsiteApprovalText     string = "WITHDRAW FIAT WITH WEBSITE APPROVAL"
	WithdrawFiatWithAPIPermissionText       string = "WITHDRAW FIAT WITH API PERMISSION"
	WithdrawCryptoViaWebsiteOnly            uint32 = (1 << 16)
	WithdrawFiatViaWebsiteOnly              uint32 = (1 << 17)
	WithdrawCryptoViaWebsiteOnlyText        string = "WITHDRAW CRYPTO VIA WEBSITE ONLY"
	WithdrawFiatViaWebsiteOnlyText          string = "WITHDRAW FIAT VIA WEBSITE ONLY"
	NoFiatWithdrawals                       uint32 = (1 << 18)
	NoFiatWithdrawalsText                   string = "NO FIAT WITHDRAWAL"

	UnknownWithdrawalTypeText string = "UNKNOWN"

	RestAuthentication      uint8 = 0
	WebsocketAuthentication uint8 = 1
)

// AccountInfo is a Generic type to hold each exchange's holdings in
// all enabled currencies
type AccountInfo struct {
	Exchange string
	Accounts []Account
}

// Account defines a singular account type with asocciated currencies
type Account struct {
	ID         string
	Currencies []AccountCurrencyInfo
}

// AccountCurrencyInfo is a sub type to store currency name and value
type AccountCurrencyInfo struct {
	CurrencyName currency.Code
	TotalValue   float64
	Hold         float64
}

// TradeHistory holds exchange history data
type TradeHistory struct {
	Timestamp   time.Time
	TID         int64
	Price       float64
	Amount      float64
	Exchange    string
	Type        string
	Fee         float64
	Description string
}

// OrderDetail holds order detail data
type OrderDetail struct {
	Exchange        string
	AccountID       string
	ID              string
	CurrencyPair    currency.Pair
	OrderSide       OrderSide
	OrderType       OrderType
	OrderDate       time.Time
	Status          string
	Price           float64
	Amount          float64
	ExecutedAmount  float64
	RemainingAmount float64
	Fee             float64
	Trades          []TradeHistory
}

// FundHistory holds exchange funding history data
type FundHistory struct {
	ExchangeName      string
	Status            string
	TransferID        string
	Description       string
	Timestamp         time.Time
	Currency          string
	Amount            float64
	Fee               float64
	TransferType      string
	CryptoToAddress   string
	CryptoFromAddress string
	CryptoTxID        string
	BankTo            string
	BankFrom          string
}

// Base stores the individual exchange information
type Base struct {
	Name                                       string
	Enabled                                    bool
	Verbose                                    bool
	RESTPollingDelay                           time.Duration
	WebsocketResponseCheckTimeout              time.Duration
	WebsocketResponseMaxLimit                  time.Duration
	WebsocketOrderbookBufferLimit              int64
	AuthenticatedAPISupport                    bool
	AuthenticatedWebsocketAPISupport           bool
	APIWithdrawPermissions                     uint32
	APIAuthPEMKeySupport                       bool
	APISecret, APIKey, APIAuthPEMKey, ClientID string
	TakerFee, MakerFee, Fee                    float64
	BaseCurrencies                             currency.Currencies
	AvailablePairs                             currency.Pairs
	EnabledPairs                               currency.Pairs
	AssetTypes                                 []string
	PairsLastUpdated                           int64
	SupportsAutoPairUpdating                   bool
	SupportsRESTTickerBatching                 bool
	HTTPTimeout                                time.Duration
	HTTPUserAgent                              string
	HTTPDebugging                              bool
	HTTPRecording                              bool
	WebsocketURL                               string
	APIUrl                                     string
	APIUrlDefault                              string
	APIUrlSecondary                            string
	APIUrlSecondaryDefault                     string
	RequestCurrencyPairFormat                  config.CurrencyPairFormatConfig
	ConfigCurrencyPairFormat                   config.CurrencyPairFormatConfig
	Websocket                                  *wshandler.Websocket
	*request.Requester
}

// IBotExchange enforces standard functions for all exchanges supported in
// GoCryptoTrader
type IBotExchange interface {
	Setup(exch *config.ExchangeConfig)
	SetAPIKeys(apiKey, apiSecret, clientID string, b64Decode bool)
	Start(wg *sync.WaitGroup)
	SetDefaults()
	GetName() string
	IsEnabled() bool
	SetEnabled(bool)
	// GetTickerPrice 最后 24 时交易数据统计
	GetTickerPrice(currency currency.Pair, assetType string) (ticker.Price, error)
	UpdateTicker(currency currency.Pair, assetType string) (ticker.Price, error)
	GetOrderbookEx(currency currency.Pair, assetType string) (orderbook.Base, error)
	UpdateOrderbook(currency currency.Pair, assetType string) (orderbook.Base, error)
	GetEnabledCurrencies() currency.Pairs
	GetAvailableCurrencies() currency.Pairs
	GetAssetTypes() []string
	GetAccountInfo() (AccountInfo, error)
	GetAuthenticatedAPISupport(endpoint uint8) bool
	SetCurrencies(pairs []currency.Pair, enabledPairs bool) error
	GetExchangeHistory(p currency.Pair, assetType string) ([]TradeHistory, error)
	SupportsAutoPairUpdates() bool
	GetLastPairsUpdateTime() int64
	SupportsRESTTickerBatchUpdates() bool
	GetFeeByType(feeBuilder *FeeBuilder) (float64, error)
	GetWithdrawPermissions() uint32
	FormatWithdrawPermissions() string
	SupportsWithdrawPermissions(permissions uint32) bool
	GetFundingHistory() ([]FundHistory, error)
	SubmitOrder(p currency.Pair, side OrderSide, orderType OrderType, amount, price float64, clientID string) (SubmitOrderResponse, error)
	ModifyOrder(action *ModifyOrder) (string, error)
	CancelOrder(order *OrderCancellation) error
	CancelAllOrders(orders *OrderCancellation) (CancelAllOrdersResponse, error)
	GetOrderInfo(orderID string) (OrderDetail, error)
	GetDepositAddress(cryptocurrency currency.Code, accountID string) (string, error)
	GetOrderHistory(getOrdersRequest *GetOrdersRequest) ([]OrderDetail, error)
	GetActiveOrders(getOrdersRequest *GetOrdersRequest) ([]OrderDetail, error)
	WithdrawCryptocurrencyFunds(withdrawRequest *WithdrawRequest) (string, error)
	WithdrawFiatFunds(withdrawRequest *WithdrawRequest) (string, error)
	WithdrawFiatFundsToInternationalBank(withdrawRequest *WithdrawRequest) (string, error)
	GetWebsocket() (*wshandler.Websocket, error)
	SubscribeToWebsocketChannels(channels []wshandler.WebsocketChannelSubscription) error
	UnsubscribeToWebsocketChannels(channels []wshandler.WebsocketChannelSubscription) error
	AuthenticateWebsocket() error
	GetSubscriptions() ([]wshandler.WebsocketChannelSubscription, error)
	// GetKlines 自定义获取 K 线
	GetKlines(arg interface{}) ([]*kline.Kline, error)
}

// SupportsRESTTickerBatchUpdates returns whether or not the
// exhange supports REST batch ticker fetching
func (e *Base) SupportsRESTTickerBatchUpdates() bool {
	return e.SupportsRESTTickerBatching
}

// SetHTTPClientTimeout sets the timeout value for the exchanges
// HTTP Client
func (e *Base) SetHTTPClientTimeout(t time.Duration) {
=======
func (e *Base) checkAndInitRequester() {
>>>>>>> upstrem/master
	if e.Requester == nil {
		e.Requester = request.New(e.Name,
			request.NewRateLimit(time.Second, 0),
			request.NewRateLimit(time.Second, 0),
			new(http.Client))
	}
}

// SetHTTPClientTimeout sets the timeout value for the exchanges
// HTTP Client
func (e *Base) SetHTTPClientTimeout(t time.Duration) {
	e.checkAndInitRequester()
	e.Requester.HTTPClient.Timeout = t
}

// SetHTTPClient sets exchanges HTTP client
func (e *Base) SetHTTPClient(h *http.Client) {
	e.checkAndInitRequester()
	e.Requester.HTTPClient = h
}

// GetHTTPClient gets the exchanges HTTP client
func (e *Base) GetHTTPClient() *http.Client {
	e.checkAndInitRequester()
	return e.Requester.HTTPClient
}

// SetHTTPClientUserAgent sets the exchanges HTTP user agent
func (e *Base) SetHTTPClientUserAgent(ua string) {
	e.checkAndInitRequester()
	e.Requester.UserAgent = ua
	e.HTTPUserAgent = ua
}

// GetHTTPClientUserAgent gets the exchanges HTTP user agent
func (e *Base) GetHTTPClientUserAgent() string {
	return e.HTTPUserAgent
}

// SetClientProxyAddress 设置 URL 和 Websocket 的代理
func (e *Base) SetClientProxyAddress(addr string) error {
	if addr != "" {
		proxy, err := url.Parse(addr)
		if err != nil {
			return fmt.Errorf("exchange.go - setting proxy address error %s",
				err)
		}

		// No needs to check err here as the only err condition is an empty
		// string which is already checked above
		_ = e.Requester.SetProxy(proxy)

		if e.Websocket != nil {
			err = e.Websocket.SetProxyAddress(addr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

<<<<<<< HEAD
// SetAutoPairDefaults 设置是否支持自动更新交易对
func (e *Base) SetAutoPairDefaults() error {
	cfg := config.GetConfig()
	exch, err := cfg.GetExchangeConfig(e.Name)
	if err != nil {
		return err
	}

	update := false

	// 如果支持自动更新交易对
	if e.SupportsAutoPairUpdating {
		if !exch.SupportsAutoPairUpdates {
			exch.SupportsAutoPairUpdates = true
			exch.PairsLastUpdated = 0
			update = true
=======
// SetFeatureDefaults sets the exchanges default feature
// support set
func (e *Base) SetFeatureDefaults() {
	if e.Config.Features == nil {
		s := &config.FeaturesConfig{
			Supports: config.FeaturesSupportedConfig{
				Websocket: e.Features.Supports.Websocket,
				REST:      e.Features.Supports.REST,
				RESTCapabilities: protocol.Features{
					AutoPairUpdates: e.Features.Supports.RESTCapabilities.AutoPairUpdates,
				},
			},
		}

		if e.Config.SupportsAutoPairUpdates != nil {
			s.Supports.RESTCapabilities.AutoPairUpdates = *e.Config.SupportsAutoPairUpdates
			s.Enabled.AutoPairUpdates = *e.Config.SupportsAutoPairUpdates
		} else {
			s.Supports.RESTCapabilities.AutoPairUpdates = e.Features.Supports.RESTCapabilities.AutoPairUpdates
			s.Enabled.AutoPairUpdates = e.Features.Supports.RESTCapabilities.AutoPairUpdates
			if !s.Supports.RESTCapabilities.AutoPairUpdates {
				e.Config.CurrencyPairs.LastUpdated = time.Now().Unix()
				e.CurrencyPairs.LastUpdated = e.Config.CurrencyPairs.LastUpdated
			}
>>>>>>> upstrem/master
		}
		e.Config.Features = s
		e.Config.SupportsAutoPairUpdates = nil
	} else {
		if e.Features.Supports.RESTCapabilities.AutoPairUpdates != e.Config.Features.Supports.RESTCapabilities.AutoPairUpdates {
			e.Config.Features.Supports.RESTCapabilities.AutoPairUpdates = e.Features.Supports.RESTCapabilities.AutoPairUpdates

			if !e.Config.Features.Supports.RESTCapabilities.AutoPairUpdates {
				e.Config.CurrencyPairs.LastUpdated = time.Now().Unix()
			}
		}

		if e.Features.Supports.REST != e.Config.Features.Supports.REST {
			e.Config.Features.Supports.REST = e.Features.Supports.REST
		}

		if e.Features.Supports.RESTCapabilities.TickerBatching != e.Config.Features.Supports.RESTCapabilities.TickerBatching {
			e.Config.Features.Supports.RESTCapabilities.TickerBatching = e.Features.Supports.RESTCapabilities.TickerBatching
		}

		if e.Features.Supports.Websocket != e.Config.Features.Supports.Websocket {
			e.Config.Features.Supports.Websocket = e.Features.Supports.Websocket
		}

		e.Features.Enabled.AutoPairUpdates = e.Config.Features.Enabled.AutoPairUpdates
	}
}

<<<<<<< HEAD
	if update {
		// 更新交易所的配置
		return cfg.UpdateExchangeConfig(&exch)
=======
// SetAPICredentialDefaults sets the API Credential validator defaults
func (e *Base) SetAPICredentialDefaults() {
	// Exchange hardcoded settings take precedence and overwrite the config settings
	if e.Config.API.CredentialsValidator == nil {
		e.Config.API.CredentialsValidator = new(config.APICredentialsValidatorConfig)
	}
	if e.Config.API.CredentialsValidator.RequiresKey != e.API.CredentialsValidator.RequiresKey {
		e.Config.API.CredentialsValidator.RequiresKey = e.API.CredentialsValidator.RequiresKey
	}

	if e.Config.API.CredentialsValidator.RequiresSecret != e.API.CredentialsValidator.RequiresSecret {
		e.Config.API.CredentialsValidator.RequiresSecret = e.API.CredentialsValidator.RequiresSecret
	}

	if e.Config.API.CredentialsValidator.RequiresBase64DecodeSecret != e.API.CredentialsValidator.RequiresBase64DecodeSecret {
		e.Config.API.CredentialsValidator.RequiresBase64DecodeSecret = e.API.CredentialsValidator.RequiresBase64DecodeSecret
	}

	if e.Config.API.CredentialsValidator.RequiresClientID != e.API.CredentialsValidator.RequiresClientID {
		e.Config.API.CredentialsValidator.RequiresClientID = e.API.CredentialsValidator.RequiresClientID
	}

	if e.Config.API.CredentialsValidator.RequiresPEM != e.API.CredentialsValidator.RequiresPEM {
		e.Config.API.CredentialsValidator.RequiresPEM = e.API.CredentialsValidator.RequiresPEM
	}
}

// SetHTTPRateLimiter sets the exchanges default HTTP rate limiter and updates the exchange's config
// to default settings if it doesn't exist
func (e *Base) SetHTTPRateLimiter() {
	e.checkAndInitRequester()

	if e.RequiresRateLimiter() {
		if e.Config.HTTPRateLimiter == nil {
			e.Config.HTTPRateLimiter = new(config.HTTPRateLimitConfig)
			e.Config.HTTPRateLimiter.Authenticated.Duration = e.GetRateLimit(true).Duration
			e.Config.HTTPRateLimiter.Authenticated.Rate = e.GetRateLimit(true).Rate
			e.Config.HTTPRateLimiter.Unauthenticated.Duration = e.GetRateLimit(false).Duration
			e.Config.HTTPRateLimiter.Unauthenticated.Rate = e.GetRateLimit(false).Rate
		} else {
			e.SetRateLimit(true, e.Config.HTTPRateLimiter.Authenticated.Duration,
				e.Config.HTTPRateLimiter.Authenticated.Rate)
			e.SetRateLimit(false, e.Config.HTTPRateLimiter.Unauthenticated.Duration,
				e.Config.HTTPRateLimiter.Unauthenticated.Rate)
		}
>>>>>>> upstrem/master
	}
}

// SupportsRESTTickerBatchUpdates returns whether or not the
// exhange supports REST batch ticker fetching
func (e *Base) SupportsRESTTickerBatchUpdates() bool {
	return e.Features.Supports.RESTCapabilities.TickerBatching
}

// SupportsAutoPairUpdates returns whether or not the exchange supports
// auto currency pair updating
func (e *Base) SupportsAutoPairUpdates() bool {
	if e.Features.Supports.RESTCapabilities.AutoPairUpdates || e.Features.Supports.WebsocketCapabilities.AutoPairUpdates {
		return true
	}
	return false
}

// GetLastPairsUpdateTime returns the unix timestamp of when the exchanges
// currency pairs were last updated
func (e *Base) GetLastPairsUpdateTime() int64 {
	return e.CurrencyPairs.LastUpdated
}

<<<<<<< HEAD
// SetAssetTypes 检查exchange资产类型（是否支持SPOT，二进制或期货）如果不存在，则将其设置为默认设置
func (e *Base) SetAssetTypes() error {
	cfg := config.GetConfig()
	exch, err := cfg.GetExchangeConfig(e.Name)
	if err != nil {
		return err
=======
// SetAssetTypes checks the exchange asset types (whether it supports SPOT,
// Binary or Futures) and sets it to a default setting if it doesn't exist
func (e *Base) SetAssetTypes() {
	if e.Config.CurrencyPairs.AssetTypes.JoinToString(",") == "" {
		e.Config.CurrencyPairs.AssetTypes = e.CurrencyPairs.AssetTypes
	} else if e.Config.CurrencyPairs.AssetTypes.JoinToString(",") != e.CurrencyPairs.AssetTypes.JoinToString(",") {
		e.Config.CurrencyPairs.AssetTypes = e.CurrencyPairs.AssetTypes
>>>>>>> upstrem/master
	}
}

// GetAssetTypes returns the available asset types for an individual exchange
func (e *Base) GetAssetTypes() asset.Items {
	return e.CurrencyPairs.AssetTypes
}

// GetPairAssetType returns the associated asset type for the currency pair
func (e *Base) GetPairAssetType(c currency.Pair) (asset.Item, error) {
	for i := range e.GetAssetTypes() {
		if e.GetEnabledPairs(e.GetAssetTypes()[i]).Contains(c, true) {
			return e.GetAssetTypes()[i], nil
		}
	}
	return "", errors.New("asset type not associated with currency pair")
}

// GetClientBankAccounts returns banking details associated with
// a client for withdrawal purposes
func (e *Base) GetClientBankAccounts(exchangeName, withdrawalCurrency string) (config.BankAccount, error) {
	cfg := config.GetConfig()
	return cfg.GetClientBankAccounts(exchangeName, withdrawalCurrency)
}

// GetExchangeBankAccounts returns banking details associated with an
// exchange for funding purposes
func (e *Base) GetExchangeBankAccounts(exchangeName, depositCurrency string) (config.BankAccount, error) {
	cfg := config.GetConfig()
	return cfg.GetExchangeBankAccounts(exchangeName, depositCurrency)
}

// SetCurrencyPairFormat checks the exchange request and config currency pair
// formats and syncs it with the exchanges SetDefault settings
func (e *Base) SetCurrencyPairFormat() {
	if e.Config.CurrencyPairs == nil {
		e.Config.CurrencyPairs = new(currency.PairsManager)
	}

<<<<<<< HEAD
// SetSymbol 设置交易对
func (e *Base) SetSymbol(baseAsset, quoteAsset string) {
	// e.BaseAsset = baseAsset
	// e.QuoteAsset = quoteAsset
}

// GetSymbol 获取格式化的交易对
func (e *Base) GetSymbol() string {
	var symbol string
	// if e.RequestCurrencyPairFormat.Uppercase {
	// 	symbol = fmt.Sprintf("%s%s%s", strings.ToUpper(e.BaseAsset), e.RequestCurrencyPairFormat.Delimiter, strings.ToUpper(e.QuoteAsset))
	// } else {
	// 	symbol = fmt.Sprintf("%s%s%s", strings.ToLower(e.BaseAsset), e.RequestCurrencyPairFormat.Delimiter, strings.ToLower(e.QuoteAsset))
	// }
	return symbol
}

// SetCurrencyPairFormat checks the exchange request and config currency pair
// formats and sets it to a default setting if it doesn't exist
func (e *Base) SetCurrencyPairFormat() error {
	cfg := config.GetConfig()
	exch, err := cfg.GetExchangeConfig(e.Name)
	if err != nil {
		return err
=======
	e.Config.CurrencyPairs.UseGlobalFormat = e.CurrencyPairs.UseGlobalFormat
	if e.Config.CurrencyPairs.UseGlobalFormat {
		e.Config.CurrencyPairs.RequestFormat = e.CurrencyPairs.RequestFormat
		e.Config.CurrencyPairs.ConfigFormat = e.CurrencyPairs.ConfigFormat
		return
>>>>>>> upstrem/master
	}

	if e.Config.CurrencyPairs.ConfigFormat != nil {
		e.Config.CurrencyPairs.ConfigFormat = nil
	}
	if e.Config.CurrencyPairs.RequestFormat != nil {
		e.Config.CurrencyPairs.RequestFormat = nil
	}

	assetTypes := e.GetAssetTypes()
	for x := range assetTypes {
		if e.Config.CurrencyPairs.Get(assetTypes[x]) == nil {
			r := e.CurrencyPairs.Get(assetTypes[x])
			if r == nil {
				continue
			}
			e.Config.CurrencyPairs.Store(assetTypes[x], *e.CurrencyPairs.Get(assetTypes[x]))
		}
	}
}

// SetConfigPairs sets the exchanges currency pairs to the pairs set in the config
func (e *Base) SetConfigPairs() {
	assetTypes := e.GetAssetTypes()
	for x := range assetTypes {
		cfgPS := e.Config.CurrencyPairs.Get(assetTypes[x])
		if cfgPS == nil {
			continue
		}
		if e.Config.CurrencyPairs.UseGlobalFormat {
			e.CurrencyPairs.StorePairs(assetTypes[x], cfgPS.Available, false)
			e.CurrencyPairs.StorePairs(assetTypes[x], cfgPS.Enabled, true)
			continue
		}
		exchPS := e.CurrencyPairs.Get(assetTypes[x])
		cfgPS.ConfigFormat = exchPS.ConfigFormat
		cfgPS.RequestFormat = exchPS.RequestFormat
		e.CurrencyPairs.StorePairs(assetTypes[x], cfgPS.Available, false)
		e.CurrencyPairs.StorePairs(assetTypes[x], cfgPS.Enabled, true)
	}
}

// GetAuthenticatedAPISupport returns whether the exchange supports
// authenticated API requests
func (e *Base) GetAuthenticatedAPISupport(endpoint uint8) bool {
	switch endpoint {
	case RestAuthentication:
		return e.API.AuthenticatedSupport
	case WebsocketAuthentication:
		return e.API.AuthenticatedWebsocketSupport
	}
	return false
}

// GetName is a method that returns the name of the exchange base
func (e *Base) GetName() string {
	return e.Name
}

// GetEnabledFeatures returns the exchanges enabled features
func (e *Base) GetEnabledFeatures() FeaturesEnabled {
	return e.Features.Enabled
}

// GetSupportedFeatures returns the exchanges supported features
func (e *Base) GetSupportedFeatures() FeaturesSupported {
	return e.Features.Supports
}

// GetPairFormat returns the pair format based on the exchange and
// asset type
func (e *Base) GetPairFormat(assetType asset.Item, requestFormat bool) currency.PairFormat {
	if e.CurrencyPairs.UseGlobalFormat {
		if requestFormat {
			return *e.CurrencyPairs.RequestFormat
		}
		return *e.CurrencyPairs.ConfigFormat
	}

	if requestFormat {
		return *e.CurrencyPairs.Get(assetType).RequestFormat
	}
	return *e.CurrencyPairs.Get(assetType).ConfigFormat
}

// GetEnabledPairs is a method that returns the enabled currency pairs of
// the exchange by asset type
func (e *Base) GetEnabledPairs(assetType asset.Item) currency.Pairs {
	format := e.GetPairFormat(assetType, false)
	pairs := e.CurrencyPairs.GetPairs(assetType, true)
	return pairs.Format(format.Delimiter, format.Index, format.Uppercase)
}

// GetAvailablePairs is a method that returns the available currency pairs
// of the exchange by asset type
func (e *Base) GetAvailablePairs(assetType asset.Item) currency.Pairs {
	format := e.GetPairFormat(assetType, false)
	pairs := e.CurrencyPairs.GetPairs(assetType, false)
	return pairs.Format(format.Delimiter, format.Index, format.Uppercase)
}

// SupportsPair returns true or not whether a currency pair exists in the
// exchange available currencies or not
func (e *Base) SupportsPair(p currency.Pair, enabledPairs bool, assetType asset.Item) bool {
	if enabledPairs {
		return e.GetEnabledPairs(assetType).Contains(p, false)
	}
	return e.GetAvailablePairs(assetType).Contains(p, false)
}

// FormatExchangeCurrencies returns a string containing
// the exchanges formatted currency pairs
func (e *Base) FormatExchangeCurrencies(pairs []currency.Pair, assetType asset.Item) (string, error) {
	var currencyItems strings.Builder
	pairFmt := e.GetPairFormat(assetType, true)

	for x := range pairs {
		currencyItems.WriteString(e.FormatExchangeCurrency(pairs[x], assetType).String())
		if x == len(pairs)-1 {
			continue
		}
		currencyItems.WriteString(pairFmt.Separator)
	}

	if currencyItems.Len() == 0 {
		return "", errors.New("returned empty string")
	}
	return currencyItems.String(), nil
}

// FormatExchangeCurrency is a method that formats and returns a currency pair
// based on the user currency display preferences
func (e *Base) FormatExchangeCurrency(p currency.Pair, assetType asset.Item) currency.Pair {
	pairFmt := e.GetPairFormat(assetType, true)
	return p.Format(pairFmt.Delimiter, pairFmt.Uppercase)
}

// SetEnabled is a method that sets if the exchange is enabled
func (e *Base) SetEnabled(enabled bool) {
	e.Enabled = enabled
}

// IsEnabled is a method that returns if the current exchange is enabled
func (e *Base) IsEnabled() bool {
	return e.Enabled
}

// SetAPISecretKeys is a method that sets the current API keys for the exchange
func (e *Base) SetAPISecretKeys(APIKey, APISecret string) {
	e.APIKey = APIKey
	e.APISecret = APISecret
}

// SetAPIKeys is a method that sets the current API keys for the exchange
func (e *Base) SetAPIKeys(apiKey, apiSecret, clientID string) {
	e.API.Credentials.Key = apiKey
	e.API.Credentials.ClientID = clientID

	if e.API.CredentialsValidator.RequiresBase64DecodeSecret {
		result, err := crypto.Base64Decode(apiSecret)
		if err != nil {
			e.API.AuthenticatedSupport = false
			e.API.AuthenticatedWebsocketSupport = false
			log.Warnf(log.ExchangeSys, warningBase64DecryptSecretKeyFailed, e.Name)
			return
		}
		e.API.Credentials.Secret = string(result)
	} else {
		e.API.Credentials.Secret = apiSecret
	}
}

// SetupDefaults sets the exchange settings based on the supplied config
func (e *Base) SetupDefaults(exch *config.ExchangeConfig) error {
	e.Enabled = true
	e.LoadedByConfig = true
	e.Config = exch
	e.Verbose = exch.Verbose

	e.API.AuthenticatedSupport = exch.API.AuthenticatedSupport
	e.API.AuthenticatedWebsocketSupport = exch.API.AuthenticatedWebsocketSupport
	if e.API.AuthenticatedSupport || e.API.AuthenticatedWebsocketSupport {
		e.SetAPIKeys(exch.API.Credentials.Key, exch.API.Credentials.Secret, exch.API.Credentials.ClientID)
	}

	if exch.HTTPTimeout <= time.Duration(0) {
		exch.HTTPTimeout = DefaultHTTPTimeout
	} else {
		e.SetHTTPClientTimeout(exch.HTTPTimeout)
	}

	if exch.CurrencyPairs == nil {
		exch.CurrencyPairs = new(currency.PairsManager)
	}

	e.HTTPDebugging = exch.HTTPDebugging
	e.SetHTTPClientUserAgent(exch.HTTPUserAgent)
	e.SetHTTPRateLimiter()
	e.SetAssetTypes()
	e.SetCurrencyPairFormat()
	e.SetConfigPairs()
	e.SetFeatureDefaults()
	e.SetAPIURL()
	e.SetAPICredentialDefaults()
	e.SetClientProxyAddress(exch.ProxyAddress)
	e.SetHTTPRateLimiter()

	e.BaseCurrencies = exch.BaseCurrencies

	if e.Features.Supports.Websocket {
		return e.Websocket.Initialise()
	}
	return nil
}

// AllowAuthenticatedRequest checks to see if the required fields have been set before sending an authenticated
// API request
func (e *Base) AllowAuthenticatedRequest() bool {
	// Skip auth check
	if e.SkipAuthCheck {
		return true
	}

	// Individual package usage, allow request if API credentials are valid a
	// and without needing to set AuthenticatedSupport to true
	if !e.LoadedByConfig && !e.ValidateAPICredentials() {
		return false
	}

	// Bot usage, AuthenticatedSupport can be disabled by user if desired, so don't
	// allow authenticated requests.
	if (!e.API.AuthenticatedSupport && !e.API.AuthenticatedWebsocketSupport) && e.LoadedByConfig {
		return false
	}

	// Check to see if the user has enabled AuthenticatedSupport, but has invalid
	// API credentials set and loaded by config
	if (e.API.AuthenticatedSupport || e.API.AuthenticatedWebsocketSupport) && e.LoadedByConfig && !e.ValidateAPICredentials() {
		return false
	}

	return true
}

// ValidateAPICredentials validates the exchanges API credentials
func (e *Base) ValidateAPICredentials() bool {
	if e.API.CredentialsValidator.RequiresKey {
		if e.API.Credentials.Key == "" ||
			e.API.Credentials.Key == config.DefaultAPIKey {
			log.Warnf(log.ExchangeSys,
				"exchange %s requires API key but default/empty one set",
				e.Name)
			return false
		}
	}

	if e.API.CredentialsValidator.RequiresSecret {
		if e.API.Credentials.Secret == "" ||
			e.API.Credentials.Secret == config.DefaultAPISecret {
			log.Warnf(log.ExchangeSys,
				"exchange %s requires API secret but default/empty one set",
				e.Name)
			return false
		}
	}

	if e.API.CredentialsValidator.RequiresPEM {
		if e.API.Credentials.PEMKey == "" ||
			strings.Contains(e.API.Credentials.PEMKey, "JUSTADUMMY") {
			log.Warnf(log.ExchangeSys,
				"exchange %s requires API PEM key but default/empty one set",
				e.Name)
			return false
		}
	}

	if e.API.CredentialsValidator.RequiresClientID {
		if e.API.Credentials.ClientID == "" ||
			e.API.Credentials.ClientID == config.DefaultAPIClientID {
			log.Warnf(log.ExchangeSys,
				"exchange %s requires API ClientID but default/empty one set",
				e.Name)
			return false
		}
	}

	if e.API.CredentialsValidator.RequiresBase64DecodeSecret && !e.LoadedByConfig {
		_, err := crypto.Base64Decode(e.API.Credentials.Secret)
		if err != nil {
			log.Warnf(log.ExchangeSys,
				"exchange %s API secret base64 decode failed: %s",
				e.Name, err)
			return false
		}
	}
	return true
}

// SetPairs sets the exchange currency pairs for either enabledPairs or
// availablePairs
func (e *Base) SetPairs(pairs currency.Pairs, assetType asset.Item, enabled bool) error {
	if len(pairs) == 0 {
		return fmt.Errorf("%s SetPairs error - pairs is empty", e.Name)
	}

	pairFmt := e.GetPairFormat(assetType, false)
	var newPairs currency.Pairs
	for x := range pairs {
		newPairs = append(newPairs, pairs[x].Format(pairFmt.Delimiter,
			pairFmt.Uppercase))
	}

	e.CurrencyPairs.StorePairs(assetType, newPairs, enabled)
	e.Config.CurrencyPairs.StorePairs(assetType, newPairs, enabled)
	return nil
}

// UpdatePairs updates the exchange currency pairs for either enabledPairs or
// availablePairs
func (e *Base) UpdatePairs(exchangeProducts currency.Pairs, assetType asset.Item, enabled, force bool) error {
	if len(exchangeProducts) == 0 {
		return fmt.Errorf("%s UpdatePairs error - exchangeProducts is empty", e.Name)
	}

	exchangeProducts = exchangeProducts.Upper()
	var products currency.Pairs
	for x := range exchangeProducts {
		if exchangeProducts[x].String() == "" {
			continue
		}
		products = append(products, exchangeProducts[x])
	}

	var newPairs, removedPairs currency.Pairs
	var updateType string
	targetPairs := e.CurrencyPairs.GetPairs(assetType, enabled)

	if enabled {
		newPairs, removedPairs = targetPairs.FindDifferences(products)
		updateType = "enabled"
	} else {
		newPairs, removedPairs = targetPairs.FindDifferences(products)
		updateType = "available"
	}

	if force || len(newPairs) > 0 || len(removedPairs) > 0 {
		if force {
			log.Debugf(log.ExchangeSys,
				"%s forced update of %s [%v] pairs.", e.Name, updateType,
				strings.ToUpper(assetType.String()))
		} else {
			if len(newPairs) > 0 {
				log.Debugf(log.ExchangeSys,
					"%s Updating pairs [%v] - New: %s.\n", e.Name,
					strings.ToUpper(assetType.String()), newPairs)
			}
			if len(removedPairs) > 0 {
				log.Debugf(log.ExchangeSys,
					"%s Updating pairs [%v] - Removed: %s.\n", e.Name,
					strings.ToUpper(assetType.String()), removedPairs)
			}
		}
<<<<<<< HEAD

		if enabled {
			exch.EnabledPairs = products
			e.EnabledPairs = products
		} else {
			exch.AvailablePairs = products
			e.AvailablePairs = products
		}
		return cfg.UpdateExchangeConfig(&exch)
=======
		e.Config.CurrencyPairs.StorePairs(assetType, products, enabled)
		e.CurrencyPairs.StorePairs(assetType, products, enabled)
>>>>>>> upstrem/master
	}
	return nil
}

<<<<<<< HEAD
// ModifyOrder is a an order modifyer
type ModifyOrder struct {
	OrderID string
	OrderType
	OrderSide
	Price           float64
	Amount          float64
	LimitPriceUpper float64
	LimitPriceLower float64
	CurrencyPair    currency.Pair

	ImmediateOrCancel bool
	HiddenOrder       bool
	FillOrKill        bool
	PostOnly          bool
}

// ModifyOrderResponse is an order modifying return type
type ModifyOrderResponse struct {
	OrderID string
}

// Format holds exchange formatting
type Format struct {
	ExchangeName string
	OrderType    map[string]string
	OrderSide    map[string]string
}

// CancelAllOrdersResponse returns the status from attempting to cancel all orders on an exchagne
type CancelAllOrdersResponse struct {
	OrderStatus map[string]string
}

// Formatting contain a range of exchanges formatting
type Formatting []Format

// OrderType enforces a standard for Ordertypes across the code base
type OrderType string

// OrderType ...types
const (
	AnyOrderType               OrderType = "ANY"
	LimitOrderType             OrderType = "LIMIT"
	MarketOrderType            OrderType = "MARKET"
	ImmediateOrCancelOrderType OrderType = "IMMEDIATE_OR_CANCEL"
	StopOrderType              OrderType = "STOP"
	TrailingStopOrderType      OrderType = "TRAILINGSTOP"
	UnknownOrderType           OrderType = "UNKNOWN"
)

// ToString changes the ordertype to the exchange standard and returns a string
func (o OrderType) ToString() string {
	return fmt.Sprintf("%v", o)
}

// OrderSide enforces a standard for OrderSides across the code base
type OrderSide string

// OrderSide types
const (
	AnyOrderSide  OrderSide = "ANY"
	BuyOrderSide  OrderSide = "BUY"
	SellOrderSide OrderSide = "SELL"
	BidOrderSide  OrderSide = "BID"
	AskOrderSide  OrderSide = "ASK"
)

// ToString changes the ordertype to the exchange standard and returns a string
func (o OrderSide) ToString() string {
	return fmt.Sprintf("%v", o)
}

=======
>>>>>>> upstrem/master
// SetAPIURL sets configuration API URL for an exchange
func (e *Base) SetAPIURL() error {
	if e.Config.API.Endpoints.URL == "" || e.Config.API.Endpoints.URLSecondary == "" {
		return fmt.Errorf("exchange %s: SetAPIURL error. URL vals are empty", e.Name)
	}

	checkInsecureEndpoint := func(endpoint string) {
		if !strings.Contains(endpoint, "https") {
			return
		}
		log.Warnf(log.ExchangeSys,
			"%s is using HTTP instead of HTTPS [%s] for API functionality, an"+
				" attacker could eavesdrop on this connection. Use at your"+
				" own risk.",
			e.Name, endpoint)
	}
<<<<<<< HEAD
	if ec.APIURLSecondary != config.APIURLNonDefaultMessage {
		e.APIUrlSecondary = ec.APIURLSecondary
=======

	if e.Config.API.Endpoints.URL != config.APIURLNonDefaultMessage {
		e.API.Endpoints.URL = e.Config.API.Endpoints.URL
		checkInsecureEndpoint(e.API.Endpoints.URL)
	}
	if e.Config.API.Endpoints.URLSecondary != config.APIURLNonDefaultMessage {
		e.API.Endpoints.URLSecondary = e.Config.API.Endpoints.URLSecondary
		checkInsecureEndpoint(e.API.Endpoints.URLSecondary)
>>>>>>> upstrem/master
	}
	return nil
}

// GetAPIURL returns the set API URL
func (e *Base) GetAPIURL() string {
	return e.API.Endpoints.URL
}

// GetSecondaryAPIURL returns the set Secondary API URL
func (e *Base) GetSecondaryAPIURL() string {
	return e.API.Endpoints.URLSecondary
}

// GetAPIURLDefault returns exchange default URL
func (e *Base) GetAPIURLDefault() string {
	return e.API.Endpoints.URLDefault
}

// GetAPIURLSecondaryDefault returns exchange default secondary URL
func (e *Base) GetAPIURLSecondaryDefault() string {
	return e.API.Endpoints.URLSecondaryDefault
}

// SupportsWebsocket returns whether or not the exchange supports
// websocket
func (e *Base) SupportsWebsocket() bool {
	return e.Features.Supports.Websocket
}

// SupportsREST returns whether or not the exchange supports
// REST
func (e *Base) SupportsREST() bool {
	return e.Features.Supports.REST
}

// IsWebsocketEnabled returns whether or not the exchange has its
// websocket client enabled
func (e *Base) IsWebsocketEnabled() bool {
	if e.Websocket != nil {
		return e.Websocket.IsEnabled()
	}
	return false
}

// GetWithdrawPermissions passes through the exchange's withdraw permissions
func (e *Base) GetWithdrawPermissions() uint32 {
	return e.Features.Supports.WithdrawPermissions
}

// SupportsWithdrawPermissions compares the supplied permissions with the exchange's to verify they're supported
func (e *Base) SupportsWithdrawPermissions(permissions uint32) bool {
	exchangePermissions := e.GetWithdrawPermissions()
	return permissions&exchangePermissions == permissions
}

// FormatWithdrawPermissions will return each of the exchange's compatible withdrawal methods in readable form
func (e *Base) FormatWithdrawPermissions() string {
	var services []string
	for i := 0; i < 32; i++ {
		var check uint32 = 1 << uint32(i)
		if e.GetWithdrawPermissions()&check != 0 {
			switch check {
			case AutoWithdrawCrypto:
				services = append(services, AutoWithdrawCryptoText)
			case AutoWithdrawCryptoWithAPIPermission:
				services = append(services, AutoWithdrawCryptoWithAPIPermissionText)
			case AutoWithdrawCryptoWithSetup:
				services = append(services, AutoWithdrawCryptoWithSetupText)
			case WithdrawCryptoWith2FA:
				services = append(services, WithdrawCryptoWith2FAText)
			case WithdrawCryptoWithSMS:
				services = append(services, WithdrawCryptoWithSMSText)
			case WithdrawCryptoWithEmail:
				services = append(services, WithdrawCryptoWithEmailText)
			case WithdrawCryptoWithWebsiteApproval:
				services = append(services, WithdrawCryptoWithWebsiteApprovalText)
			case WithdrawCryptoWithAPIPermission:
				services = append(services, WithdrawCryptoWithAPIPermissionText)
			case AutoWithdrawFiat:
				services = append(services, AutoWithdrawFiatText)
			case AutoWithdrawFiatWithAPIPermission:
				services = append(services, AutoWithdrawFiatWithAPIPermissionText)
			case AutoWithdrawFiatWithSetup:
				services = append(services, AutoWithdrawFiatWithSetupText)
			case WithdrawFiatWith2FA:
				services = append(services, WithdrawFiatWith2FAText)
			case WithdrawFiatWithSMS:
				services = append(services, WithdrawFiatWithSMSText)
			case WithdrawFiatWithEmail:
				services = append(services, WithdrawFiatWithEmailText)
			case WithdrawFiatWithWebsiteApproval:
				services = append(services, WithdrawFiatWithWebsiteApprovalText)
			case WithdrawFiatWithAPIPermission:
				services = append(services, WithdrawFiatWithAPIPermissionText)
			case WithdrawCryptoViaWebsiteOnly:
				services = append(services, WithdrawCryptoViaWebsiteOnlyText)
			case WithdrawFiatViaWebsiteOnly:
				services = append(services, WithdrawFiatViaWebsiteOnlyText)
			case NoFiatWithdrawals:
				services = append(services, NoFiatWithdrawalsText)
			default:
				services = append(services, fmt.Sprintf("%s[1<<%v]", UnknownWithdrawalTypeText, i))
			}
		}
	}
	if len(services) > 0 {
		return strings.Join(services, " & ")
	}

	return NoAPIWithdrawalMethodsText
}

// SupportsAsset whether or not the supplied asset is supported
// by the exchange
func (e *Base) SupportsAsset(a asset.Item) bool {
	return e.CurrencyPairs.AssetTypes.Contains(a)
}

// PrintEnabledPairs prints the exchanges enabled asset pairs
func (e *Base) PrintEnabledPairs() {
	for k, v := range e.CurrencyPairs.Pairs {
		log.Infof(log.ExchangeSys, "%s Asset type %v:\n\t Enabled pairs: %v",
			e.Name, strings.ToUpper(k.String()), v.Enabled)
	}
}

// GetBase returns the exchange base
func (e *Base) GetBase() *Base { return e }
