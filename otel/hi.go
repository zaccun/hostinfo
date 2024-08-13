package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"strconv"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
)

func main() {
	//Setup Otel log provider
	logExporter, err := stdoutlog.New()
	if err != nil {
		os.Exit(1)
	}
	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	defer loggerProvider.Shutdown(context.Background())
	global.SetLoggerProvider(loggerProvider)
	logger := otelslog.NewLogger("hostinfo")

	//Set listening port
	port := 9898
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			logger.Info(err.Error())
			os.Exit(1)
		}
	}
	//Find default output IP and hostname
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.Info(err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	hostname, err := os.Hostname()
	if err != nil {
		logger.Info(err.Error())
		os.Exit(1)
	}
	//Register handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ClientAddress := r.Header.Get("X-Real-Ip")
		if ClientAddress == "" {
			ClientAddress = r.Header.Get("X-Forwarded-For")
		}
		if ClientAddress == "" {
			ClientAddress = r.RemoteAddr
		}
		logger.Info("Request received", "ClientAddress", ClientAddress)
		w.Write([]byte(hostname + " " + localAddr.IP.String() + "\n"))
	})
	//Serve
	logger.Info("Serving", "Hostname", hostname, "IP", localAddr.IP.String(), "Port", strconv.Itoa(port))
	logger.Info(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil).Error())
}

