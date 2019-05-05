package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/Fize/logcatbeat/config"
)

// Logcatbeat configuration.
type Logcatbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of logcatbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Logcatbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts logcatbeat.
func (bt *Logcatbeat) Run(b *beat.Beat) error {
	logp.Info("logcatbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}
	var command string
	t := getRegisterTime()

	var osPreCommand string
	if bt.config.OS == "android" {
		osPreCommand = "logcat python:S"
	} else {
		osPreCommand = "adb logcat python:S"
	}

	if t == "" {
		command = osPreCommand
	} else {
		command = fmt.Sprintf("%s -T '%s'", osPreCommand, t)
	}
	ticker := time.NewTicker(bt.config.Period)
	msgs := make(chan string, 256)
	logcat := NewExecutor(command)
	go logcat.Run(msgs)
	go func() {
		for m := range msgs {
			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type":    "android-log",
					"message": m,
				},
			}
			bt.client.Publish(event)
			err := setRegisterTime(m)
			if err != nil {
				logp.Err(err.Error())
			}
		}
	}()
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
	}
}

// Stop stops logcatbeat.
func (bt *Logcatbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
