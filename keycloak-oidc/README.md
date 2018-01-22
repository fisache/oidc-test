<pre>
Referred to https://jacobmartins.com/2016/02/29/getting-started-with-oauth2-in-go/
</pre>

```
Install Keycloak and run
Create Realm, Client and change Access Type to confidential
Add 'http://localhost:3000/*' in Valid Redirect URIs

$ export keycloakUrl='http://localhost:8080'
$ export realm='<your realm>'
$ export clientId='<your client id>'
$ export clientSecret='<your client secret>'

$ cd oidc-test/keycloak-oidc
$ go get golang.org/x/oauth2
$ go get github.com/dgrijalva/jwt-go
$ go build
$ ./keycloak-oidc

Connect to localhost:3000
```
