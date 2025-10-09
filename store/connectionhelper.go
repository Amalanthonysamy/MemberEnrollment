package store

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnectionHelper() (*mongo.Client, error) {
	// MongoDB connection helper
	_ = godotenv.Load(".env")
	uri := os.Getenv("uri")
	// Placeholder for MongoDB connection and insertion logic
	//timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB connection URI
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	//defer client.Disconnect(ctx)
	client.Database("MemberEnrollment").Collection("Members")
	return client, nil

}

// CreateMember godoc
// @Summary      Create a new member
// @Description  Adds a new member
// @Tags         members
// @Accept       json
// @Produce      json
// @Param        member  body      Member  true  "Member to create"
// @Success      201    {object}  Member
// @Failure      400    {object}  map[string]string "Invalid input"
// @Router       /members/v1.0 [post]
func SaveMember(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for save member logic
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}

	//request to db
	collection := mongoClient.Database("MemberEnrollment").Collection("Members")
	var member Member
	json.NewDecoder(request.Body).Decode(&member)
	_, err = collection.InsertOne(context.TODO(), member)
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(member)

}

// GetMembers godoc
// @Summary      Get all members
// @Description  Returns list of members
// @Tags         members
// @Accept       json
// @Produce      json
// @Success      200  {array}   Member
// @Router       /members/v1.0 [get]
func GetMembers(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for get members logic
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("MemberEnrollment").Collection("Members")
	cursor, err := collection.Find(context.TODO(), struct{}{})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var members []Member
	if err = cursor.All(context.TODO(), &members); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(members)

}

// GetMemberById godoc
// @Summary     Get details of requested member
// @Description Get details of requested member
// @Tags        members
// @Accept      json
// @Produce     json
// @Param       memberid path int true "ID of the Member"
// @Success     200 {object} Member
// @Failure     400 {object} map[string]string "Invalid ID supplied"
// @Failure     404 {object} map[string]string "Member not found"
// @Router      /members/v1.0/{memberid} [get]
func GetMemberByID(writer http.ResponseWriter, request *http.Request) {

	idStr := request.PathValue("memberid")
	memberID, err := strconv.Atoi(idStr)
	if err != nil || memberID <= 0 {
		http.Error(writer, `{"error":"Invalid ID supplied"}`, http.StatusBadRequest)
		return
	}
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("MemberEnrollment").Collection("Members")
	//get member by id
	var member Member
	if err := collection.FindOne(context.TODO(), bson.M{"memberid": memberID}).Decode(&member); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(writer, `{"error":"Member not found"}`, http.StatusNotFound)
			return
		}
		http.Error(writer, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(member)

}

// UpdateMember godoc
// @Summary     Update existing member
// @Description Update existing member with the input payload
// @Tags        members
// @Accept      json
// @Produce     json
// @Param       memberid path int   true "ID of the Member to update"
// @Param       member   body Member true "Updated member fields"
// @Success     200 {object} Member
// @Failure     400 {object} map[string]string "Invalid ID supplied"
// @Failure     404 {object} map[string]string "Member not found"
// @Router      /members/v1.0/{memberid} [put]
func UpdateMember(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for delete member logic
	idStr := request.PathValue("memberid")
	memberID, err := strconv.Atoi(idStr)
	if err != nil || memberID <= 0 {
		http.Error(writer, `{"error":"Invalid ID supplied"}`, http.StatusBadRequest)
		return
	}
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("MemberEnrollment").Collection("Members")
	var member Member
	json.NewDecoder(request.Body).Decode(&member)
	_, err = collection.UpdateOne(context.TODO(), bson.M{"memberid": memberID}, bson.M{"$set": member})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(member)

}

// DeleteMemberById godoc
// @Summary Delete requested member
// @Description Delete requested member
// @Tags members
// @Accept  json
// @Produce  json
// @Param memberid path int true "ID of the Member"
// @Success 200 {object} Member
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Member not found"
// @Router /members/v1.0/{memberid} [delete]
func DeleteMember(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for delete member logic
	idStr := request.PathValue("memberid")
	memberID, err := strconv.Atoi(idStr)
	if err != nil || memberID <= 0 {
		http.Error(writer, `{"error":"Invalid ID supplied"}`, http.StatusBadRequest)
		return
	}
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("MemberEnrollment").Collection("Members")
	_, err = collection.DeleteOne(context.TODO(), map[string]interface{}{"memberid": memberID})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
