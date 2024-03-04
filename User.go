package main

import (
	"errors"
	"log"
	"net/http"
)

type User struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	passwordHash string
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreate struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserWithScore struct {
	Username  string `json:"username"`
	Rating    int    `json:"rating"`
	GamesWon  int    `json:"gamesWon"`
	GamesLost int    `json:"gamesLost"`
}

// In momentul in care se creeaza un user, va fi initiat si scorul acestuia cu valori nule (rating, games_won, games_lost)
func createUser(user *UserCreate) error {
	if err := DB.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password).Err(); err != nil {
		return err
	}
	return nil
}

func getUserByUsername(username string) (*User, error) {
	var myUser User
	if err := DB.QueryRow("SELECT username, email, password FROM users WHERE username = $1", username).Scan(&myUser.Username, &myUser.Email, &myUser.passwordHash); err != nil {
		return nil, err
	}
	return &myUser, nil
}

func getUserByEmail(email string) (*User, error) {
	var myUser User
	if err := DB.QueryRow("SELECT username, email, password FROM users WHERE email = $1", email).Scan(&myUser.Username, &myUser.Email, &myUser.passwordHash); err != nil {
		return nil, err
	}
	return &myUser, nil
}

func getFriendsOfUser(user User) ([]User, error) {
	var user_id int
	if err := DB.QueryRow("select id from users where username = $1", user.Username).Scan(&user_id); err != nil {
		return make([]User, 0), err
	}
	friends := make([]User, 0)
	rows, err := DB.Query("select username from users where id in ((select user_id1 from are_friends where user_id2 = $1 and confirmed_1 and confirmed_2) union (select user_id2 from are_friends where user_id1 = $1 and confirmed_1 and confirmed_2))", user_id)
	if err != nil {
		return make([]User, 0), err
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return make([]User, 0), err
		}
		friends = append(friends, User{Username: username})
	}

	return friends, nil
}

func getFriendRequestsOfUser(user User) ([]User, error) {
	var user_id int
	if err := DB.QueryRow("select id from users where username = $1", user.Username).Scan(&user_id); err != nil {
		return make([]User, 0), err
	}
	friend_requests := make([]User, 0)
	rows, err := DB.Query("select username from users where id in ((select user_id1 from are_friends where user_id2 = $1 and confirmed_1 and not confirmed_2) union (select user_id2 from are_friends where user_id1 = $1 and not confirmed_1 and confirmed_2))", user_id)
	if err != nil {
		return make([]User, 0), err
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return make([]User, 0), err
		}
		friend_requests = append(friend_requests, User{Username: username})
	}
	return friend_requests, nil
}

func getUsersNotRelatedToMe(user User) ([]User, error) {
	var user_id int
	if err := DB.QueryRow("select id from users where username = $1", user.Username).Scan(&user_id); err != nil {
		return make([]User, 0), err
	}

	usersNotRelated := make([]User, 0)
	rows, err := DB.Query("select username from users where id not in (select user_id1 from are_friends where user_id2 = $1 union select user_id2 from are_friends where user_id1 = $1)", user_id)
	if err != nil {
		return make([]User, 0), err
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return make([]User, 0), err
		}
		if username == user.Username {
			continue
		}
		usersNotRelated = append(usersNotRelated, User{Username: username})
	}
	return usersNotRelated, nil
}

func areFriends(user1 User, user2 User) (bool, error) {
	var result bool
	var err error
	var user1_id int
	var user2_id int
	if err = DB.QueryRow("select id from users where username = $1", user1.Username).Scan(&user1_id); err != nil {
		return false, err
	}

	if err = DB.QueryRow("select id from users where username = $1", user2.Username).Scan(&user2_id); err != nil {
		return false, err
	}

	if err = DB.QueryRow("select exists(select 1 from are_friends where ((user_id1, user_id2) = ($1, $2) or (user_id1, user_id2) = ($2, $1)) and confirmed_1 and confirmed_2)", user1_id, user2_id).Scan(&result); err != nil {
		return false, err
	}

	return result, nil
}

func sendFriendRequest(sender User, receiver User) error {
	var sender_id int
	var receiver_id int

	if err := DB.QueryRow("select id from users where username = $1", sender.Username).Scan(&sender_id); err != nil {
		return err
	}

	if err := DB.QueryRow("select id from users where username = $1", receiver.Username).Scan(&receiver_id); err != nil {
		return err
	}

	if sender_id == receiver_id {
		return errors.New("Cannot send friend request to yourself!")
	}

	if _, err := DB.Exec("insert into are_friends(user_id1, user_id2, confirmed_1, confirmed_2) values($1, $2, true, false)", sender_id, receiver_id); err != nil {
		return err
	}

	return nil
}

func acceptFriendRequest(accepter User, other User) error {
	var accepter_id int
	var other_id int

	if err := DB.QueryRow("select id from users where username = $1", accepter.Username).Scan(&accepter_id); err != nil {
		return err
	}

	if err := DB.QueryRow("select id from users where username = $1", other.Username).Scan(&other_id); err != nil {
		return err
	}

	if accepter_id == other_id {
		return errors.New("Cannot accept friend request from yourself!")
	}

	result, err := DB.Exec("update are_friends set confirmed_2 = true where user_id1 = $1 and user_id2 = $2", other_id, accepter_id)
	if err != nil {
		return err
	}
	affected_rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected_rows == 0 {
		return errors.New("No friend request to accept!")
	}
	return nil
}

func declineFriendRequest(rejecter User, other User) error {
	var rejecter_id int
	var other_id int

	if err := DB.QueryRow("select id from users where username = $1", rejecter.Username).Scan(&rejecter_id); err != nil {
		return err
	}

	if err := DB.QueryRow("select id from users where username = $1", other.Username).Scan(&other_id); err != nil {
		return err
	}

	if rejecter_id == other_id {
		return errors.New("Cannot decline a friend request from yourself!")
	}

	result, err := DB.Exec("delete from are_friends where (user_id1, userid_2) = ($1, $2) or (user_id1, userid_2) = ($2, $1)", other_id, rejecter_id)
	if err != nil {
		return err
	}

	affected_rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected_rows == 0 {
		return errors.New("No friend request to decline!")
	}
	return nil
}

func unfriend(user1 User, user2 User) error {
	are_friends, err := areFriends(user1, user2)
	if err != nil {
		return err
	}
	if !are_friends {
		return errors.New("Users are not friends!")
	}

	var user1_id int
	var user2_id int

	if err := DB.QueryRow("select id from users where username = $1", user1.Username).Scan(&user1_id); err != nil {
		return err
	}

	if err := DB.QueryRow("select id from users where username = $1", user2.Username).Scan(&user2_id); err != nil {
		return err
	}

	if user1_id == user2_id {
		return errors.New("Cannot unfriend yourself!")
	}

	if _, err := DB.Exec("delete from are_friends where (user_id1, user_id2) = ($1, $2) or (user_id1, user_id2) = ($2, $1)", user1_id, user2_id); err != nil {
		return err
	}
	return nil
}

func getAllUsersWithScoresDescending() ([]UserWithScore, error) {
	usersWithScore := make([]UserWithScore, 0)
	rows, err := DB.Query("select username, games_won, games_lost, rating from users u join scores s on u.id = s.user_id order by rating desc;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		var games_won int
		var games_lost int
		var rating int
		if err := rows.Scan(&username, &games_won, &games_lost, &rating); err != nil {
			return nil, err
		}
		usersWithScore = append(usersWithScore, UserWithScore{Username: username, GamesWon: games_won, GamesLost: games_lost, Rating: rating})
	}

	return usersWithScore, nil
}

func getScoreInfoOfUser(user User) (UserWithScore, error) {
	var username string
	var games_won int
	var games_lost int
	var rating int
	if err := DB.QueryRow("select username, games_won, games_lost, rating from users u join scores s on u.id = s.user_id where username = $1;", user.Username).Scan(&username, &games_won, &games_lost, &rating); err != nil {
		return UserWithScore{}, err
	}
	return UserWithScore{Username: username, GamesWon: games_won, GamesLost: games_lost, Rating: rating}, nil
}

// NU este nevoie de o functie care sa actualizeze rating-ul unui jucator; acesta se actualizeaza automat la cresterea numarului de
// jocuri castigate sau pierdute dupa formula: RATING = 10*GAMES_WON - 3*GAMES_LOST
func incrNumberOfGamesWon(user User) error {
	result, err := DB.Exec("update scores set games_won = games_won + 1 where user_id = (select id from users where username = $1);", user.Username)
	if err != nil {
		return err
	}
	affected_rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected_rows == 0 {
		return errors.New("No row updated!")
	}

	return nil
}

func incrNumberOfGamesLost(user User) error {
	result, err := DB.Exec("update scores set games_lost = games_lost + 1 where user_id = (select id from users where username = $1);", user.Username)
	if err != nil {
		return err
	}
	affected_rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected_rows == 0 {
		return errors.New("No row updated!")
	}

	return nil
}

// Test function for user creation
func testUserCreate() { // modify to test the creation of other users
	testUser := UserCreate{
		Email:           "email@gmail.com",
		Username:        "username",
		Password:        "password",
		ConfirmPassword: "password",
	}

	if status := register(&testUser); status != http.StatusOK {
		log.Printf("Couldn't create user: %v (Status code: %v)\n", testUser.Username, status)
	} else {
		log.Printf("Created user: %v\n", testUser.Username)
	}
}
