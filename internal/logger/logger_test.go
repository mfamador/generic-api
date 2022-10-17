package logger_test

import (
	"bytes"
	"encoding/json"
	"genericsapi/internal/logger"
	"reflect"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestLogger(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger.SetPretty(false)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var (
	output          = new(bytes.Buffer)
	timeBeforeTests = time.Now().UnixNano() / int64(time.Millisecond)
)

func testLogging(logFunction func(string), logString, logLevel string) {
	logFunction(logString)

	// Ensure we're outputting a JSON message
	logMessageMap := make(map[string]interface{})
	err := json.Unmarshal(output.Bytes(), &logMessageMap)
	Expect(err).NotTo(HaveOccurred())

	// Assert field types
	Expect(reflect.TypeOf(logMessageMap["level"]).String()).To(Equal("string"))
	// JSON has no distinction between int and float. go defaults to float
	Expect(reflect.TypeOf(logMessageMap["time"]).String()).To(Equal("float64"))
	Expect(reflect.TypeOf(logMessageMap["msg"]).String()).To(Equal("string"))

	// Assert field value
	Expect(logMessageMap["level"]).To(Equal(logLevel))
	Expect(logMessageMap["msg"]).To(Equal(logString))

	// Check that the logging message timestamp is greater or equal than a sampled
	// timestamp taken before the tests begin.
	// Convert first to float64 then int64 to be consistent with TypeOf
	Expect(timeBeforeTests <= int64(logMessageMap["time"].(float64))).To(BeTrue())
}

var _ = Describe("Test logging", func() {
	log.Logger = log.Output(output)

	BeforeEach(func() {
		output.Reset()
	})

	Context("should output correct logs", func() {
		It("should output correct debug logs", func() {
			testLogging(log.Debug().Msg, "Test message debug", "debug")
		})

		It("should output correct info logs", func() {
			testLogging(log.Info().Msg, "Test message info", "info")
		})

		It("should output correct warn logs", func() {
			testLogging(log.Warn().Msg, "Test message warn", "warn")
		})

		It("should output correct error logs", func() {
			testLogging(log.Error().Msg, "Test message error", "error")
		})
	})
})
