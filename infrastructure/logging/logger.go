package logging

import (
	"net"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once         sync.Once
	logstashConn net.Conn
	logger       *zap.Logger
)

/*
 * Initialize the connection to logstash
 * Note: Requires the .env file to be loaded before calling this function
 */
func InitLogstashConnection() {
	host := os.Getenv("LOGSTASH_HOST")
	port := os.Getenv("LOGSTASH_PORT")
	address := host + ":" + port

	conn, err := net.Dial("tcp", address)

	if err != nil {
		panic("Error connecting to logstash")
	}

	logstashConn = conn
}

/*
 * Instantiate a new logger that will be used for logstash
 * communications.
 *
 */
func InitLogger() {
	if logstashConn == nil {
		InitLogstashConnection()
	}

	if logstashConn == nil {
		logger.Sugar().Fatalw("Logstash connection is nil",
			"level", "critical",
			"timestamp", zap.Time("timestamp", time.Now()),
			"event", "logstash_connection")
	}

	once.Do(func() {
		writer := zapcore.AddSync(logstashConn)
		config := zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.CallerKey = "caller"
		encoder := zapcore.NewJSONEncoder(config.EncoderConfig)
		core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)

		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	})
}

/**
 * Return the logger instance connected to logstash
 */
func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		InitLogger()
	}

	return logger.Sugar()
}

/**
 * Close the connection to logstash
 */
func CloseLogstashConnection() {
	if logstashConn != nil {
		logstashConn.Close()
	}
}

/**
 * Sync the logger
 */
func SyncLogger() {
	if logger != nil {
		logger.Sync()
	}
}
