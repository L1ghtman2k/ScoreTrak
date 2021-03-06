package queueing

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/resolver"
	"github.com/gofrs/uuid"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type ScoringData struct {
	Service    QService
	Properties map[string]string
	Deadline   time.Time
	MasterTime time.Time
	Host       string
	RoundID    uint64
}

type QService struct {
	ID             uuid.UUID
	Group          string
	Name           string
	ReturningTopic string
}

type QCheck struct {
	Service QService
	RoundID uint64
	Passed  bool
	Log     string
	Err     string
}

type IndexedQueue struct {
	Q *QCheck
	I int
}

type Config struct {
	Use   string `default:"none"`
	Kafka struct {
	}
	NSQ struct {
		ProducerNSQD                 string   `default:"nsqd:4150"`
		IgnoreAllScoresIfWorkerFails bool     `default:"true"`
		Topic                        string   `default:"default"`
		MaxInFlight                  int      `default:"200"` //This should be more than min(NumberOfChecks, #NSQD Nodes)
		AuthSecret                   string   `default:""`
		ClientRootCA                 string   `default:""`
		ClientSSLKey                 string   `default:""`
		ClientSSLCert                string   `default:""`
		ConcurrentHandlers           int      `default:"200"`
		NSQLookupd                   []string `default:"[\"nsqlookupd:4161\"]"`
		ConsumerNSQDPool             []string `default:"[\"\"]"` //"[\"nsqd:4150\"]"`
	}
}

type MasterConfig struct {
	ReportForceRefreshSeconds uint   `default:"60"`
	ChannelPrefix             string `default:"master"`
}

func TopicFromServiceRound(roundID uint64) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return "round_" + strconv.FormatUint(roundID, 10) + "_" + strconv.Itoa(seededRand.Int()) + "_ack"
}

func CommonExecute(sd *ScoringData, execDeadline time.Time) QCheck {
	if time.Now().After(sd.Deadline) {
		return QCheck{Service: sd.Service, Passed: false, Log: "", Err: "The check arrived late to the worker. Make sure the time is synced between workers and masters, and there are enough workers to handle the load", RoundID: sd.RoundID}
	}
	executable := resolver.ExecutableByName(sd.Service.Name)

	err := exec.UpdateExecutableProperties(executable, sd.Properties)
	if err != nil {
		errLog := fmt.Sprintf("Failed to set properties for %+v. Resolved Service: %+v. Properties provided %v. See Error details for additional information", sd.Service, executable, sd.Properties)
		return QCheck{Service: sd.Service, Passed: false, Log: errLog, Err: err.Error(), RoundID: sd.RoundID}
	}

	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, execDeadline)
	defer cancel()

	e := exec.NewExec(ctx, sd.Host, executable)
	type checkRet struct {
		passed bool
		log    string
		err    error
	}
	cq := make(chan checkRet)
	go func() {
		passed, l, err := e.Execute()
		cq <- checkRet{passed: passed, log: l, err: err}
	}()
	select {
	case res := <-cq:
		var errstr string
		if res.err != nil {
			errstr = res.err.Error()
		}
		return QCheck{Service: sd.Service, Passed: res.passed, Log: res.log, Err: errstr, RoundID: sd.RoundID}
	case <-time.After(time.Until(execDeadline.Add(time.Second))):
		log.Println("Check is possibly causing resource leakage", sd.Service, execDeadline)
		panic(errors.New("check timed out. this is most likely due to services timing out"))
	}
}

type RoundTookTooLongToExecute struct {
	Msg string
}

func (e *RoundTookTooLongToExecute) Error() string { return e.Msg }
