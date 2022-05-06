package pkg

import (
	"os"
	"referral/model"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func SetupTest(t *testing.T) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	err = model.InitDB(os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_NAME"))
	if err != nil {
		t.Fatalf(err.Error())
	}
	model.ClearTables(&model.User{}, &model.Referral{})
}

func ShutdownTest(cleanup ...func()) {
	for _, c := range cleanup {
		c()
	}
	model.CloseDB()
}
