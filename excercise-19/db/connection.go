package db

type DbBase struct {
	Connection interface{}
}

type DbAdapter interface {
	Connect() bool
	DbClientName() string
}

func GetDataBase(dbType string) DbAdapter {
	if dbType == "NoSQL" {
		return new(DbNoSql)
	}

	return new(DbSql)
}
