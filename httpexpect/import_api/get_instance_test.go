package import_api

import (
	"github.com/ONSdigital/dp-cmd-api-test-spike/config"
	"github.com/gavv/httpexpect"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
	"os"
)

func TestMain(m *testing.M) {

	config.Init()
	os.Exit(m.Run())
}

func TestGetInstance(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("Given an existing job", t, func() {

		json := importAPI.POST("/jobs").WithBytes([]byte(validJSON)).
			Expect().Status(http.StatusCreated).JSON().Object()

		expectedInstanceID := json.Value("instances").Array().Element(0).
			Object().Value("id").String().Raw()

		Convey("Getting the instance data returns the expected values", func() {

			json = importAPI.GET("/instances/{id}", expectedInstanceID).
				Expect().Status(http.StatusOK).JSON().Object()

			actualInstanceID := json.Value("instance_id").String().Raw()
			actualState := json.Value("state").String().Raw()

			So(actualInstanceID, ShouldEqual, expectedInstanceID)
			So(actualState, ShouldEqual, "created")
		})
	})
}

func TestGetInstance_InstanceIDDoesNotExists(t *testing.T) {

	importAPI := httpexpect.New(t, config.ImportAPIURL())

	Convey("A get request for an instance that does not exist returns 404 not found", t, func() {

		importAPI.GET("/instances/{id}", "76543").
			Expect().Status(http.StatusNotFound)
	})
}
