package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
)

type ReportClient struct {
	s ScoretrakClient
}

func NewReportClient(c ScoretrakClient) *ReportClient {
	return &ReportClient{c}
}

func (c ReportClient) Get() (*report.Report, error) {
	conf := &report.Report{}
	err := c.s.GenericGet(conf, fmt.Sprintf("/report"))
	if err != nil {
		return nil, err
	}
	return conf, nil
}
