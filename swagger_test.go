package swagger_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/marstr/swagger"
)

func TestSwagger_UnmarshalJSON(t *testing.T) {
	testcases := []string{
		filepath.Join("testdata", "petstore", "swagger.json"),
		filepath.Join("testdata", "azure-rest-api-specs", "arm-hdinsight", "2015-03-01-preview", "swagger", "operations.json"),
	}

	for _, tc := range testcases {
		t.Run(tc, func(t *testing.T) {
			contents, err := ioutil.ReadFile(tc)
			if err != nil {
				t.Error(err)
			}

			var unmarshaled swagger.Swagger
			err = json.Unmarshal(contents, &unmarshaled)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
