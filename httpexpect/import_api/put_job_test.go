package import_api

import (
	"github.com/ONSdigital/dp-cmd-api-test-spike/config"
	"github.com/gavv/httpexpect"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestPutJob(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("Given an existing job", t, func() {

		response := importAPI.POST("/jobs").WithBytes([]byte(validJSON)).
			Expect().Status(http.StatusCreated)

		jobId := response.JSON().Object().Value("job_id").String().Raw()

		Convey("An invalid job state update sent via the /jobs put endpoint returns 400 bad request", func() {

			importAPI.PUT("/jobs/{id}", jobId).WithBytes([]byte(validJobStateJSON)).
				Expect().Status(http.StatusForbidden)
		})
	})
}

func TestPutJob_InvalidStateUpdate(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("Given an existing job", t, func() {

		response := importAPI.POST("/jobs").WithBytes([]byte(validJSON)).
			Expect().Status(http.StatusCreated)

		jobId := response.JSON().Object().Value("job_id").String().Raw()

		Convey("An invalid job state update sent via the /jobs put endpoint returns 400 bad request", func() {

			importAPI.PUT("/jobs/{id}", jobId).WithBytes([]byte(invalidJobStateJSON)).
				Expect().Status(http.StatusBadRequest)
		})
	})
}

func TestPutJob_JobIDDoesNotExists(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("A put request for a job that does not exist returns 404 not found", t, func() {

		importAPI.PUT("/jobs/{id}", "2345").WithBytes([]byte(validJobStateJSON)).
			Expect().Status(http.StatusNotFound)
	})
}

