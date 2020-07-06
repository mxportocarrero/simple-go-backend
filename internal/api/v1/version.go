package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-go-backend/internal/config"

	"github.com/sirupsen/logrus"
)

type ServerVersion struct {
	Version string `json:"version"`
}

var versionJSON []byte

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writting version")
	}
}
