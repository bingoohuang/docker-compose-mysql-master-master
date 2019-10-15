package mysqlclusterinit

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/dealancer/validate.v2"
)

// Settings 表示舒适化MySQL集群所需要的参数结构
type Settings struct {
	Master1Addr  string   `validate:"empty=false"` // Master1的地址(IP，域名)
	Master2Addr  string   `validate:"empty=false"` // Master2的地址(IP，域名)
	SlaveAddrs   []string // Slave的地址(IP，域名)
	RootPassword string   `validate:"empty=false"` // Root用户密码
	Port         int      // MySQL 端口号
	ReplUsr      string   // 复制用用户名
	ReplPassword string   // 复制用户密码
	Debug        bool     // 测试模式，只打印SQL和HAProxy配置, 不实际执行
	LocalAddr    string   // 指定本机的IP地址，不指定则自动从网卡中获取
	MySQLCnf     string   // MySQL 配置文件的地址， 例如：/etc/mysql/conf.d/my.cnf, /etc/my.cnf
	HAProxyCfg   string   // HAProxy配置文件地址，
	// 例如：/etc/haproxy/haproxy.cfg, /etc/opt/rh/rh-haproxy18/haproxy/haproxy.cfg
	HAProxyRestartShell string // HAProxy重启命令
}

func (s *Settings) ValidateAndSetDefault() error {
	if err := validate.Validate(s); err != nil {
		logrus.Errorf("error %v", err)
		return err
	}

	if s.Port <= 0 {
		s.Port = 3306
	}

	if s.ReplUsr == "" {
		s.ReplUsr = "repl"
	}

	if s.ReplPassword == "" {
		s.ReplPassword = "984d-CE5679F93918"
	}

	if s.MySQLCnf == "" {
		s.MySQLCnf = "/etc/my.cnf"
	}

	if s.HAProxyCfg == "" {
		s.HAProxyCfg = "/etc/haproxy/haproxy.cfg"
	}

	if s.HAProxyRestartShell == "" {
		s.HAProxyRestartShell = "systemctl restart haproxy"
	}

	return nil
}

// Result 表示初始化结果
type Result struct {
	Error   error
	Sqls    []string
	HAProxy string
}

// ShowSlaveStatusBean 表示MySQL Slave Status
type ShowSlaveStatusBean struct {
	SlaveIOState         string `gorm:"column:Slave_IO_State"`
	MasterHost           string `gorm:"column:Master_Host"`
	MasterUser           string `gorm:"column:Master_User"`
	MasterPort           int    `gorm:"column:Master_Port"`
	SlaveSQLRunningState string `gorm:"column:Slave_SQL_Running_State"`
	AutoPosition         bool   `gorm:"column:Auto_Position"`
	SlaveIoRunning       string `gorm:"column:Slave_IO_Running"`
	SlaveSQLRunning      string `gorm:"column:Slave_SQL_Running"`
	MasterServerID       string `gorm:"column:Master_Server_Id"`
	SecondsBehindMaster  string `gorm:"column:Seconds_Behind_Master"`
}

// ShowVariablesBean 表示MySQL 服务器参数结果行
type ShowVariablesBean struct {
	VariableName string `gorm:"column:Variable_name"`
	Value        string `gorm:"column:Value"`
}

// Variables 表示MySQL 服务器参数
type Variables struct {
	ServerID               string `var:"server_id"`
	LogBin                 string `var:"log_bin"`
	SQLLogBin              string `var:"sql_log_bin"`
	GtidMode               string `var:"gtid_mode"`
	GtidNext               string `var:"gtid_next"`
	SlaveSkipErrors        string `var:"slave_skip_errors"`
	BinlogFormat           string `var:"binlog_format"`
	MasterInfoRepository   string `var:"master_info_repository"`
	RelayLogInfoRepository string `var:"relay_log_info_repository"`
	InnodbVersion          string `var:"innodb_version"`
}
