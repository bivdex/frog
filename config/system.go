package config

type System struct {
	RouterPrefix           string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr                   int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	BotToken               string `mapstructure:"bot-token" json:"bot-token" yaml:"bot-token"`
	ChatID                 string `mapstructure:"chat-id" json:"chat-id" yaml:"chat-id"`
	MasterAddress          string `mapstructure:"master-address" json:"master-address" yaml:"master-address"`
	MasterPK               string `mapstructure:"master-pk" json:"master-pk" yaml:"master-pk"`
	TrxfastUsername        string `mapstructure:"trxfast-username" json:"trxfast-username" yaml:"trxfast-username"`
	TrxfastPassword        string `mapstructure:"trxfast-password" json:"trxfast-password" yaml:"trxfast-password"`
	TRON_FULL_NODE         string `mapstructure:"tron-full-node" json:"tron-full-node" yaml:"tron-full-node"`
	LOCAL_DB_ENABLE        bool   `mapstructure:"local-db-enable" json:"local-db-enable" yaml:"local-db-enable"`
	TRON_MONITOR_ADDRESSES string `mapstructure:"tron-monitor-addresses" json:"tron-monitor-addresses" yaml:"tron-monitor-addresses"`
	BOT_ENABLE             bool   `mapstructure:"bot-enable" json:"bot-enable" yaml:"bot-enable"`
	SUOJINSUANFA           int64  `mapstructure:"suojinsuanfa" json:"suojinsuanfa" yaml:"suojinsuanfa"`
	DAY_LIMIT              int    `mapstructure:"day-limit" json:"day-limit" yaml:"day-limit"`
	TRONGRID_KEYS          string `mapstructure:"trongrid-keys" json:"trongrid-keys" yaml:"trongrid-keys"`
	TX_COUNT_RANGE         string `mapstructure:"tx-count-range" json:"tx-count-range" yaml:"tx-count-range"`
	TRANSFER_AMOUNT_RANGE  string `mapstructure:"transfer-amount-range" json:"transfer-amount-range" yaml:"transfer-amount-range"`
	BALANCE_RANGE          string `mapstructure:"balance-range" json:"balance-range" yaml:"balance-range"`
	AVG_AMOUNT             string `mapstructure:"avg-amount" json:"avg-amount" yaml:"avg-amount"`
	TRON_GRPC_NODE         string `mapstructure:"tron-grpc-node" json:"tron-grpc-node" yaml:"tron-grpc-node"`

	FROM_ADDRESS_PART_RULE          string `mapstructure:"from-address-part-rule" json:"from-address-part-rule" yaml:"from-address-part-rule"`
	TRANSFER_TX_COUNT_MIN           int    `mapstructure:"transfer-tx-count-min" json:"transfer-tx-count-min" yaml:"transfer-tx-count-min"`
	IGNORE_TRANSFER_USDT_AMOUNT_MIN int    `mapstructure:"ignore-transfer-usdt-amount-min" json:"ignore-transfer-usdt-amount-min" yaml:"ignore-transfer-usdt-amount-min"`
	ENABLE_MAIN_ADDR_TASK           bool   `mapstructure:"enable-main-addr-task" json:"enable-main-addr-task" yaml:"enable-main-addr-task"`
	GET_FROM_ADDR_INTERVAL          int    `mapstructure:"get-from-addr-interval" json:"get-from-addr-interval" yaml:"get-from-addr-interval"`
}
