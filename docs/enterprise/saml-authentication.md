---
title: "SAML Authentication"
description: "Configure SAML single sign-on authentication for a self-hosted Convox Console on the Enterprise plan."
---

# SAML Authentication

Enterprise users who host their own private Convox Console can take advantage of SAML SSO authentication and access. For more information on the Enterprise plan, see [here](https://www.convox.com/enterprise).

SSO access through SAML can be enabled by setting two environment variables on your Console app.

```bash
$ convox env set -a console AUTHENTICATION=saml
$ convox env set -a console SAML_METADATA=https://login.microsoftonline.com/common/FederationMetadata/2007-06/FederationMetadata.xml
```

* `AUTHENTICATION` set to `saml`
* `SAML_METADATA` set to the metadata endpoint for your SAML Identity Provider.  This varies from provider to provider, so check your provider's documentation.
  * For example, Microsoft/Azure IdP information can be found [here](https://docs.microsoft.com/en-us/azure/active-directory/develop/azure-ad-federation-metadata)
  * For Okta, the IdP information can be found [here](https://support.okta.com/help/s/article/How-do-we-download-the-IDP-XML-metadata-file-from-a-SAML-Template-App)


Once configured, promote the environment changes

```bash
$ convox releases promote -a console --wait
```

The Console will retrieve the metadata from the provider endpoint to configure everything else.

## Disabling SAML

To disable SAML SSO access, change the `AUTHENTICATION` environment variable back to its previous value or remove the value, as appropriate.

## See Also

- [LDAP Authentication](/enterprise/ldap-authentication)
- [Console Access Control](/console/access-control)
- [Environment Variables](/application/environment)
- [Releases](/deployment/releases)
