package common

type SelectAllColumnsFromTable struct {
	TableName string `yaml:"table-name"`
	Alias     string `yaml:"alias"`
}

type AlterTableTool struct {
	TableName    string   `yaml:"table-name"`
	AlterCmd     string   `yaml:"alter-cmd"`
	AlterColumns []string `yaml:"alter-columns"`
}
