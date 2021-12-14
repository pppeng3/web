package mysql

import "testing"

func TestMysqlInstance(t *testing.T) {
	mysqlClient := Instance()
	t.Log(mysqlClient)
}
