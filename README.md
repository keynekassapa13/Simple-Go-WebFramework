# SimpleGoWebFramework
a simple golang web app framework with a very simple authentication system.
The database use NoSQL MongoDB.

Client --> API --> Database

Any response to the html is sent by AJAX or the variable from the controller.
Controllers and Routes configuration are set in /routes folder.

## Available model

### Users

- Create Users (Sign Up)
- Update Users
- Delete Users
- Get Users
- Encrypted Password
- Login
- Logout

# Packages used
```
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"github.com/dgrijalva/jwt-go"
"github.com/rs/xid"
"golang.org/x/crypto/bcrypt"
"github.com/rivo/sessions"
"github.com/gorilla/mux"
```
