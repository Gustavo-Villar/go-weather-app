package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gustavo-villar/go-weather-tracker/service-b/internal/api"
	"github.com/gustavo-villar/go-weather-tracker/service-b/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func init() {
	os.Setenv("APP_NAME", "service-b")
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	// Set up context with signal handling
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Initialize telemetry
	shutdown, err := telemetry.SetupProvider(ctx, os.Getenv("APP_NAME"))
	if err != nil {
		return
	}

	// Set up HTTP multiplexer and tracer
	mux := http.NewServeMux()
	tracer := otel.Tracer("weather")

	// Register handler with OpenTelemetry tracing
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		// Extract tracing context from the incoming request
		carrier := propagation.HeaderCarrier(r.Header)
		ctx := r.Context()
		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

		ctx, span := tracer.Start(ctx, "HandleGetWeather")
		defer span.End()

		// Call the existing handler function
		api.HandleGetWeather(w, r.WithContext(ctx))
	})

	// Set up HTTP server configuration
	srv := &http.Server{
		Addr:         ":8001",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	defer func() {
		err = errors.Join(err, shutdown(context.Background()))
	}()

	// Channel to capture server errors
	srvErr := make(chan error, 1)
	go func() {
		log.Printf("Server starting on port %s...", srv.Addr)
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for an interrupt signal or server error
	select {
	case err = <-srvErr:
		// Server start error
		return
	case <-ctx.Done():
		// Wait for first CTRL+C
		stop()
	}

	// Graceful shutdown
	err = srv.Shutdown(context.Background())
	return
}
