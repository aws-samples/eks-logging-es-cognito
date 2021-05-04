package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

 
// Cognito API calls. You should store this information 
// in other way, because this way it would be possible 
// to extract those strings from your compiled app if it 
// leaks somewhere.
const clientId = ""
const clientSecret = ""
// This is the username and password of a user from your
// Cognito user pool.
const username = ""
const password = ""

func main() {
    conf := &aws.Config{Region: aws.String("us-east-1")}
    sess := session.Must(session.NewSession(conf))

    // This is the part where we generate the hash.
    mac := hmac.New(sha256.New, []byte(clientSecret))
    mac.Write([]byte(username + clientId))

    secretHash := base64.StdEncoding.EncodeToString(mac.Sum(nil))

    cognitoClient := cognitoidentityprovider.New(sess)

    authTry := &cognitoidentityprovider.InitiateAuthInput{
        AuthFlow: aws.String("USER_PASSWORD_AUTH"),
        AuthParameters: map[string]*string{
            "USERNAME": aws.String(username),
            "PASSWORD": aws.String(password),
            "SECRET_HASH": aws.String(secretHash),
        },
        ClientId: aws.String(clientId),
    }

    res, err := cognitoClient.InitiateAuth(authTry)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("authenticated")
        fmt.Println(res.AuthenticationResult)
    }
}