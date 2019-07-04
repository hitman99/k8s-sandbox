package cmd

import (
	"context"
	"fmt"
	pgloadsvr "github.com/hitman99/k8s-sandbox/gen/http/pgload/server"
	"github.com/hitman99/k8s-sandbox/gen/pgload"
	"github.com/hitman99/k8s-sandbox/services"
	"github.com/spf13/cobra"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

func serve() {
	var (
		hostF     = "0.0.0.0"
		domainF   = ""
		httpPortF = "8080"
		secureF   = false
		dbgF      = false
	)
	logger := log.New(os.Stderr, "[k8s-sandbox] ", log.Ltime)
	// Initialize the services.
	pgloadSvc := services.NewPgload(logger)
	pgloadEndpoints := pgload.NewEndpoints(pgloadSvc)

	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	addr := "http://" + hostF
	u, err := url.Parse(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
		os.Exit(1)
	}
	if secureF {
		u.Scheme = "https"
	}
	if domainF != "" {
		u.Host = domainF
	}
	if httpPortF != "" {
		h := strings.Split(u.Host, ":")[0]
		u.Host = h + ":" + httpPortF
	} else if u.Port() == "" {
		u.Host += ":80"
	}
	handleHTTPServer(ctx, u, pgloadEndpoints, &wg, errc, logger, dbgF)
	logger.Printf("exiting (%v)", <-errc)
	// Send cancellation signal to the goroutines.
	cancel()
	wg.Wait()
	logger.Println("exited")
}

func handleHTTPServer(ctx context.Context, u *url.URL, pgloadEndpoints *pgload.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {
	adapter := middleware.NewLogger(logger)

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		pgloadServer *pgloadsvr.Server
	)
	{
		eh := errorHandler(logger)
		pgloadServer = pgloadsvr.New(pgloadEndpoints, mux, dec, enc, eh)
	}
	// Configure the mux.
	pgloadsvr.Mount(mux, pgloadServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		if debug {
			handler = httpmdlwr.Debug(mux, os.Stdout)(handler)
		}
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range pgloadServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

var httpServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "serves an application by starting an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}
