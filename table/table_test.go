package table

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var filePath = "../"
var fileName = "LocalContent.txt"
var loadedTable Table

var totalOk = 3
var oKCount = 0

// run test by the order of function declared/defined in code
func TestParse(t *testing.T) {
	dataBytes, err := ioutil.ReadFile(filePath + fileName)
	if err != nil {
		t.Errorf("%v can not be read", filePath+fileName)
	}

	data := string(dataBytes)

	loadedTable, err = Parse(fileName, data)
	if err != nil {
		t.Errorf("%v parse error: %v", fileName, err)
	}
	if loadedTable.GetName() == "" {
		t.Errorf("table name: %v is empty", loadedTable.GetName())
	}
	if loadedTable.GetName() != fileName {
		t.Errorf("table name: %v is not equal to file name: %v", loadedTable.GetName(), fileName)
	}

	oKCount++
	fmt.Println("TestParse OK. ", oKCount, "/", totalOk)
}

func TestGet(t *testing.T) {
	if loadedTable.GetName() == "" {
		t.Error("no table to test Get()")
	}

	if loadedTable.Get("1000", "EN") != "System message" {
		t.Errorf("%v Get 1000, EN is %v NOT \"System message\"", fileName, loadedTable.Get("1000", "EN"))
	}
	if loadedTable.Get("1000", "yoyo") != "" {
		t.Errorf("%v Get 1000, yoyo is %v NOT empty", fileName, loadedTable.Get("1000", "yoyo"))
	}
	if loadedTable.Get("-52", "CH") != "" {
		t.Errorf("%v Get -52, CH is %v NOT empty", fileName, loadedTable.Get("-52", "CH"))
	}
	if loadedTable.Get("-52", "wuwu") != "" {
		t.Errorf("%v Get -52, wuwu is %v NOT empty", fileName, loadedTable.Get("-52", "wuwu"))
	}
	if loadedTable.Get("666", "TH") != "" {
		t.Errorf("%v Get 666, TH is %v NOT empty", fileName, loadedTable.Get("666", "TH"))
	}
	if loadedTable.Get("666", "lala") != "" {
		t.Errorf("%v Get 666, lala is %v NOT empty", fileName, loadedTable.Get("666", "lala"))
	}

	oKCount++
	fmt.Println("TestGet OK. ", oKCount, "/", totalOk)
}

func TestGetAll(t *testing.T) {
	if loadedTable.GetName() == "" {
		t.Error("no table to test GetAll()")
	}

	test := loadedTable.GetAll("1000")
	if test["CH"] != "系统讯息" {
		t.Errorf("%v GetAll 1000 [\"CH\"] is %q NOT \"系统讯息\"", fileName, test["CH"])
	}
	if test["EN"] != "System message" {
		t.Errorf("%v GetAll 1000 [\"EN\"] is %q NOT \"System message\"", fileName, test["EN"])
	}
	if test["KR"] != "System message" {
		t.Errorf("%v GetAll 1000 [\"KR\"] is %q NOT \"System message\"", fileName, test["KR"])
	}
	if test["TH"] != "System message" {
		t.Errorf("%v GetAll 1000 [\"TH\"] is %q NOT \"System message\"", fileName, test["TH"])
	}
	if test["TW"] != "系統訊息" {
		t.Errorf("%v GetAll 1000 [\"TW\"] is %q NOT \"系統訊息\"", fileName, test["TW"])
	}

	test = loadedTable.GetAll("-52")
	if test["CH"] != "" {
		t.Errorf("%v GetAll -52 [\"CH\"] is %q NOT empty", fileName, test["CH"])
	}
	if test["EN"] != "" {
		t.Errorf("%v GetAll -52 [\"EN\"] is %q NOT empty", fileName, test["EN"])
	}
	if test["KR"] != "" {
		t.Errorf("%v GetAll -52 [\"KR\"] is %q NOT empty", fileName, test["KR"])
	}
	if test["TH"] != "" {
		t.Errorf("%v GetAll -52 [\"TH\"] is %q NOT empty", fileName, test["TH"])
	}
	if test["TW"] != "" {
		t.Errorf("%v GetAll -52 [\"TW\"] is %q NOT empty", fileName, test["TW"])
	}

	test = loadedTable.GetAll("666")
	if test["CH"] != "" {
		t.Errorf("%v GetAll 666 [\"CH\"] is %q NOT empty", fileName, test["CH"])
	}
	if test["EN"] != "" {
		t.Errorf("%v GetAll 666 [\"EN\"] is %q NOT empty", fileName, test["EN"])
	}
	if test["KR"] != "" {
		t.Errorf("%v GetAll 666 [\"KR\"] is %q NOT empty", fileName, test["KR"])
	}
	if test["TH"] != "" {
		t.Errorf("%v GetAll 666 [\"TH\"] is %q NOT empty", fileName, test["TH"])
	}
	if test["TW"] != "" {
		t.Errorf("%v GetAll 666 [\"TW\"] is %q NOT empty", fileName, test["TW"])
	}

	oKCount++
	fmt.Println("TestGetAll OK. ", oKCount, "/", totalOk)
}
