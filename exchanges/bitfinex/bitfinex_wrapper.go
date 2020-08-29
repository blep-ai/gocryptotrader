package bitfinex

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/currency"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/account"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/kline"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
	"github.com/thrasher-corp/gocryptotrader/exchanges/orderbook"
	"github.com/thrasher-corp/gocryptotrader/exchanges/protocol"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"github.com/thrasher-corp/gocryptotrader/exchanges/stream"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
	"github.com/thrasher-corp/gocryptotrader/log"
	"github.com/thrasher-corp/gocryptotrader/portfolio/withdraw"
)

// GetDefaultConfig returns a default exchange config
func (b *Bitfinex) GetDefaultConfig() (*config.ExchangeConfig, error) {
	b.SetDefaults()
	exchCfg := new(config.ExchangeConfig)
	exchCfg.Name = b.Name
	exchCfg.HTTPTimeout = exchange.DefaultHTTPTimeout
	exchCfg.BaseCurrencies = b.BaseCurrencies

	err := b.SetupDefaults(exchCfg)
	if err != nil {
		return nil, err
	}

	if b.Features.Supports.RESTCapabilities.AutoPairUpdates {
		err = b.UpdateTradablePairs(true)
		if err != nil {
			return nil, err
		}
	}

	return exchCfg, nil
}

// SetDefaults sets the basic defaults for bitfinex
func (b *Bitfinex) SetDefaults() {
	b.Name = "Bitfinex"
	b.Enabled = true
	b.Verbose = true
	b.WebsocketSubdChannels = make(map[int]WebsocketChanInfo)
	b.API.CredentialsValidator.RequiresKey = true
	b.API.CredentialsValidator.RequiresSecret = true

	fmt1 := currency.PairStore{
		RequestFormat: &currency.PairFormat{Uppercase: true},
		ConfigFormat:  &currency.PairFormat{Uppercase: true},
	}

	fmt2 := currency.PairStore{
		RequestFormat: &currency.PairFormat{Uppercase: true},
		ConfigFormat:  &currency.PairFormat{Uppercase: true},
	}

	perpfmt := currency.PairStore{
		RequestFormat: &currency.PairFormat{Uppercase: true},
		ConfigFormat:  &currency.PairFormat{Uppercase: true, Delimiter: ":"},
	}

	err := b.StoreAssetPairFormat(asset.Spot, fmt1)
	if err != nil {
		log.Errorln(log.ExchangeSys, err)
	}
	err = b.StoreAssetPairFormat(asset.Margin, fmt2)
	if err != nil {
		log.Errorln(log.ExchangeSys, err)
	}
	err = b.StoreAssetPairFormat(asset.MarginFunding, fmt1)
	if err != nil {
		log.Errorln(log.ExchangeSys, err)
	}
	err = b.StoreAssetPairFormat(asset.PerpetualSwap, perpfmt)
	if err != nil {
		log.Errorln(log.ExchangeSys, err)
	}

	b.Features = exchange.Features{
		Supports: exchange.FeaturesSupported{
			REST:      true,
			Websocket: true,
			RESTCapabilities: protocol.Features{
				TickerBatching:      true,
				TickerFetching:      true,
				OrderbookFetching:   true,
				AutoPairUpdates:     true,
				AccountInfo:         true,
				CryptoDeposit:       true,
				CryptoWithdrawal:    true,
				FiatWithdraw:        true,
				GetOrder:            true,
				GetOrders:           true,
				CancelOrders:        true,
				CancelOrder:         true,
				SubmitOrder:         true,
				SubmitOrders:        true,
				DepositHistory:      true,
				WithdrawalHistory:   true,
				TradeFetching:       true,
				UserTradeHistory:    true,
				TradeFee:            true,
				FiatDepositFee:      true,
				FiatWithdrawalFee:   true,
				CryptoDepositFee:    true,
				CryptoWithdrawalFee: true,
			},
			WebsocketCapabilities: protocol.Features{
				AccountBalance:         true,
				CancelOrders:           true,
				CancelOrder:            true,
				SubmitOrder:            true,
				ModifyOrder:            true,
				TickerFetching:         true,
				KlineFetching:          true,
				TradeFetching:          true,
				OrderbookFetching:      true,
				AccountInfo:            true,
				Subscribe:              true,
				AuthenticatedEndpoints: true,
				MessageCorrelation:     true,
				DeadMansSwitch:         true,
				GetOrders:              true,
				GetOrder:               true,
			},
			WithdrawPermissions: exchange.AutoWithdrawCryptoWithAPIPermission |
				exchange.AutoWithdrawFiatWithAPIPermission,
			Kline: kline.ExchangeCapabilitiesSupported{
				DateRanges: true,
				Intervals:  true,
			},
		},
		Enabled: exchange.FeaturesEnabled{
			AutoPairUpdates: true,
			Kline: kline.ExchangeCapabilitiesEnabled{
				Intervals: map[string]bool{
					kline.OneMin.Word():     true,
					kline.ThreeMin.Word():   true,
					kline.FiveMin.Word():    true,
					kline.FifteenMin.Word(): true,
					kline.ThirtyMin.Word():  true,
					kline.OneHour.Word():    true,
					kline.TwoHour.Word():    true,
					kline.FourHour.Word():   true,
					kline.SixHour.Word():    true,
					kline.TwelveHour.Word(): true,
					kline.OneDay.Word():     true,
					kline.OneWeek.Word():    true,
					kline.TwoWeek.Word():    true,
				},
				ResultLimit: 10000,
			},
		},
	}

	b.Requester = request.New(b.Name,
		common.NewHTTPClientWithTimeout(exchange.DefaultHTTPTimeout),
		request.WithLimiter(SetRateLimit()))

	b.API.Endpoints.URLDefault = bitfinexAPIURLBase
	b.API.Endpoints.URL = b.API.Endpoints.URLDefault
	b.API.Endpoints.WebsocketURL = publicBitfinexWebsocketEndpoint
	b.Websocket = stream.New()
	b.WebsocketResponseMaxLimit = exchange.DefaultWebsocketResponseMaxLimit
	b.WebsocketResponseCheckTimeout = exchange.DefaultWebsocketResponseCheckTimeout
	b.WebsocketOrderbookBufferLimit = exchange.DefaultWebsocketOrderbookBufferLimit
}

// Setup takes in the supplied exchange configuration details and sets params
func (b *Bitfinex) Setup(exch *config.ExchangeConfig) error {
	if !exch.Enabled {
		b.SetEnabled(false)
		return nil
	}

	err := b.SetupDefaults(exch)
	if err != nil {
		return err
	}

	err = b.Websocket.Setup(&stream.WebsocketSetup{
		Enabled:                          exch.Features.Enabled.Websocket,
		Verbose:                          exch.Verbose,
		AuthenticatedWebsocketAPISupport: exch.API.AuthenticatedWebsocketSupport,
		WebsocketTimeout:                 exch.WebsocketTrafficTimeout,
		DefaultURL:                       publicBitfinexWebsocketEndpoint,
		ExchangeName:                     exch.Name,
		RunningURL:                       exch.API.Endpoints.WebsocketURL,
		Connector:                        b.WsConnect,
		Subscriber:                       b.Subscribe,
		UnSubscriber:                     b.Unsubscribe,
		GenerateSubscriptions:            b.GenerateDefaultSubscriptions,
		Features:                         &b.Features.Supports.WebsocketCapabilities,
		OrderbookBufferLimit:             exch.WebsocketOrderbookBufferLimit,
		UpdateEntriesByID:                true,
	})
	if err != nil {
		return err
	}

	err = b.Websocket.SetupNewConnection(stream.ConnectionSetup{
		ResponseCheckTimeout: exch.WebsocketResponseCheckTimeout,
		ResponseMaxLimit:     exch.WebsocketResponseMaxLimit,
		URL:                  publicBitfinexWebsocketEndpoint,
	})
	if err != nil {
		return err
	}

	return b.Websocket.SetupNewConnection(stream.ConnectionSetup{
		ResponseCheckTimeout: exch.WebsocketResponseCheckTimeout,
		ResponseMaxLimit:     exch.WebsocketResponseMaxLimit,
		URL:                  authenticatedBitfinexWebsocketEndpoint,
		Authenticated:        true,
	})
}

// Start starts the Bitfinex go routine
func (b *Bitfinex) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		b.Run()
		wg.Done()
	}()
}

// Run implements the Bitfinex wrapper
func (b *Bitfinex) Run() {
	if b.Verbose {
		log.Debugf(log.ExchangeSys,
			"%s Websocket: %s.",
			b.Name,
			common.IsEnabled(b.Websocket.IsEnabled()))
		b.PrintEnabledPairs()
	}

	if !b.GetEnabledFeatures().AutoPairUpdates {
		return
	}

	err := b.UpdateTradablePairs(false)
	if err != nil {
		log.Errorf(log.ExchangeSys,
			"%s failed to update tradable pairs. Err: %s",
			b.Name,
			err)
	}
}

// FetchTradablePairs returns a list of the exchanges tradable pairs
func (b *Bitfinex) FetchTradablePairs(a asset.Item) ([]string, error) {
	items, err := b.GetTickerBatch()
	if err != nil {
		return nil, err
	}

	var symbols []string
	switch a {
	case asset.Spot:
		for k := range items {
			if !strings.HasPrefix(k, "t") {
				continue
			}
			symbols = append(symbols, k[1:])
		}
	case asset.PerpetualSwap:
		for k := range items {
			if !strings.Contains(k, ":") {
				continue
			}
			symbols = append(symbols, k[1:])
		}
	case asset.MarginFunding:
		for k := range items {
			if !strings.HasPrefix(k, "f") {
				continue
			}
			symbols = append(symbols, k[1:])
		}
	default:
		return nil, errors.New("asset type not supported by this endpoint")
	}

	return symbols, nil
}

// UpdateTradablePairs updates the exchanges available pairs and stores
// them in the exchanges config
func (b *Bitfinex) UpdateTradablePairs(forceUpdate bool) error {
	assets := b.CurrencyPairs.GetAssetTypes()
	for i := range assets {
		pairs, err := b.FetchTradablePairs(assets[i])
		if err != nil {
			return err
		}

		p, err := currency.NewPairsFromStrings(pairs)
		if err != nil {
			return err
		}

		err = b.UpdatePairs(p, assets[i], false, forceUpdate)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateTicker updates and returns the ticker for a currency pair
func (b *Bitfinex) UpdateTicker(p currency.Pair, assetType asset.Item) (*ticker.Price, error) {
	enabledPairs, err := b.GetEnabledPairs(assetType)
	if err != nil {
		return nil, err
	}

	tickerNew, err := b.GetTickerBatch()
	if err != nil {
		return nil, err
	}

	for k, v := range tickerNew {
		pair, err := currency.NewPairFromString(k[1:]) // Remove prefix
		if err != nil {
			return nil, err
		}

		if !enabledPairs.Contains(pair, true) {
			continue
		}

		err = ticker.ProcessTicker(&ticker.Price{
			Last:         v.Last,
			High:         v.High,
			Low:          v.Low,
			Bid:          v.Bid,
			Ask:          v.Ask,
			Volume:       v.Volume,
			Pair:         pair,
			AssetType:    assetType,
			ExchangeName: b.Name})
		if err != nil {
			return nil, err
		}
	}
	return ticker.GetTicker(b.Name, p, assetType)
}

// FetchTicker returns the ticker for a currency pair
func (b *Bitfinex) FetchTicker(p currency.Pair, assetType asset.Item) (*ticker.Price, error) {
	b.appendOptionalDelimiter(&p)
	tick, err := ticker.GetTicker(b.Name, p, asset.Spot)
	if err != nil {
		return b.UpdateTicker(p, assetType)
	}
	return tick, nil
}

// FetchOrderbook returns the orderbook for a currency pair
func (b *Bitfinex) FetchOrderbook(p currency.Pair, assetType asset.Item) (*orderbook.Base, error) {
	b.appendOptionalDelimiter(&p)
	ob, err := orderbook.Get(b.Name, p, assetType)
	if err != nil {
		return b.UpdateOrderbook(p, assetType)
	}
	return ob, nil
}

// UpdateOrderbook updates and returns the orderbook for a currency pair
func (b *Bitfinex) UpdateOrderbook(p currency.Pair, assetType asset.Item) (*orderbook.Base, error) {
	b.appendOptionalDelimiter(&p)
	var prefix = "t"
	if assetType == asset.MarginFunding {
		prefix = "f"
	}

	orderbookNew, err := b.GetOrderbook(prefix+p.String(), "P0", 100)
	if err != nil {
		return nil, err
	}

	var o orderbook.Base
	for x := range orderbookNew.Asks {
		o.Asks = append(o.Asks, orderbook.Item{
			Price:  orderbookNew.Asks[x].Price,
			Amount: orderbookNew.Asks[x].Amount,
		})
	}

	for x := range orderbookNew.Bids {
		o.Bids = append(o.Bids, orderbook.Item{
			Price:  orderbookNew.Bids[x].Price,
			Amount: orderbookNew.Bids[x].Amount,
		})
	}

	o.Pair = p
	o.ExchangeName = b.Name
	o.AssetType = assetType

	err = o.Process()
	if err != nil {
		return nil, err
	}

	return orderbook.Get(b.Name, p, assetType)
}

// UpdateAccountInfo retrieves balances for all enabled currencies on the
// Bitfinex exchange
func (b *Bitfinex) UpdateAccountInfo() (account.Holdings, error) {
	var response account.Holdings
	response.Exchange = b.Name

	accountBalance, err := b.GetAccountBalance()
	if err != nil {
		return response, err
	}

	var Accounts = []account.SubAccount{
		{ID: "deposit"},
		{ID: "exchange"},
		{ID: "trading"},
		{ID: "margin"},
		{ID: "funding "},
	}

	for x := range accountBalance {
		for i := range Accounts {
			if Accounts[i].ID == accountBalance[x].Type {
				Accounts[i].Currencies = append(Accounts[i].Currencies,
					account.Balance{
						CurrencyName: currency.NewCode(accountBalance[x].Currency),
						TotalValue:   accountBalance[x].Amount,
						Hold:         accountBalance[x].Amount - accountBalance[x].Available,
					})
			}
		}
	}

	response.Accounts = Accounts
	err = account.Process(&response)
	if err != nil {
		return account.Holdings{}, err
	}

	return response, nil
}

// FetchAccountInfo retrieves balances for all enabled currencies
func (b *Bitfinex) FetchAccountInfo() (account.Holdings, error) {
	acc, err := account.GetHoldings(b.Name)
	if err != nil {
		return b.UpdateAccountInfo()
	}

	return acc, nil
}

// GetFundingHistory returns funding history, deposits and
// withdrawals
func (b *Bitfinex) GetFundingHistory() ([]exchange.FundHistory, error) {
	return nil, common.ErrFunctionNotSupported
}

// GetExchangeHistory returns historic trade data within the timeframe provided.
func (b *Bitfinex) GetExchangeHistory(p currency.Pair, assetType asset.Item, timestampStart, timestampEnd time.Time) ([]exchange.TradeHistory, error) {
	return nil, common.ErrNotYetImplemented
}

// SubmitOrder submits a new order
func (b *Bitfinex) SubmitOrder(o *order.Submit) (order.SubmitResponse, error) {
	var submitOrderResponse order.SubmitResponse
	err := o.Validate()
	if err != nil {
		return submitOrderResponse, err
	}

	fpair, err := b.FormatExchangeCurrency(o.Pair, o.AssetType)
	if err != nil {
		return submitOrderResponse, err
	}

	if b.Websocket.CanUseAuthenticatedWebsocketForWrapper() {
		submitOrderResponse.OrderID, err = b.WsNewOrder(&WsNewOrderRequest{
			CustomID: b.Websocket.AuthConn.GenerateMessageID(false),
			Type:     o.Type.String(),
			Symbol:   fpair.String(),
			Amount:   o.Amount,
			Price:    o.Price,
		})
		if err != nil {
			return submitOrderResponse, err
		}
	} else {
		var response Order
		isBuying := o.Side == order.Buy
		b.appendOptionalDelimiter(&fpair)
		orderType := o.Type.Lower()
		if o.AssetType == asset.Spot {
			orderType = "exchange " + orderType
		}
		response, err = b.NewOrder(fpair.String(),
			orderType,
			o.Amount,
			o.Price,
			isBuying,
			false)
		if err != nil {
			return submitOrderResponse, err
		}
		if response.ID > 0 {
			submitOrderResponse.OrderID = strconv.FormatInt(response.ID, 10)
		}
		if response.RemainingAmount == 0 {
			submitOrderResponse.FullyMatched = true
		}

		submitOrderResponse.IsOrderPlaced = true
	}
	return submitOrderResponse, err
}

// ModifyOrder will allow of changing orderbook placement and limit to
// market conversion
func (b *Bitfinex) ModifyOrder(action *order.Modify) (string, error) {
	orderIDInt, err := strconv.ParseInt(action.ID, 10, 64)
	if err != nil {
		return action.ID, err
	}
	if b.Websocket.CanUseAuthenticatedWebsocketForWrapper() {
		if action.Side == order.Sell && action.Amount > 0 {
			action.Amount = -1 * action.Amount
		}
		err = b.WsModifyOrder(&WsUpdateOrderRequest{
			OrderID: orderIDInt,
			Price:   action.Price,
			Amount:  action.Amount,
		})
		return action.ID, err
	}
	return "", common.ErrNotYetImplemented
}

// CancelOrder cancels an order by its corresponding ID number
func (b *Bitfinex) CancelOrder(order *order.Cancel) error {
	orderIDInt, err := strconv.ParseInt(order.ID, 10, 64)
	if err != nil {
		return err
	}
	if b.Websocket.CanUseAuthenticatedWebsocketForWrapper() {
		err = b.WsCancelOrder(orderIDInt)
	} else {
		_, err = b.CancelExistingOrder(orderIDInt)
	}
	return err
}

// CancelAllOrders cancels all orders associated with a currency pair
func (b *Bitfinex) CancelAllOrders(_ *order.Cancel) (order.CancelAllResponse, error) {
	var err error
	if b.Websocket.CanUseAuthenticatedWebsocketForWrapper() {
		err = b.WsCancelAllOrders()
	} else {
		_, err = b.CancelAllExistingOrders()
	}
	return order.CancelAllResponse{}, err
}

// GetOrderInfo returns information on a current open order
func (b *Bitfinex) GetOrderInfo(orderID string) (order.Detail, error) {
	var orderDetail order.Detail
	return orderDetail, common.ErrNotYetImplemented
}

// GetDepositAddress returns a deposit address for a specified currency
func (b *Bitfinex) GetDepositAddress(c currency.Code, accountID string) (string, error) {
	if accountID == "" {
		accountID = "deposit"
	}

	method, err := b.ConvertSymbolToDepositMethod(c)
	if err != nil {
		return "", err
	}

	resp, err := b.NewDeposit(method, accountID, 0)
	return resp.Address, err
}

// WithdrawCryptocurrencyFunds returns a withdrawal ID when a withdrawal is submitted
func (b *Bitfinex) WithdrawCryptocurrencyFunds(withdrawRequest *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	// Bitfinex has support for three types, exchange, margin and deposit
	// As this is for trading, I've made the wrapper default 'exchange'
	// TODO: Discover an automated way to make the decision for wallet type to withdraw from
	walletType := "exchange"
	resp, err := b.WithdrawCryptocurrency(walletType,
		withdrawRequest.Crypto.Address,
		withdrawRequest.Description,
		withdrawRequest.Amount,
		withdrawRequest.Currency)
	if err != nil {
		return nil, err
	}

	return &withdraw.ExchangeResponse{
		ID:     strconv.FormatInt(resp.WithdrawalID, 10),
		Status: resp.Status,
	}, err
}

// WithdrawFiatFunds returns a withdrawal ID when a withdrawal is submitted
// Returns comma delimited withdrawal IDs
func (b *Bitfinex) WithdrawFiatFunds(withdrawRequest *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	withdrawalType := "wire"
	// Bitfinex has support for three types, exchange, margin and deposit
	// As this is for trading, I've made the wrapper default 'exchange'
	// TODO: Discover an automated way to make the decision for wallet type to withdraw from
	walletType := "exchange"
	resp, err := b.WithdrawFIAT(withdrawalType, walletType, withdrawRequest)
	if err != nil {
		return nil, err
	}

	return &withdraw.ExchangeResponse{
		ID:     strconv.FormatInt(resp.WithdrawalID, 10),
		Status: resp.Status,
	}, err
}

// WithdrawFiatFundsToInternationalBank returns a withdrawal ID when a withdrawal is submitted
// Returns comma delimited withdrawal IDs
func (b *Bitfinex) WithdrawFiatFundsToInternationalBank(withdrawRequest *withdraw.Request) (*withdraw.ExchangeResponse, error) {
	v, err := b.WithdrawFiatFunds(withdrawRequest)
	if err != nil {
		return nil, err
	}
	return &withdraw.ExchangeResponse{
		ID:     v.ID,
		Status: v.Status,
	}, nil
}

// GetFeeByType returns an estimate of fee based on type of transaction
func (b *Bitfinex) GetFeeByType(feeBuilder *exchange.FeeBuilder) (float64, error) {
	if !b.AllowAuthenticatedRequest() && // Todo check connection status
		feeBuilder.FeeType == exchange.CryptocurrencyTradeFee {
		feeBuilder.FeeType = exchange.OfflineTradeFee
	}
	return b.GetFee(feeBuilder)
}

// GetActiveOrders retrieves any orders that are active/open
func (b *Bitfinex) GetActiveOrders(req *order.GetOrdersRequest) ([]order.Detail, error) {
	var orders []order.Detail
	resp, err := b.GetOpenOrders()
	if err != nil {
		return nil, err
	}

	for i := range resp {
		orderSide := order.Side(strings.ToUpper(resp[i].Side))
		timestamp, err := strconv.ParseFloat(resp[i].Timestamp, 64)
		if err != nil {
			log.Warnf(log.ExchangeSys,
				"Unable to convert timestamp '%s', leaving blank",
				resp[i].Timestamp)
		}

		pair, err := currency.NewPairFromString(resp[i].Symbol)
		if err != nil {
			return nil, err
		}

		orderDetail := order.Detail{
			Amount:          resp[i].OriginalAmount,
			Date:            time.Unix(int64(timestamp), 0),
			Exchange:        b.Name,
			ID:              strconv.FormatInt(resp[i].ID, 10),
			Side:            orderSide,
			Price:           resp[i].Price,
			RemainingAmount: resp[i].RemainingAmount,
			Pair:            pair,
			ExecutedAmount:  resp[i].ExecutedAmount,
		}

		switch {
		case resp[i].IsLive:
			orderDetail.Status = order.Active
		case resp[i].IsCancelled:
			orderDetail.Status = order.Cancelled
		case resp[i].IsHidden:
			orderDetail.Status = order.Hidden
		default:
			orderDetail.Status = order.UnknownStatus
		}

		// API docs discrepancy. Example contains prefixed "exchange "
		// Return type suggests “market” / “limit” / “stop” / “trailing-stop”
		orderType := strings.Replace(resp[i].Type, "exchange ", "", 1)
		if orderType == "trailing-stop" {
			orderDetail.Type = order.TrailingStop
		} else {
			orderDetail.Type = order.Type(strings.ToUpper(orderType))
		}

		orders = append(orders, orderDetail)
	}

	order.FilterOrdersBySide(&orders, req.Side)
	order.FilterOrdersByType(&orders, req.Type)
	order.FilterOrdersByTickRange(&orders, req.StartTicks, req.EndTicks)
	order.FilterOrdersByCurrencies(&orders, req.Pairs)
	return orders, nil
}

// GetOrderHistory retrieves account order information
// Can Limit response to specific order status
func (b *Bitfinex) GetOrderHistory(req *order.GetOrdersRequest) ([]order.Detail, error) {
	var orders []order.Detail
	resp, err := b.GetInactiveOrders()
	if err != nil {
		return nil, err
	}

	for i := range resp {
		orderSide := order.Side(strings.ToUpper(resp[i].Side))
		timestamp, err := strconv.ParseInt(resp[i].Timestamp, 10, 64)
		if err != nil {
			log.Warnf(log.ExchangeSys, "Unable to convert timestamp '%v', leaving blank", resp[i].Timestamp)
		}
		orderDate := time.Unix(timestamp, 0)

		pair, err := currency.NewPairFromString(resp[i].Symbol)
		if err != nil {
			return nil, err
		}

		orderDetail := order.Detail{
			Amount:          resp[i].OriginalAmount,
			Date:            orderDate,
			Exchange:        b.Name,
			ID:              strconv.FormatInt(resp[i].ID, 10),
			Side:            orderSide,
			Price:           resp[i].Price,
			RemainingAmount: resp[i].RemainingAmount,
			ExecutedAmount:  resp[i].ExecutedAmount,
			Pair:            pair,
		}

		switch {
		case resp[i].IsLive:
			orderDetail.Status = order.Active
		case resp[i].IsCancelled:
			orderDetail.Status = order.Cancelled
		case resp[i].IsHidden:
			orderDetail.Status = order.Hidden
		default:
			orderDetail.Status = order.UnknownStatus
		}

		// API docs discrepency. Example contains prefixed "exchange "
		// Return type suggests “market” / “limit” / “stop” / “trailing-stop”
		orderType := strings.Replace(resp[i].Type, "exchange ", "", 1)
		if orderType == "trailing-stop" {
			orderDetail.Type = order.TrailingStop
		} else {
			orderDetail.Type = order.Type(strings.ToUpper(orderType))
		}

		orders = append(orders, orderDetail)
	}

	order.FilterOrdersBySide(&orders, req.Side)
	order.FilterOrdersByType(&orders, req.Type)
	order.FilterOrdersByTickRange(&orders, req.StartTicks, req.EndTicks)
	for i := range req.Pairs {
		b.appendOptionalDelimiter(&req.Pairs[i])
	}
	order.FilterOrdersByCurrencies(&orders, req.Pairs)
	return orders, nil
}

// AuthenticateWebsocket sends an authentication message to the websocket
func (b *Bitfinex) AuthenticateWebsocket() error {
	return b.WsSendAuth()
}

// appendOptionalDelimiter ensures that a delimiter is present for long character currencies
func (b *Bitfinex) appendOptionalDelimiter(p *currency.Pair) {
	if len(p.Quote.String()) > 3 ||
		len(p.Base.String()) > 3 {
		p.Delimiter = ":"
	}
}

// ValidateCredentials validates current credentials used for wrapper
// functionality
func (b *Bitfinex) ValidateCredentials() error {
	_, err := b.UpdateAccountInfo()
	return b.CheckTransientError(err)
}

// FormatExchangeKlineInterval returns Interval to exchange formatted string
func (b *Bitfinex) FormatExchangeKlineInterval(in kline.Interval) string {
	switch in {
	case kline.OneDay:
		return "1D"
	case kline.OneWeek:
		return "7D"
	case kline.OneWeek * 2:
		return "14D"
	default:
		return in.Short()
	}
}

// GetHistoricCandles returns candles between a time period for a set time interval
func (b *Bitfinex) GetHistoricCandles(pair currency.Pair, a asset.Item, start, end time.Time, interval kline.Interval) (kline.Item, error) {
	if !b.KlineIntervalEnabled(interval) {
		return kline.Item{}, kline.ErrorKline{
			Interval: interval,
		}
	}

	if kline.TotalCandlesPerInterval(start, end, interval) > b.Features.Enabled.Kline.ResultLimit {
		return kline.Item{}, errors.New(kline.ErrRequestExceedsExchangeLimits)
	}

	cf, err := b.fixCasing(pair, a)
	if err != nil {
		return kline.Item{}, err
	}

	candles, err := b.GetCandles(cf, b.FormatExchangeKlineInterval(interval),
		start.Unix()*1000, end.Unix()*1000,
		b.Features.Enabled.Kline.ResultLimit, true)
	if err != nil {
		return kline.Item{}, err
	}
	ret := kline.Item{
		Exchange: b.Name,
		Pair:     pair,
		Asset:    a,
		Interval: interval,
	}

	for x := range candles {
		ret.Candles = append(ret.Candles, kline.Candle{
			Time:   candles[x].Timestamp,
			Open:   candles[x].Open,
			High:   candles[x].Close,
			Low:    candles[x].Low,
			Close:  candles[x].Close,
			Volume: candles[x].Volume,
		})
	}

	ret.SortCandlesByTimestamp(false)
	return ret, nil
}

// GetHistoricCandlesExtended returns candles between a time period for a set time interval
func (b *Bitfinex) GetHistoricCandlesExtended(pair currency.Pair, a asset.Item, start, end time.Time, interval kline.Interval) (kline.Item, error) {
	if !b.KlineIntervalEnabled(interval) {
		return kline.Item{}, kline.ErrorKline{
			Interval: interval,
		}
	}

	ret := kline.Item{
		Exchange: b.Name,
		Pair:     pair,
		Asset:    a,
		Interval: interval,
	}

	dates := kline.CalcDateRanges(start, end, interval, b.Features.Enabled.Kline.ResultLimit)
	cf, err := b.fixCasing(pair, a)
	if err != nil {
		return kline.Item{}, err
	}

	for x := range dates {
		candles, err := b.GetCandles(cf, b.FormatExchangeKlineInterval(interval),
			dates[x].Start.Unix()*1000, dates[x].End.Unix()*1000,
			b.Features.Enabled.Kline.ResultLimit, true)
		if err != nil {
			return kline.Item{}, err
		}

		for i := range candles {
			ret.Candles = append(ret.Candles, kline.Candle{
				Time:   candles[i].Timestamp,
				Open:   candles[i].Open,
				High:   candles[i].Close,
				Low:    candles[i].Low,
				Close:  candles[i].Close,
				Volume: candles[i].Volume,
			})
		}
	}

	ret.SortCandlesByTimestamp(false)
	return ret, nil
}

func (b *Bitfinex) fixCasing(in currency.Pair, a asset.Item) (string, error) {
	var checkString [2]byte
	if a == asset.Spot {
		checkString[0] = 't'
		checkString[1] = 'T'
	} else if a == asset.Margin {
		checkString[0] = 'f'
		checkString[1] = 'F'
	}

	fmt, err := b.FormatExchangeCurrency(in, a)
	if err != nil {
		return "", err
	}

	y := in.Base.String()
	if (y[0] != checkString[0] && y[0] != checkString[1]) ||
		(y[0] == checkString[1] && y[1] == checkString[1]) || in.Base == currency.TNB {
		return string(checkString[0]) + fmt.Upper().String(), nil
	}

	runes := []rune(fmt.Upper().String())
	runes[0] = unicode.ToLower(runes[0])
	return string(runes), nil
}
