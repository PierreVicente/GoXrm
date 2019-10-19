package Client

import (
	"github.com/PierreVicente/GoXrm/AADAuth"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type CrmServiceClient struct {
	CallerId      uuid.UUID
	ApiVersion    string
	aadAuthResult AADAuth.AADAuthResult
	loginUrl      string
	tenantId      string
	clientId      string
	resourceUrl   string
	userName      string
	password      string
	secret        string
}

const defaultApiVersion = "9.1"

func (this *CrmServiceClient) SetAadAuthResult(AadAuthResult AADAuth.AADAuthResult) {
	this.aadAuthResult = AadAuthResult
}

func (this *CrmServiceClient) GetAadAuthResult() AADAuth.AADAuthResult {

	if this.aadAuthResult.AuthType == AADAuth.AuthType_Office365 {
		iexp, _ := strconv.ParseInt(this.aadAuthResult.ExpiresOn, 10, 64)
		exp := time.Unix(iexp, 0)
		diff := exp.Sub(time.Now().UTC())
		if diff.Minutes() < 2 {
			refrsh := AADAuth.AADAuthResult{}

			if this.aadAuthResult.RefreshToken != "" {
				refrsh = AADAuth.RenewToken(this.loginUrl, this.tenantId, this.clientId, this.resourceUrl, this.aadAuthResult.RefreshToken)
			}

			if refrsh.RefreshToken == "" {
				switch this.aadAuthResult.AuthMode {
				case AADAuth.AuthMode_Password:
					refrsh = AADAuth.AADAuthenticateWithPassword(this.loginUrl, this.tenantId, this.clientId,
						this.resourceUrl, this.userName, this.password)
					break
				case AADAuth.AuthMode_Secret:
					refrsh = AADAuth.AADAuthenticateWithSecret(this.loginUrl, this.tenantId, this.clientId,
						this.resourceUrl, this.secret)
					break
				}
			}

			this.aadAuthResult.AccessToken = refrsh.AccessToken
			this.aadAuthResult.ExpiresIn = refrsh.ExpiresIn
			this.aadAuthResult.ExpiresOn = refrsh.ExpiresOn
			this.aadAuthResult.ExtExpiresIn = refrsh.ExtExpiresIn
			this.aadAuthResult.NotBefore = refrsh.NotBefore
			this.aadAuthResult.RefreshToken = refrsh.RefreshToken
			this.aadAuthResult.Resource = refrsh.Resource
			this.aadAuthResult.Scope = refrsh.Scope
			this.aadAuthResult.TokenType = refrsh.TokenType

			return this.aadAuthResult
		}
		return this.aadAuthResult
	}
	panic("not supported auth")
}

func NewCrmServiceClientAadAuth(authObject AADAuth.AADAuthResult) CrmServiceClient {
	srv := CrmServiceClient{}
	srv.SetAadAuthResult(authObject)
	srv.ApiVersion = defaultApiVersion
	return srv
}

func NewCrmServiceClientPassword(_loginUrl string, _tenantId string, _clientId string, _resourceUrl string, _userName string, _password string) CrmServiceClient {
	srv := CrmServiceClient{
		loginUrl:      _loginUrl,
		tenantId:      _tenantId,
		clientId:      _clientId,
		resourceUrl:   _resourceUrl,
		userName:      _userName,
		password:      _password,
		ApiVersion:    defaultApiVersion,
		aadAuthResult: AADAuth.AADAuthenticateWithPassword(_loginUrl, _tenantId, _clientId, _resourceUrl, _userName, _password),
	}
	return srv
}

func NewCrmServiceClientSecret(_loginUrl string, _tenantId string, _clientId string, _resourceUrl string, _secret string) CrmServiceClient {
	srv := CrmServiceClient{
		loginUrl:      _loginUrl,
		tenantId:      _tenantId,
		clientId:      _clientId,
		resourceUrl:   _resourceUrl,
		secret:        _secret,
		ApiVersion:    defaultApiVersion,
		aadAuthResult: AADAuth.AADAuthenticateWithSecret(_loginUrl, _tenantId, _clientId, _resourceUrl, _secret),
	}
	return srv
}

func NewCrmServiceClientRefreshToken(_loginUrl string, _resourceUrl string, _tenantId string, _clientId string, _refreshToken string) CrmServiceClient {
	srv := CrmServiceClient{}
	srv.loginUrl = _loginUrl
	srv.tenantId = _tenantId
	srv.clientId = _clientId
	srv.resourceUrl = _resourceUrl
	srv.ApiVersion = defaultApiVersion
	srv.SetAadAuthResult(AADAuth.RenewToken(_loginUrl, _tenantId, _clientId, _resourceUrl, _refreshToken))
	return srv
}
