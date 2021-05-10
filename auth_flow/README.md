# Steps to attach MFA to User in Amazon Cognito

- Add MFA to UserPool using this [link](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa.html)

- Get the Access Token with the Go script in this folder.

- Run the AssociateSoftwareToken Command

```shell
aws cognito-idp associate-software-token --access-token ACCESS_TOKEN
```

Output will look like something like this

```json
{
    "SecretCode":
    "AETQ6XXMDFYMEPFQQ7FD4HKXXXXAOY3MBXIVRBLRXX3SXLSHHWOA"
}
```

- Open Google Authenticator or any other Authenticator and click in **Get started**

- Setup the **SecretCode** that you get with the above command.

- Choose the Type of key dropdown list, and then select Time based.

- For Account name, enter an account name. For example, BobPhone.

- Verify the software token using the time-based password that appears on the screen and the following code

```shell
aws cognito-idp verify-software-token --access-token ACCESS_TOKEN --user-code AUTHETICATOR_CODE --friendly-device-name BobPhone
```

Output will look like the following

```json
{
    "Status": "SUCCESS"
}
```

- Configure the user's MFA configuration to TOTP MFA

```shell
aws cognito-idp admin-set-user-mfa-preference --software-token-mfa-settings Enabled=true,PreferredMfa=true --username Bob --user-pool-id us-east-1_123456789
```
