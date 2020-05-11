package Model

type Databases struct {
	Database string `gorm:"column:Database"`
}

// 表字段信息参数
type TablesInfo struct {
	TableCatakog   string `json:"table_catakog" gorm:"column:TABLE_CATALOG"`     // 表目录
	TableSchema    string `json:"table_schema" gorm:"column:TABLE_SCHEMA"`       // 数据库名称
	TableName      string `json:"table_name" gorm:"column:TABLE_NAME"`           // 表名
	TableType      string `json:"table_type" gorm:"column:TABLE_TYPE"`           // 表类型
	Engine         string `json:"engine" gorm:"column:ENGINE"`                   // 表引擎
	Version        string `json:"version" gorm:"column:VERSION"`                 // 表版本?
	RowFormat      string `json:"row_format" gorm:"column:ROW_FORMAT"`           //　行格式类型
	TableRows      string `json:"table_rows" gorm:"column:TABLE_ROWS"`           // 表条目数
	AvgRowLength   string `json:"avg_row_length" gorm:"column:AVG_ROW_LEGTH"`    // 平均每行的占用大小
	DataLength     string `json:"data_length" gorm:"column:DATA_LENGTH"`         // 数据大小
	MaxDataLength  string `json:"max_data_length" gorm:"column:MAX_DATA_LENGTH"` // ???
	IndexLength    string `json:"index_length" gorm:"column:INDEX_LENGTH"`       // 索引大小
	DataFree       string `json:"data_free" gorm:"column:DATA_FREE"`             // 数据碎片大小
	AutoIncrement  string `json:"auto_increment" gorm:"column:AUTO_INCREMENT"`   // 自增最大值
	CreateTime     string `json:"create_time" gorm:"column:CREATE_TIME"`         // 创建时间
	UpdateTime     string `json:"update_time" gorm:"column:UPDATE_TIME"`         // 更新时间
	CheckTime      string `json:"check_time" gorm:"column:CHECK_TIME"`           // 检查时间
	TableCollation string `json:"table_collation" gorm:"column:TABLE_COLLATION"` // 排序字符集
	CheckSum       string `json:"check_sum" gorm:"column:CHECKSUM"`              // 检查次数
	TableComment   string `json:"table_comment" gorm:"column:TABLE_COMMENT"`     // 表注释
}

type TableInfo struct {
	TableSchema            string `json:"table_schema" gorm:"column:TABLE_SCHEMA"`                         // 数据库名称
	TableName              string `json:"table_name" gorm:"column:TABLE_NAME"`                             // 表名
	ColumnName             string `json:"column_name" gorm:"column:COLUMN_NAME"`                           // 字段名称
	OrdinalPosition        string `json:"ordinal_position" gorm:"column:ORDINAL_POSITION"`                 // 排序
	ColumnDefault          string `json:"column_default" gorm:"column:COLUMN_DEFAULT"`                     // 字段默认值
	IsNullable             string `json:"is_nullable" gorm:"column:IS_NULLABLE"`                           //　是否可为null
	DataType               string `json:"data_type" gorm:"column:DATA_TYPE"`                               // 字段类型
	CharacterMaximumLength string `json:"character_maximum_length" gorm:"column:CHARACTER_MAXIMUM_LENGTH"` // 字符限制最大长度
	CharacterOcterLength   string `json:"character_octer_length" gorm:"column:CHARACTER_OCTER_LENGTH"`     // 字节限制最大长度
	NumeriPrecision        string `json:"numeri_precision" gorm:"column:NUMERIC_PRECISION"`                // 数字类型长度
	NumeriScale            string `json:"numeri_scale" gorm:"column:NUMERIC_SCALE"`                        // 数字类型精度
	CharacterSetName       string `json:"character_set_name" gorm:"column:CHARACTER_SET_NAME"`             // 字段字符集
	ColumnType             string `json:"column_type" gorm:"column:COLUMN_TYPE"`                           // 字段所属类型
	ColumnKey              string `json:"column_key" gorm:"column:COLUMN_KEY"`                             // 索引
	Extra                  string `json:"extra" gorm:"column:EXTRA"`                                       // 拓展描述
	ColumnComment          string `json:"column_comment" gorm:"column:COLUMN_COMMENT"`                     // 字段注释
}

type SelectDbs struct {
	PageModel
	Uuid string `json:"uuid" form:"uuid" binding:"required"`
}

type SelectTableInfo struct {
	DbName    string `json:"db_name" form:"db_name"`
	TableName string `json:"table_name" form:"table_name"`
}
