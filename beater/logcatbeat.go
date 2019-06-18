package beater

import (
	"fmt"
	"os"
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

	// analysis os system and logcat option
	var osPreCommand string
	if bt.config.OS == "android" {
		osPreCommand = fmt.Sprintf("logcat %s", bt.config.Option)
	} else {
		osPreCommand = fmt.Sprintf("adb logcat %s", bt.config.Option)
	}
	logp.Info(osPreCommand)

	// tags
	tags := bt.config.Tags
	if len(tags) == 0 {
		hostname, err := os.Hostname()
		if err != nil {
			logp.Error(err)
		}
		tags = append(tags, hostname)
	}

	if t == "" {
		command = osPreCommand
	} else {
		command = fmt.Sprintf("%s -T '%s'", osPreCommand, t)
	}
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
					"tag":     tags,
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
		}
	}
}

// Stop stops logcatbeat.
func (bt *Logcatbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
