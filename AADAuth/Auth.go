package AADAuth

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func AADAuthenticateWithSecret(loginUrl string, tenantId uuid.UUID, clientId uuid.UUID, resourceUrl string, secret string) AADAuthResult {

	tokenEndpoint := loginUrl
	if !strings.HasSuffix(tokenEndpoint, "/") {
		tokenEndpoint += "/"
	}

	tokenEndpoint += tenantId.String() + "/oauth2/token"

	var body = "resource=" + url.QueryEscape(resourceUrl) + "&"
	body += "client_id=" + clientId.String() + "&"
	body += "grant_type=client_credentials" + "&"
	body += "client_secret=" + url.QueryEscape(secret)

	var stringContent = []byte(body)

	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBuffer(stringContent))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-TYpe", "application/x-www-form-urlencoded")

	client := &http.Client{}
	var resp, err2 = client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	rbody, _ := ioutil.ReadAll(resp.Body)

	aada := AADAuthResult{}
	err3 := json.Unmarshal(rbody, &aada)
	if err3 != nil {
		panic(err3)
	}

	aada.AuthMode = AuthMode_Secret
	aada.AuthType = AuthType_Office365

	return aada

}

func AADAuthenticateWithPassword(loginUrl string, tenantId uuid.UUID, clientId uuid.UUID, resourceUrl, userName, password string) AADAuthResult {

	var tokenEndpoint = loginUrl
	if !strings.HasSuffix(tokenEndpoint, "/") {
		tokenEndpoint += "/"
	}
	tokenEndpoint += tenantId.String() + "/oauth2/token"
	//https://login.microsoftonline.com/065377e9-ac15-4669-9c10-c5e2b3848e99/oauth2/token";
	var body = "resource=" + url.QueryEscape(resourceUrl) + "&"
	body += "client_id=" + clientId.String() + "&"
	body += "grant_type=password" + "&"
	body += "username=" + url.QueryEscape(userName) + "&"
	body += "password=" + url.QueryEscape(password)

	//var stringContent = new StringContent(body, Encoding.UTF8, "application/x-www-form-urlencoded");
	var stringContent = []byte(body)

	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBuffer(stringContent))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-TYpe", "application/x-www-form-urlencoded")

	client := &http.Client{}
	var resp, err2 = client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	rbody, _ := ioutil.ReadAll(resp.Body)

	aada := AADAuthResult{}
	err3 := json.Unmarshal(rbody, &aada)
	if err3 != nil {
		panic(err3)
	}

	aada.AuthMode = AuthMode_Password
	aada.AuthType = AuthType_Office365

	return aada

}

func RenewToken(loginUrl string, tenantId uuid.UUID, clientId uuid.UUID, resourceUrl, rToken string) AADAuthResult {

	var tokenEndpoint = loginUrl
	if !strings.HasSuffix(tokenEndpoint, "/") {
		tokenEndpoint += "/"
	}
	tokenEndpoint += tenantId.String() + "/oauth2/token"
	//https://login.microsoftonline.com/065377e9-ac15-4669-9c10-c5e2b3848e99/oauth2/token";
	var body = "resource=" + url.QueryEscape(resourceUrl) + "&"
	body += "client_id=" + clientId.String() + "&"
	body += "grant_type=refresh_token" + "&"
	body += "refresh_token=" + rToken

	var stringContent = []byte(body)

	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBuffer(stringContent))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-TYpe", "application/x-www-form-urlencoded")

	client := &http.Client{}
	var resp, err2 = client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	rbody, _ := ioutil.ReadAll(resp.Body)

	var aada = AADAuthResult{}
	err3 := json.Unmarshal(rbody, &aada)
	if err3 != nil {
		panic(err3)
	}

	return aada
}
