package model

import (
	_ "embed"
)

var (
	//go:embed sql/testdata.sql
	testdataScript string
	//go:embed sql/clear.sql
	clearScript string
)

func setupDB() {
	result := db.Exec(testdataScript)
	if result.Error != nil {
		panic(result.Error)
	}
}

func teardownDB() {
	result := db.Exec(clearScript)
	if result.Error != nil {
		panic(result.Error)
	}
}
