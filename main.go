package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

<<<<<<< HEAD
	"github.com/idoall/gocryptotrader/common"
	"github.com/idoall/gocryptotrader/communications"
	"github.com/idoall/gocryptotrader/config"
	"github.com/idoall/gocryptotrader/connchecker"
	"github.com/idoall/gocryptotrader/currency"
	"github.com/idoall/gocryptotrader/currency/coinmarketcap"
	exchange "github.com/idoall/gocryptotrader/exchanges"
	log "github.com/idoall/gocryptotrader/logger"
	"github.com/idoall/gocryptotrader/ntpclient"
	"github.com/idoall/gocryptotrader/portfolio"
=======
	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/core"
	"github.com/thrasher-corp/gocryptotrader/dispatch"
	"github.com/thrasher-corp/gocryptotrader/engine"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	log "github.com/thrasher-corp/gocryptotrader/logger"
	"github.com/thrasher-corp/gocryptotrader/signaler"
>>>>>>> upstrem/master
)

func main() {
	// Handle flags
	var settings engine.Settings
	versionFlag := flag.Bool("version", false, "retrieves current GoCryptoTrader version")

	// Core settings
	flag.StringVar(&settings.ConfigFile, "config", "", "config file to load")
	flag.StringVar(&settings.DataDir, "datadir", common.GetDefaultDataDir(runtime.GOOS), "default data directory for GoCryptoTrader files")
	flag.IntVar(&settings.GoMaxProcs, "gomaxprocs", runtime.GOMAXPROCS(-1), "sets the runtime GOMAXPROCS value")
	flag.BoolVar(&settings.EnableDryRun, "dryrun", false, "dry runs bot, doesn't save config file")
	flag.BoolVar(&settings.EnableAllExchanges, "enableallexchanges", false, "enables all exchanges")
	flag.BoolVar(&settings.EnableAllPairs, "enableallpairs", false, "enables all pairs for enabled exchanges")
	flag.BoolVar(&settings.EnablePortfolioManager, "portfoliomanager", true, "enables the portfolio manager")
	flag.BoolVar(&settings.EnableGRPC, "grpc", true, "enables the grpc server")
	flag.BoolVar(&settings.EnableGRPCProxy, "grpcproxy", false, "enables the grpc proxy server")
	flag.BoolVar(&settings.EnableWebsocketRPC, "websocketrpc", true, "enables the websocket RPC server")
	flag.BoolVar(&settings.EnableDeprecatedRPC, "deprecatedrpc", true, "enables the deprecated RPC server")
	flag.BoolVar(&settings.EnableCommsRelayer, "enablecommsrelayer", true, "enables available communications relayer")
	flag.BoolVar(&settings.Verbose, "verbose", false, "increases logging verbosity for GoCryptoTrader")
	flag.BoolVar(&settings.EnableExchangeSyncManager, "syncmanager", true, "enables to exchange sync manager")
	flag.BoolVar(&settings.EnableWebsocketRoutine, "websocketroutine", true, "enables the websocket routine for all loaded exchanges")
	flag.BoolVar(&settings.EnableCoinmarketcapAnalysis, "coinmarketcap", false, "overrides config and runs currency analysis")
	flag.BoolVar(&settings.EnableEventManager, "eventmanager", true, "enables the event manager")
	flag.BoolVar(&settings.EnableOrderManager, "ordermanager", true, "enables the order manager")
	flag.BoolVar(&settings.EnableDepositAddressManager, "depositaddressmanager", true, "enables the deposit address manager")
	flag.BoolVar(&settings.EnableConnectivityMonitor, "connectivitymonitor", true, "enables the connectivity monitor")
	flag.BoolVar(&settings.EnableDatabaseManager, "databasemanager", true, "enables database manager")
	flag.DurationVar(&settings.EventManagerDelay, "eventmanagerdelay", time.Duration(0), "sets the event managers sleep delay between event checking")
	flag.BoolVar(&settings.EnableNTPClient, "ntpclient", true, "enables the NTP client to check system clock drift")
	flag.BoolVar(&settings.EnableDispatcher, "dispatch", true, "enables the dispatch system")
	flag.IntVar(&settings.DispatchMaxWorkerAmount, "dispatchworkers", dispatch.DefaultMaxWorkers, "sets the dispatch package max worker generation limit")
	flag.IntVar(&settings.DispatchJobsLimit, "dispatchjobslimit", dispatch.DefaultJobsLimit, "sets the dispatch package max jobs limit")

	// Exchange syncer settings
	flag.BoolVar(&settings.EnableTickerSyncing, "tickersync", true, "enables ticker syncing for all enabled exchanges")
	flag.BoolVar(&settings.EnableOrderbookSyncing, "orderbooksync", true, "enables orderbook syncing for all enabled exchanges")
	flag.BoolVar(&settings.EnableTradeSyncing, "tradesync", false, "enables trade syncing for all enabled exchanges")
	flag.IntVar(&settings.SyncWorkers, "syncworkers", engine.DefaultSyncerWorkers, "the amount of workers (goroutines) to use for syncing exchange data")
	flag.BoolVar(&settings.SyncContinuously, "synccontinuously", true, "whether to sync exchange data continuously (ticker, orderbook and trade history info")
	flag.DurationVar(&settings.SyncTimeout, "synctimeout", engine.DefaultSyncerTimeout,
		"the amount of time before the syncer will switch from one protocol to the other (e.g. from REST to websocket)")

	// Forex provider settings
	flag.BoolVar(&settings.EnableCurrencyConverter, "currencyconverter", false, "overrides config and sets up foreign exchange Currency Converter")
	flag.BoolVar(&settings.EnableCurrencyLayer, "currencylayer", false, "overrides config and sets up foreign exchange Currency Layer")
	flag.BoolVar(&settings.EnableFixer, "fixer", false, "overrides config and sets up foreign exchange Fixer.io")
	flag.BoolVar(&settings.EnableOpenExchangeRates, "openexchangerates", false, "overrides config and sets up foreign exchange Open Exchange Rates")

	// Exchange tuning settings
	flag.BoolVar(&settings.EnableExchangeAutoPairUpdates, "exchangeautopairupdates", false, "enables automatic available currency pair updates for supported exchanges")
	flag.BoolVar(&settings.DisableExchangeAutoPairUpdates, "exchangedisableautopairupdates", false, "disables exchange auto pair updates")
	flag.BoolVar(&settings.EnableExchangeWebsocketSupport, "exchangewebsocketsupport", false, "enables Websocket support for exchanges")
	flag.BoolVar(&settings.EnableExchangeRESTSupport, "exchangerestsupport", true, "enables REST support for exchanges")
	flag.BoolVar(&settings.EnableExchangeVerbose, "exchangeverbose", false, "increases exchange logging verbosity")
	flag.BoolVar(&settings.ExchangePurgeCredentials, "exchangepurgecredentials", false, "purges the stored exchange API credentials")
	flag.BoolVar(&settings.EnableExchangeHTTPRateLimiter, "ratelimiter", true, "enables the rate limiter for HTTP requests")
	flag.IntVar(&settings.MaxHTTPRequestJobsLimit, "requestjobslimit", request.DefaultMaxRequestJobs, "sets the max amount of jobs the HTTP request package stores")
	flag.IntVar(&settings.RequestTimeoutRetryAttempts, "exchangehttptimeoutretryattempts", request.DefaultTimeoutRetryAttempts, "sets the amount of retry attempts after a HTTP request times out")
	flag.DurationVar(&settings.ExchangeHTTPTimeout, "exchangehttptimeout", time.Duration(0), "sets the exchangs HTTP timeout value for HTTP requests")
	flag.StringVar(&settings.ExchangeHTTPUserAgent, "exchangehttpuseragent", "", "sets the exchanges HTTP user agent")
	flag.StringVar(&settings.ExchangeHTTPProxy, "exchangehttpproxy", "", "sets the exchanges HTTP proxy server")
	flag.BoolVar(&settings.EnableExchangeHTTPDebugging, "exchangehttpdebugging", false, "sets the exchanges HTTP debugging")

	// Common tuning settings
	flag.DurationVar(&settings.GlobalHTTPTimeout, "globalhttptimeout", time.Duration(0), "sets common HTTP timeout value for HTTP requests")
	flag.StringVar(&settings.GlobalHTTPUserAgent, "globalhttpuseragent", "", "sets the common HTTP client's user agent")
	flag.StringVar(&settings.GlobalHTTPProxy, "globalhttpproxy", "", "sets the common HTTP client's proxy server")

	flag.Parse()

	if *versionFlag {
		fmt.Print(core.Version(true))
		os.Exit(0)
	}

	fmt.Println(core.Banner)
	fmt.Println(core.Version(false))

<<<<<<< HEAD
	err = common.CreateDir(bot.dataDir)
	if err != nil {
		log.Fatalf("Failed to open/create data directory: %s. Err: %s", bot.dataDir, err)
	}
	log.Debugf("Using data directory: %s.\n", bot.dataDir)

	err = bot.config.CheckLoggerConfig()
	if err != nil {
		log.Errorf("Failed to configure logger reason: %s", err)
	}

	err = log.SetupLogger()
	if err != nil {
		log.Errorf("Failed to setup logger reason: %s", err)
	}

	ActivateNTP()
	ActivateConnectivityMonitor()
	AdjustGoMaxProcs()

	log.Debugf("Bot '%s' started.\n", bot.config.Name)
	log.Debugf("Bot dry run mode: %v.\n", common.IsEnabled(bot.dryRun))

	log.Debugf("Available Exchanges: %d. Enabled Exchanges: %d.\n",
		len(bot.config.Exchanges),
		bot.config.CountEnabledExchanges())

	// 设置全局的 http client 超时时间
	common.HTTPClient = common.NewHTTPClientWithTimeout(bot.config.GlobalHTTPTimeout)
	log.Debugf("Global HTTP request timeout: %v.\n", common.HTTPClient.Timeout)

	SetupExchanges()
	if len(bot.exchanges) == 0 {
		log.Fatal("No exchanges were able to be loaded. Exiting")
	}

	log.Debugf("Starting communication mediums..")
	cfg := bot.config.GetCommunicationsConfig()
	bot.comms = communications.NewComm(&cfg)
	bot.comms.GetEnabledCommunicationMediums()

	var newFxSettings []currency.FXSettings
	for _, d := range bot.config.Currency.ForexProviders {
		newFxSettings = append(newFxSettings, currency.FXSettings(d))
	}

	err = currency.RunStorageUpdater(currency.BotOverrides{
		Coinmarketcap:       *Coinmarketcap,
		FxCurrencyConverter: *FxCurrencyConverter,
		FxCurrencyLayer:     *FxCurrencyLayer,
		FxFixer:             *FxFixer,
		FxOpenExchangeRates: *FxOpenExchangeRates,
	},
		&currency.MainConfiguration{
			ForexProviders:         newFxSettings,
			CryptocurrencyProvider: coinmarketcap.Settings(bot.config.Currency.CryptocurrencyProvider),
			Cryptocurrencies:       bot.config.Currency.Cryptocurrencies,
			FiatDisplayCurrency:    bot.config.Currency.FiatDisplayCurrency,
			CurrencyDelay:          bot.config.Currency.CurrencyFileUpdateDuration,
			FxRateDelay:            bot.config.Currency.ForeignExchangeUpdateDuration,
		},
		bot.dataDir,
		*verbosity)
	if err != nil {
		log.Fatalf("currency updater system failed to start %v", err)

	}

	bot.portfolio = &portfolio.Portfolio
	bot.portfolio.SeedPortfolio(bot.config.Portfolio)
	SeedExchangeAccountInfo(GetAllEnabledExchangeAccountInfo().Data)

	ActivateWebServer()

	go portfolio.StartPortfolioWatcher()

	go TickerUpdaterRoutine()
	go OrderbookUpdaterRoutine()
	go WebsocketRoutine(*verbosity)

	<-bot.shutdown
	Shutdown()
}

// ActivateWebServer Sets up a local web server
func ActivateWebServer() {
	if bot.config.Webserver.Enabled {
		listenAddr := bot.config.Webserver.ListenAddress
		log.Debugf(
			"HTTP Webserver support enabled. Listen URL: http://%s:%d/\n",
			common.ExtractHost(listenAddr), common.ExtractPort(listenAddr),
		)

		router := NewRouter()
		go func() {
			err := http.ListenAndServe(listenAddr, router)
			if err != nil {
				log.Fatal(err)
			}
		}()

		log.Debugln("HTTP Webserver started successfully.")
		log.Debugln("Starting websocket handler.")
		StartWebsocketHandler()
	} else {
		log.Debugln("HTTP RESTful Webserver support disabled.")
	}
}

// ActivateConnectivityMonitor Sets up internet connectivity monitor
func ActivateConnectivityMonitor() {
=======
>>>>>>> upstrem/master
	var err error
	settings.CheckParamInteraction = true
	engine.Bot, err = engine.NewFromSettings(&settings)
	if engine.Bot == nil || err != nil {
		log.Errorf(log.Global, "Unable to initialise bot engine. Error: %s\n", err)
		os.Exit(1)
	}

	engine.PrintSettings(&engine.Bot.Settings)
	if err = engine.Bot.Start(); err != nil {
		log.Errorf(log.Global, "Unable to start bot engine. Error: %s\n", err)
		os.Exit(1)
	}

	interrupt := signaler.WaitForInterrupt()
	log.Infof(log.Global, "Captured %v, shutdown requested.\n", interrupt)
	engine.Bot.Stop()
	log.Infoln(log.Global, "Exiting.")
}
