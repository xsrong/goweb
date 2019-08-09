package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func setup() (users []User) {
	file, _ := os.Open("test_user_data.json")
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &users)
	return users
}

func TestUserCreate(t *testing.T) {
	users := setup()
	err := users[0].Create() // 测试Email为空的情况
	if err == nil {
		t.Error("excepted get an error but no error occured.")
	}

	err = users[1].Create() // 测试Passworld为空的情况
	if err == nil {
		t.Error("excepted get an error but no error occured.")
	}

	err = users[2].Create() // 测试Username为空的情况
	if err == nil {
		t.Error("excepted get an error but no error occured.")
	}

	err = users[3].Create() // 测试Message为空的情况，注：Message允许为空
	if err != nil {
		t.Errorf("excepted no error but an error occured:\n %s\n", err)
	}

	err = users[4].Create() // 测试都不为空的情况
	if err != nil {
		t.Errorf("excepted no error but an error occured:\n %s\n", err)
	}

	err = users[5].Create() // 测试Email发生重复的情况
	if err == nil {
		t.Error("excepted get an error but no error occured.")
	}

	err = users[6].Create() // 测试Username发生重复的情况
	if err == nil {
		t.Error("excepted get an error but no error occured.")
	}
}
