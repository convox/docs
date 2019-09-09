---
title: "SAML Authentication"
---

Enterprise users who host their own private Convox Console can take advantage of SAML SSO authentication and access.  For more information on the Enterprise plan, see [here](https://www.convox.com/enterprise).

---

SSO access through SAML can be enabled by simply setting two environment variables on your Console app.

* `AUTHENTICATION` set to `saml`
* `SAML_METADATA` set to the metadata endpoint for your SAML Identity Provider.  This varies from provider to provider so please check your documentation from them.
  * For example, Microsoft/Azure IdP information can be found [here](https://docs.microsoft.com/en-us/azure/active-directory/develop/azure-ad-federation-metadata)
  * For Okta, the IdP information can be found [here](https://support.okta.com/help/s/article/How-do-we-download-the-IDP-XML-metadata-file-from-a-SAML-Template-App)

The Console will retrieve the metadata from the provider endpoint to configure everything else.

---

To disable SAML SSO access, change the `AUTHENTICATION` environment variable back to it's previous value or simply remove the value, as appropriate.

