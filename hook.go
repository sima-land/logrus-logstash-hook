package logrus_logstash

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/sirupsen/logrus"
	"net"
)

// Hook represents a connection to a Logstash instance
type Hook struct {
	conn     net.Conn
	typeName string
}

// NewHook creates a new hook to a Logstash instance, which listens on
// `protocol`://`address`.
func NewHook(protocol, address, typeName string) (*Hook, error) {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return nil, err
	}
	return &Hook{conn: conn, typeName: typeName}, nil
}

func (h *Hook) format(entry *logrus.Entry) ([]byte, error) {
	fields := make(logrus.Fields)
	for k, v := range entry.Data {
		fields[k] = v
	}
	fields["@timestamp"] = entry.Time.Format(time.RFC3339)
	fields["message"] = entry.Message
	fields["level"] = entry.Level.String()
	fields["type"] = h.typeName
	serialized, err := json.Marshal(fields)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	dataBytes, err := h.format(entry)
	if err != nil {
		return err
	}
	if _, err = h.conn.Write(dataBytes); err != nil {
		return err
	}
	return nil
}

func (h *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
