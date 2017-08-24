package import_api

import (
	"github.com/ONSdigital/dp-cmd-api-test-spike/config"
	"github.com/gavv/httpexpect"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestPutFile(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("Given an existing job", t, func() {

		response := importAPI.POST("/jobs").WithBytes([]byte(validJSON)).
			Expect().Status(http.StatusCreated)

		jobId := response.JSON().Object().Value("job_id").String().Raw()

		Convey("A valid file update sent via the /files put endpoint returns 200 OK", func() {

			importAPI.PUT("/jobs/{id}/files", jobId).WithBytes([]byte(validFileJSON)).
				Expect().Status(http.StatusOK)
		})
	})
}

func TestPutFile_InvalidFileJSON(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("Given an existing job", t, func() {

		response := importAPI.POST("/jobs").WithBytes([]byte(validJSON)).
			Expect().Status(http.StatusCreated)

		jobId := response.JSON().Object().Value("job_id").String().Raw()

		Convey("An invalid file update sent via the /files put endpoint returns 200 OK", func() {

			importAPI.PUT("/jobs/{id}/files", jobId).WithBytes([]byte(invalidFileJSON)).
				Expect().Status(http.StatusBadRequest)
		})
	})
}

func TestPutFile_JobIDDoesNotExists(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("A put request for a job that does not exist returns 404 not found", t, func() {

		importAPI.PUT("/jobs/{id}files", "2345").WithBytes([]byte(validFileJSON)).
			Expect().Status(http.StatusNotFound)
	})
}
