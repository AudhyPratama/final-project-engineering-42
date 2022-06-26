package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginSuccessResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
type AuthErrorResponse struct {
	Error string `json:"error"`
}

type Profile struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type ProfileSuccessResponse struct {
	Profil []Profile `json:"profile"`
}

var jwtKey = []byte("secret")

type Claims struct {
	Name     string `json:"name"`
	Email    string
	Password string `json:"password"`
	Role     string
	jwt.StandardClaims
}

func (api *API) signup(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	res, _ := api.userRepo.FetchUsersByEmail(user.Email)
	if len(res) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email already used"))
		return
	} else if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email or Password cannot be empty"))
		return
	} else if !api.userRepo.Valid(user.Email) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email is not valid"))
		return
	} else if !(user.Role == "admin" || user.Role == "user") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Role must be admin or user"))
		return
	} else {
		err := api.userRepo.Signup(user.Name, user.Email, user.Password, user.Role)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Signup failed"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Account registered successfully"))
	}

}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.userRepo.Login(user.Email, user.Password)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	userRole, _ := api.userRepo.FetchUserRole(*res)

	// Deklarasi expiry time untuk token jwt
	expirationTime := time.Now().Add(60 * time.Minute)

	// Buat claim menggunakan variable yang sudah didefinisikan diatas
	claims := &Claims{
		Email: *res,
		Role:  *userRole,
		StandardClaims: jwt.StandardClaims{
			// expiry time menggunakan time millisecond
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// return internal error ketika ada kesalahan ketika pembuatan JWT string
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set token string kedalam cookie response
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
	})

	json.NewEncoder(w).Encode(LoginSuccessResponse{Email: *res, Token: tokenString})
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	token, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// return bad request ketika field token tidak ada
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logged out"))
}

func (api *API) updatePassword(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.userRepo.FetchUsersByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(res) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email not found"))
		return
	}

	err = api.userRepo.UpdatePassword(user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Update password failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password updated successfully"))
}

func (api *API) profile(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	token, err := req.Cookie("token")
	encoder := json.NewEncoder(w)
	if err != nil {
		if err == http.ErrNoCookie {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
		// return bad request ketika field token tidak ada
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := token.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	profiles, err := api.userRepo.GetProfile(claims.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	encoder.Encode(profiles)

}

// cara membuat update profile
func (api *API) updateProfile(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	token, err := req.Cookie("token")
	encoder := json.NewEncoder(w)
	if err != nil {
		if err == http.ErrNoCookie {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
		// return bad request ketika field token tidak ada
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := token.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	profiles, err := api.userRepo.GetProfile(claims.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	encoder.Encode(profiles)
}
