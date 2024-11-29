# Email Validation API

The **Email Validation API** is a powerful microservice available on RapidAPI, designed to validate email addresses for validity, spam risk, and potential maliciousness. RapidAPI handles API key management, making integration seamless.

---

## Features

- **Syntax Validation**: Ensures the email address conforms to standard syntax rules.
- **Domain Validation**: Verifies the existence of DNS MX records for the domain.
- **Disposable Email Detection**: Detects domains commonly used for disposable or temporary emails (e.g., `mailinator.com`).
- **Risky Domain Detection**: Flags email domains with spammy or risky TLDs such as `.ru`, `.xyz`, etc.
- **Blacklist Validation**: Compares the email's domain against a predefined blacklist of known spam or malicious domains.
- **Cyrillic Character Detection**: Identifies email addresses containing Cyrillic or other unusual characters.
- **SMTP Validation**: Optionally validates the email address by connecting to its SMTP server.

---

## How to Use

This API is available on **RapidAPI**. Follow these steps to get started:

### 1. Sign Up or Log In to RapidAPI

Visit [RapidAPI](https://rapidapi.com/) and sign up or log in to your account.

### 2. Subscribe to the API

- Search for **Email Validation API** on RapidAPI.
- Choose a subscription plan that suits your needs (Free or Paid).

### 3. Make Requests

Use the `/validate` endpoint to validate email addresses. Details below.

---

### **Endpoint: `/validate`**

- **Base URL**: `https://email-classifier2.p.rapidapi.com/`
- **Method**: `POST`
- **Headers**:
  - `Content-Type`: `application/json`
  - `X-RapidAPI-Key`: Automatically added by RapidAPI.
- **Request Body**:
  ```json
  {
    "email": "example@test.com"
  }
  ```

---

### **Example Response**

```json
{
  "email": "example@test.com",
  "syntax_valid": true,
  "domain_valid": true,
  "disposable_domain": false,
  "risky_domain": false,
  "on_blacklist": false,
  "contains_cyrillic": false,
  "smtp_valid": true
}
```

#### **Response Fields**

| Field              | Type    | Description                                                |
|--------------------|---------|------------------------------------------------------------|
| `email`            | string  | The input email address being validated.                  |
| `syntax_valid`     | boolean | Indicates if the email syntax is valid.                   |
| `domain_valid`     | boolean | Checks if the email domain has valid MX records.          |
| `disposable_domain`| boolean | Detects if the email belongs to a disposable provider.     |
| `risky_domain`     | boolean | Flags if the email domain has a risky or spammy TLD.      |
| `on_blacklist`     | boolean | Indicates if the email domain is on a blacklist.          |
| `contains_cyrillic`| boolean | Detects if the email contains Cyrillic characters.         |
| `smtp_valid`       | boolean | Optionally checks if the email is valid via SMTP.         |

---

## Testing the API with Postman

1. Open Postman and create a new request.
2. Set the following details:
   - **Method**: `POST`
   - **URL**: `https://email-classifier2.p.rapidapi.com/validate`
   - **Headers**:
     - `Content-Type`: `application/json`
     - `X-RapidAPI-Key`: Automatically added by RapidAPI when using your account.
   - **Body**:
     ```json
     {
       "email": "example@test.com"
     }
     ```
3. Click **Send** and view the response.

---

## Error Responses

- **400 Bad Request**: The request payload is invalid or the `email` field is missing.
- **500 Internal Server Error**: An unexpected error occurred while processing the request.

---

## Common Use Cases

- **Sign-Up Form Validation**: Ensure that users provide valid, non-disposable email addresses during sign-ups.
- **Spam Prevention**: Block users from registering with spammy or risky email addresses.
- **Email List Cleansing**: Clean and validate large email lists for marketing campaigns or user databases.

---

## Support

For assistance or to report issues, contact us through the **RapidAPI Support Portal**.

---

## License

This API is proprietary and available exclusively via **RapidAPI**. Unauthorized usage is prohibited.

