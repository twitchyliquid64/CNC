package user

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
  b64 "encoding/base64"
)

//Checks if a usename and password pair are valid, returning true and the user
//if it is.
func CheckAuthByPassword(username, password string, db gorm.DB)(bool,User) {
  success, tmp := GetByUsername(username, db)

  if !success {
    return false, tmp
  }

  LoadAuthMethods(&tmp, db)

  for _, authmethod := range tmp.AuthMethods {
    if authmethod.MethodType == AUTH_HASHPW {
      hashedPassword, err := b64.StdEncoding.DecodeString(authmethod.Value)

      if err != nil {
        logging.Error("auth", "Corrupted password in db. ID: ", authmethod.ID)
      } else if bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil {
        return true, tmp
      }
    }
  }

  return false, User{}
}

const HASH_COST = 10
func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    logging.Error("auth", "Error hashing password", err)
    return "", err
  }

  return b64.StdEncoding.EncodeToString(bytes), nil
}
