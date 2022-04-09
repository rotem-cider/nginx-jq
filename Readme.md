# Nginx JQ

Takes an nginx file and turns it into json

### For example to search for any proxy header of ldap

```
 go get -v github.com/rotem-cider/nginx-jq@21b292fc
 wget https://raw.githubusercontent.com/rotem-cider/nginx-jq/main/tests/nginx.conf
./nginx-jq nginx.conf | jq -r ".. | (select(.Name==\"proxy_set_header\"))? | select(.Parameters[0] | startswith(\"X-Ldap\")) | [.Parameters[0], .Parameters[1]] | @csv"
```