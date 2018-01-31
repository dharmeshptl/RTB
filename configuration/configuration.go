package configuration

type SystemConfig struct {
	MessageQueue MessageQueueConfig `json:"message_queue"`
	NewRelic     NewRelicConfig     `json:"new_relic"`
	DB           AeroSpikeConfig    `json:"db"`
	App          AppConfig          `json:"app"`
	Rtb          RtbConfig          `json:"rtb"`
	JobQueue     JobQueueConfig     `json:"job_queue"`
}

type MessageQueueConfig struct {
	ConnectionString         string `json:"connection_string"`
	WinConfirmQueue          string `json:"winconfirm_queue"`
	WinConfirmConsumerNumber int    `json:"winconfirm_consumer_number"`
}

type NewRelicConfig struct {
	AppName string `json:"app_name"`
	License string `json:"license"`
}

type AeroSpikeConfig struct {
	Host              string `json:"host"`
	Port              string `json:"port"`
	Namespace         string `json:"namespace"`
	SSPSet            string `json:"ssp_set"`
	DSPSet            string `json:"dsp_set"`
	NURLSet           string `json:"nurl_set"`
	SSPApiKeyMapIdSet string `json:"ssp_apikey_id_map_set"`
	SSPLogSet         string `json:"ssp_logs_set"`
	DSPLogSet         string `json:"dsp_logs_set"`
	StatSet           string `json:"stat_set"`
}

type AppConfig struct {
	Env      string `json:"env"`
	AppPort  int    `json:"app_port"`
	LogLevel string `json:"log_level"`
	BaseUrl  string `json:"base_url"`
}

type RtbConfig struct {
	DspRequestTimeOut int `json:"dsp_request_timeout"`
}

type JobQueueConfig struct {
	WorkerNumber int `json:"worker_number"`
	QueueLength  int `json:"queue_length"`
}
