# GoXrm
This library has been written to reproduce Microsoft.Xrm.Sdk functions.
The initial idea is to provide the same "objects" and "methods" as there is in Microsoft SDK. So that .net developpers do not arrive in a totally new world.
It connects to Dynamics 365 online using the Web API, and gets identity against Azure AD. It shouldn't work with onpremise deployments.

<b>Getting started</b>
First of all, you need to create a CrmServiceClient. You need to know the "Tenant Id" of your organization and the Application Id that you got when registering the application.
Import "github.com/PierreVicente/GoXrm/Client"
then use one of the NewCrmServiceClient functions
The possible parameters are
- Login Url: "https://login.microsoftonline.com" or "https://login.windows.net" work 
- Tenant Id:  the id of your azure AD organization (it's a guid)
- Client Id: the id of your registered application (it's a guid too)
- Resource Url: the url of your Dynamics organization: https://XXXXX.crmX.dynamics.com
- User id: use it if you want to connect as an organization user: xxxuserxxx@yyyorgyyy.onmicrosoft.com
- Password: no need to explain
- Secret: Means that your are connecting with an application user: this is the secret string provided by Azure AD registration

