package main

import (
	"log"
	"net/http"

	_ "MemberEnrollment/APIDevelopment/docs"
	"MemberEnrollment/APIDevelopment/store"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Claim API
// @version 1.0
// @description This is api service for managing Claims
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email parameswaribala@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7177
// @BasePath /
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /members/v1.0", store.GetMembers)
	mux.HandleFunc("POST /members/v1.0", store.SaveMember)
	mux.HandleFunc("GET /members/v1.0/{memberid}", store.GetMemberByID)
	mux.HandleFunc("PUT /members/v1.0/{memberid}", store.UpdateMember)
	mux.HandleFunc("DELETE /members/v1.0/{memberid}", store.DeleteMember)
	// Swagger UI served at /swagger/
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Your own handlers
	// mux.HandleFunc("/members", membersHandler)

	log.Println("Server running at http://localhost:7177")
	log.Fatal(http.ListenAndServe(":7177", mux))

}
