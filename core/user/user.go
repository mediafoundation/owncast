package user

import (
	// "github.com/owncast/owncast/models"

	"database/sql"
	"time"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/owncast/owncast/core/data"
	"github.com/owncast/owncast/utils"
	"github.com/teris-io/shortid"

	log "github.com/sirupsen/logrus"
)

var _db *sql.DB

type User struct {
	Id           string
	AccessToken  string
	DisplayName  string
	DisplayColor float64
	CreatedAt    time.Time
	DisabledAt   *time.Time
}

func SetupUsers() {
	_db = data.GetDatabase()
	createUsersTable()
}

func createUsersTable() {
	log.Traceln("Creating users table...")

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" TEXT PRIMARY KEY,
		"access_token" string NOT NULL,
		"color" TEXT NOT NULL,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"disabled_at" DATETIME 
	);`

	stmt, err := _db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Warnln(err)
	}
}

func CreateAnonymousUser() (error, *User) {
	id := shortid.MustGenerate()
	accessToken, err := utils.GenerateAccessToken()
	if err != nil {
		log.Errorln("Unable to create access token for new user")
		return err, nil
	}

	displayName := utils.GeneratePhrase() //fmt.Sprintf("User %d", rand.Intn(99-10)+10)
	displayColor := randomcolor.GetRandomColorInHSV().Hue

	user := &User{
		Id:           id,
		AccessToken:  accessToken,
		DisplayName:  displayName,
		DisplayColor: displayColor,
	}

	setCachedIdUser(id, user)
	setCachedAccessTokenUser(accessToken, user)

	return nil, user
}

func create(user *User) error {
	tx, err := _db.Begin()
	stmt, err := tx.Prepare("INSERT INTO users(id, accessToken, displayName, displayColor, createdAt) values(?, ?, ?, ?, ?)")

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.AccessToken, user.DisplayName, user.DisplayColor, user.CreatedAt)
	if err != nil {
		panic(err)
	}

	return tx.Commit()
}

func Disable(user *User) error {
	tx, err := _db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	stmt, err := tx.Prepare("UPDATE users SET disabled_at=? WHERE id IS ?")
	stmt.Exec(user.Id, time.Now())

	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	return nil
}

// GetUserByToken will return a user by an access token.
func GetUserByToken(token string) *User {
	if user := getCachedAccessTokenUser(token); user != nil {
		return user
	}

	query := "SELECT * FROM users WHERE accessToken = ?"
	row := _db.QueryRow(query, token)
	return getUserFromRow(row)
}

// GetUserById will return a user by a user ID.
func GetUserById(id string) *User {
	if user := getCachedIdUser(id); user != nil {
		return user
	}

	query := "SELECT * FROM users WHERE id = ?"
	row := _db.QueryRow(query, id)
	return getUserFromRow(row)
}

func getUserFromRow(row *sql.Row) *User {
	var id string
	if err := row.Scan(&id); err != nil {
		return nil
	}

	return &User{
		Id: id,
	}
}
