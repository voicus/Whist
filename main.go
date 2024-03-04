package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Port = 5000

func main() {
	http.FileServer(http.Dir("style"))
	http.FileServer(http.Dir("deckOfCards/SVG-cards-1.3"))
	if err := openDatabase(); err != nil {
		log.Printf("Error opening database: %v", err)
	}
	defer closeDatabase() // close the database after main1 returns

	router := mux.NewRouter()
	router.Use(authMiddleware) // Adding the auth middleware to the router

	router.HandleFunc("/rules", getRules).Methods("GET")
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/lobbies", lobbiesHandler).Methods("GET")
	router.HandleFunc("/connect", onConnect).Methods("POST")
	router.HandleFunc("/disconnect", onDisconnect).Methods("POST")
	router.HandleFunc("/getAllLobbies", getAllLobbies).Methods("GET")
	router.HandleFunc("/lobby/{lobbyName}", lobbyHandler).Methods("GET")
	router.HandleFunc("/send", sendMessageHandler).Methods("POST")
	router.HandleFunc("/login", loginGETHandler).Methods("GET")
	router.HandleFunc("/login", loginPOSTHandler).Methods("POST")
	router.HandleFunc("/setCookie", cookieTestHandler).Methods("GET")
	router.HandleFunc("/getCookies", getCookiesHandler).Methods("GET")
	router.HandleFunc("/register", registerPOSTHandler).Methods("POST")
	router.HandleFunc("/register", registerGETHandler).Methods("GET")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")
	router.HandleFunc("/style/{filename}", getStyleFile).Methods("GET")
	router.HandleFunc("/deckOfCards/SVG-cards-1.3/{filename}", getCard).Methods("GET")
	router.HandleFunc("/addToLobby", addToLobbyHandler).Methods("POST")
	router.HandleFunc("/removeFromLobby", removeFromLobbyHandler).Methods("POST")
	router.HandleFunc("/lobbyMembers/{lobbyName}", lobbyMembers).Methods("GET")
	router.HandleFunc("/manageFriends", manageFriendsHandler).Methods("GET")
	router.HandleFunc("/getFriends", getFriendsHandler).Methods("GET")
	router.HandleFunc("/addFriend/{username}", addFriendHandler).Methods("POST")
	router.HandleFunc("/acceptFriend/{username}", acceptFriendHandler).Methods("POST")
	router.HandleFunc("/declineFriend/{username}", declineFriendHandler).Methods("POST")
	router.HandleFunc("/getFriendRequests", getFriendRequestsHandler).Methods("GET")
	router.HandleFunc("/removeFriend/{username}", removeFriendHandler).Methods("POST")
	router.HandleFunc("/getUsersNotRelatedToMe", getUsersNotRelatedToMeHandler).Methods("GET")
	router.HandleFunc("/getAllUsersWithScore", getUsersWithScoreHandler).Methods("GET")
	router.HandleFunc("/getMyScore", getMyScoreHandler).Methods("GET")
	router.HandleFunc("/leaderboard", leaderboardHandler).Methods("GET")
	router.HandleFunc("/makeBid", getPlayerBid).Methods("POST")
	router.HandleFunc("/playCard", getPlayedCard).Methods("POST")
	router.HandleFunc("/winner", winnerHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		renderError(w, http.StatusNotFound)
	})
	origins := handlers.AllowedOrigins([]string{"http://proiect.home.ro"})
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET"})
	http.ListenAndServe(":"+fmt.Sprint(Port), handlers.CORS(credentials, methods, origins)(router))
}
