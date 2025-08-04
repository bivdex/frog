package config

type Server struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`

	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql_A     Mysql `mapstructure:"mysql_a" json:"mysql" yaml:"mysql"`
	Mysql_B     Mysql `mapstructure:"mysql_b" json:"mysql" yaml:"mysql"`
	Mysql_C     Mysql `mapstructure:"mysql_c" json:"mysql" yaml:"mysql"`
	Mysql_D     Mysql `mapstructure:"mysql_d" json:"mysql" yaml:"mysql"`
	Mysql_E     Mysql `mapstructure:"mysql_e" json:"mysql" yaml:"mysql"`
	Mysql_F     Mysql `mapstructure:"mysql_f" json:"mysql" yaml:"mysql"`
	Mysql_G     Mysql `mapstructure:"mysql_g" json:"mysql" yaml:"mysql"`
	Mysql_Local Mysql `mapstructure:"mysql_local" json:"mysql" yaml:"mysql"`

	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
