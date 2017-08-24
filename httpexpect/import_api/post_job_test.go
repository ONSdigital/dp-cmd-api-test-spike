package import_api

import (
	"github.com/ONSdigital/dp-cmd-api-test-spike/config"
	"github.com/gavv/httpexpect"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestPostJob_CreatesJob(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("Given a valid json input to create a job", t, func() {

		Convey("The jobs endpoint returns 201 created ", func() {

			importAPI.POST("/jobs").WithBytes([]byte(validJSON)).
				Expect().Status(http.StatusCreated)
		})
	})
}

func TestPostJob_InvalidInput(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("Given invalid json input to create a job", t, func() {

		Convey("The jobs endpoint returns 201 created ", func() {

			importAPI.POST("/jobs").WithBytes([]byte(invalidJSON)).
				Expect().Status(http.StatusBadRequest)
		})
	})
}

