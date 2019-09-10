---
title: "LDAP Authentication"
---

Enterprise users who host their own private Convox Console can take advantage of LDAP authentication and access.  For more information on the Enterprise plan, see [here](https://www.convox.com/enterprise).

---

SSO access through LDAP can be enabled by simply setting three or four environment variables on your Console app.

```
$ convox env set -a console AUTHENTICATION=ldap
$ convox env set -a console LDAP_ADDR=auth.example.org:636
$ convox env set -a console LDAP_BIND=uid=%s,dc=example,dc=org
```

* `AUTHENTICATION` set to `ldap`
* `LDAP_ADDR` set to your LDAP server address.  
* `LDAP_BIND` to a full bind string where `%s` will be substituted for the user's email address.

If your LDAP server does not have a valid certificate issued by a known CA, you can disable certificate validation:

```
$ convox env set -a console LDAP_VERIFY=no
```

Once configured, promote the environment changes

```
$ convox releases promote -a console --wait
```

---

To disable LDAP SSO access, change the `AUTHENTICATION` environment variable back to it's previous value or simply remove the value, as appropriate.

