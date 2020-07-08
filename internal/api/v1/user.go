package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simple-go-backend/internal/api/utils"
	"simple-go-backend/internal/database"
	"simple-go-backend/internal/model"
	"time"

	"github.com/gorilla/mux"

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
	logger := logrus.WithField("func", "user.go -> GetAllUsers()")
	users := []model.User{} // empty slice

	params := r.URL.Query()
	logger.Debug("GET params", params)

	ctx := r.Context()

	cursor, err := api.DB.GetAllUsers(ctx, params.Get("limit"), params.Get("filterKey"), params.Get("filterValue"))
	if err != nil {
		logger.WithError(err).Warn("Error getting all users")
		utils.WriteError(w, http.StatusBadRequest, "Error getting all users", nil)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, user)

	}

	if err := cursor.Err(); err != nil {
		logger.WithError(err).Warn("Error getting all users with cursor")
		utils.WriteError(w, http.StatusBadRequest, "Error getting all users with cursor", nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, users)
}

func (api *UserAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	logger := logrus.WithField("func", "user.go -> UpdateUser()")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Warn("Error insert valid user info")
		utils.WriteError(w, http.StatusBadRequest, "Error insert valid user info", nil)
		return
	}

	// build new user
	user := model.User{
		ID: id,
	}
	json.Unmarshal(reqBody, &user)
	logger.Debugf("%+v", user)

	ctx := r.Context()

	err = api.DB.UpdateUser(ctx, &user)

	if err != nil {
		logger.WithError(err).Warn("Error updating user")
		utils.WriteError(w, http.StatusBadRequest, "Error updating", nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, user)
}
