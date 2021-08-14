package config_test

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/ahsanulks/waitress/config"
	"go.uber.org/zap"
)

func TestLog_Store(t *testing.T) {
	zapLog, _ := zap.NewProduction()
	logger := config.NewLog(zapLog)
	type args struct {
		err     error
		message string
		options map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "error nil",
			args: args{nil, "ok", map[string]interface{}{
				"tags": []string{"test"},
			}},
		},
		{
			name: "error",
			args: args{errors.New("error"), "error", map[string]interface{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)
			defer func() {
				log.SetOutput(os.Stderr)
			}()
			logger.Store(tt.args.err, tt.args.message, tt.args.options)
			t.Log(buf.String())
		})
	}
}
