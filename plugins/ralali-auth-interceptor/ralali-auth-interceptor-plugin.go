package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-contrib/uuid"
)

// const outputHeaderName = "X-User-ID"

func init() {
	fmt.Println("loading ralali-auth verificator plugin!!")
}

func main() {}

// HandlerRegisterer is the name of the symbol krakend looks up to try and register plugins
var HandlerRegisterer = registerer("ralaliAuthInterceptorPlugin")

type registerer string

// type OauthData struct {
// 	Success *SuccesData `json:"success"`
// }

// type SuccesData struct {
// 	UserSession *TokenData `json:"user_session"`
// }

// type TokenData struct {
// 	Token *TokenDetail `json:"token"`
// }

// type TokenDetail struct {
// 	UserSession string `json:"user_session"`
// 	ID          int    `json:"id"`
// 	UserID      string `json:"user_id"`
// 	ClientID    string `json:"client_id"`
// 	FCMToken    string `json:"fcm_token"`
// }

const outputHeaderName = "X-Friend-User"
const pluginName = "ralaliAuthInterceptorPlugin"

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(
		context.Context,
		map[string]interface{},
		http.Handler) (http.Handler, error),
)) {
	f(pluginName, r.registerHandlers)
}

func (r registerer) registerHandlers(ctx context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {

	// oauthURL, ok := extra["oauthURL"].(string)
	// if !ok {
	// 	panic(errors.New("incorrect config").Error())
	// }

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// token := r.Header.Get("x-access-token")
		// if len(token) == 0 {
		// 	http.Error(w, "", http.StatusForbidden)
		// 	return
		// }

		// rq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://oauthdev.ralali.xyz/v1/profile"), nil)
		// // rq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s", oauthURL), nil)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		// 	return
		// }
		// rq.Header.Set("Content-Type", "application/json")
		// rq.Header.Set("x-access-token", token)

		// rs, err := client.Do(rq)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusUnauthorized)
		// 	return
		// }
		// defer rs.Body.Close()

		// rsBodyBytes, err := ioutil.ReadAll(rs.Body)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusNotAcceptable)
		// 	return
		// }
		// stringBody := string(rsBodyBytes)
		// data := &OauthData{}
		// errParse := json.Unmarshal([]byte(stringBody), data)
		// if errParse != nil {
		// 	http.Error(w, err.Error(), http.StatusUnauthorized)
		// 	return
		// }
		// if nil == errParse {
		// 	fmt.Println("found data")
		// 	fmt.Println(data.Success.UserSession.Token.ID)
		// 	fmt.Println(data.Success.UserSession.Token.ClientID)
		// 	fmt.Println(data.Success.UserSession.Token.FCMToken)
		// 	fmt.Println(data.Success.UserSession.Token.UserID)
		// }

		// choose it or below
		// r.Header.Set("Coba-Set-P", "Value-Set-P")
		// r.Header.Add("Coba-Add-P", "Value-Add-P")
		// w.Header().Add("Coba-H-Add-P", "Value-H-Add-P")
		// w.Header().Set("Coba-H-Set-P", "Value-H-Set-P")
		// handler.ServeHTTP(w, r)
		fmt.Println("adding additional header request")
		r.Header.Set("X-Request-Id", uuid.NewV4().String())
		handler.ServeHTTP(w, r)

		// r2 := new(http.Request)
		// *r2 = *r
		// r2.Header.Set(outputHeaderName, "Sample Header Value")
		// handler.ServeHTTP(w, r2)

		// r2 := new(http.Request)
		// *r2 = *r
		// r2.Header.Set("Content-Type", "application/json")
		// r2.Header.Set("X-User-ID", "sampleID")
		// r2.Header.Set("X-Application-ID", "1")
		// handler.ServeHTTP(w, r2)
	}), nil
}
