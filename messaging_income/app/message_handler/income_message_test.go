package message_handler

import (
	"testing"
)

func TestIsIncomePattern(t *testing.T) {
	t.Run("It should return ok if message is เงินเดิือน +10000", func(t *testing.T) {
		r := isIncomePattern("เงินเดือน +10000")
		if r != true {
			t.Error("Expect", true, "Actual", r)
		}
	})

	t.Run("It should return ok if message is เงินเดิือน +10k", func(t *testing.T) {
		r := isIncomePattern("เงินเดิือน +10k")
		if r != true {
			t.Error("Expect", true, "Actual", r)
		}
	})

	t.Run("It should return ok if message is +10k เงินเดิือน", func(t *testing.T) {
		r := isIncomePattern("+10k เงินเดิือน")
		if r != true {
			t.Error("Expect", true, "Actual", r)
		}
	})

	t.Run("It should return ok if message is +10000 เงินเดิือน", func(t *testing.T) {
		r := isIncomePattern("+10000 เงินเดิือน")
		if r != true {
			t.Error("Expect", true, "Actual", r)
		}
	})

	t.Run("It should return ok if message is -60 กินข้าว", func(t *testing.T) {
		r := isIncomePattern("-60 กินข้าว")
		if r != false {
			t.Error("Expect", false, "Actual", r)
		}
	})

	t.Run("It should return ok if message is กินข้าว -60", func(t *testing.T) {
		r := isIncomePattern("กินข้าว -60")
		if r != false {
			t.Error("Expect", false, "Actual", r)
		}
	})
}

func TestExtractValue(t *testing.T) {
	t.Run("it should return 100 when enter เพื่อนจ่ายหนี้ +100", ExtractValueTestCase("เพื่อนจ่ายหนี้ +100", 100))
	t.Run("it should return 100 when enter +100 เพื่อนจ่ายหนี้", ExtractValueTestCase("+100 เพื่อนจ่ายหนี้", 100))
	t.Run("it should return 10000 when enter เงินเดือน +10k", ExtractValueTestCase("เงินเดือน +10k", 10000))
	t.Run("it should return 10000 when enter เงินเดือน +10K", ExtractValueTestCase("เงินเดือน +10K", 10000))
	t.Run("it should return 10000 when enter เงินเดือน +10m", ExtractValueTestCase("เงินเดือน +10m", 10000000))
	t.Run("it should return 10000 when enter เงินเดือน +10M", ExtractValueTestCase("เงินเดือน +10M", 10000000))
	t.Run("it should return error when enter +100 เพื่อนจ่ายหนี้ +100", func(t *testing.T) {
		_, err := extractValue("+100 เพื่อนจ่ายหนี้ +1000")
		if err == nil {
			t.Error("expect to have a error")
		}
	})
}

func ExtractValueTestCase(s string, expect float64) func(*testing.T) {
	return func(t *testing.T) {
		v, err := extractValue(s)
		if err != nil {
			t.Error(err)
		}

		if v != expect {
			t.Error("expect", expect, "actual", v)
		}
	}
}
