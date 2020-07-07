package v1

import (
	"encoding/json"
	"net/http"
	"simple-go-backend/internal/api/utils"
	"simple-go-backend/internal/database"
	"simple-go-backend/internal/model"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAPI struct {
	DB database.Database
}

type CreateUserParams struct {
	model.User
}

func (api *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("func", "user.go -> CreateUser()")
	var createUserParams CreateUserParams
	var err error

	err = json.NewDecoder(r.Body).Decode(&createUserParams)
	if err != nil {
		logger.WithError(err).Warn("could not decode params")
		utils.WriteError(w, http.StatusBadRequest, "could not decode parameters", map[string]string{
			"error": err.Error(),
		})
		return
	}

	// TODO print log to backend console

	if err := createUserParams.ValidateEmail(createUserParams.Email); err != nil {
		logger.WithError(err).Warn("Input Email is invalid")
		utils.WriteError(w, http.StatusBadRequest, "Input Email is invalid", nil)
		return
	}

	if err := createUserParams.ValidatePhone(createUserParams.Phone); err != nil {
		logger.WithError(err).Warn("Input Phone is invalid")
		utils.WriteError(w, http.StatusBadRequest, "Input Phone is invalid", nil)
		return
	}

	newUser := &model.User{
		Name:      createUserParams.Name,
		Email:     createUserParams.Email,
		Phone:     createUserParams.Phone,
		CreatedAt: time.Now(),
		DeletedAt: time.Now(),
		Deleted:   false,
	}

	// use the request context
	ctx := r.Context()

	result, err := api.DB.CreateUser(ctx, newUser)

	if err != nil {
		logger.WithError(err).Warn("Error creating user")
		utils.WriteError(w, http.StatusBadRequest, "Error creating user", nil)
		return
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		newUser.ID = oid
	}

	utils.WriteResponse(w, http.StatusCreated, newUser)
}

func (api *UserAPI) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func (api *UserAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
