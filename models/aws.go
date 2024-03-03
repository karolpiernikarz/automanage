package models

import "time"

type PolicyDocument struct {
	Version   string
	Statement []StatementEntry
}

type StatementEntry struct {
	Effect    string
	Action    []string
	Resource  []string
	Condition map[string]interface{}
}

type AmazonSesFeedback struct {
	EventType string `json:"eventType"`
	Bounce    struct {
		FeedbackId        string `json:"feedbackId"`
		BounceType        string `json:"bounceType"`
		BounceSubType     string `json:"bounceSubType"`
		BouncedRecipients []struct {
			EmailAddress   string `json:"emailAddress"`
			Action         string `json:"action"`
			Status         string `json:"status"`
			DiagnosticCode string `json:"diagnosticCode"`
		} `json:"bouncedRecipients"`
		Timestamp    time.Time `json:"timestamp"`
		ReportingMTA string    `json:"reportingMTA"`
	} `json:"bounce,omitempty"`
	Mail struct {
		Timestamp        time.Time `json:"timestamp"`
		Source           string    `json:"source"`
		SourceArn        string    `json:"sourceArn"`
		SendingAccountId string    `json:"sendingAccountId"`
		MessageId        string    `json:"messageId"`
		Destination      []string  `json:"destination"`
		HeadersTruncated bool      `json:"headersTruncated"`
		Headers          []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`
		CommonHeaders struct {
			From      []string `json:"from"`
			Date      string   `json:"date"`
			To        []string `json:"to"`
			MessageId string   `json:"messageId"`
			Subject   string   `json:"subject"`
		} `json:"commonHeaders"`
		Tags struct {
			SesSourceTlsVersion []string `json:"ses:source-tls-version"`
			SesOperation        []string `json:"ses:operation"`
			SesConfigurationSet []string `json:"ses:configuration-set"`
			SesRecipientIsp     []string `json:"ses:recipient-isp"`
			SesSourceIp         []string `json:"ses:source-ip"`
			SesFromDomain       []string `json:"ses:from-domain"`
			SesSenderIdentity   []string `json:"ses:sender-identity"`
			SesCallerIdentity   []string `json:"ses:caller-identity"`
		} `json:"tags"`
	} `json:"mail"`
}
