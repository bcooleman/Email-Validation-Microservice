# Email Validation Microservice - RapidAPI Usage Guide

This Email Validation Microservice is available on **RapidAPI** to help validate and classify email addresses. It analyzes syntax, domain risks, and unusual characters to determine if an email is benign, spam, or malicious. Follow this guide to integrate and use the API via RapidAPI.

---

## Features

- **Syntax Validation**: Checks if the email address has a valid format (e.g., `a@b.com`).
- **Risky Domain Detection**: Flags suspicious domains commonly used for spam or phishing (e.g., `.ru`, `.xyz`).
- **Character Analysis**: Identifies unusual or Cyrillic characters to detect potential malicious intent.
- **Categorization**: Returns detailed flags classifying the email as:
  - **Benign**: No issues detected.
  - **Spam**: Suspicious domain or pattern.
  - **Likely Malicious**: High-risk content or unusual characters.

---

## Accessing the API

You can access the Email Validation Microservice through the RapidAPI platform. Here’s how:

### 1. **Subscribe to the API**
- Visit the RapidAPI page for the Email Validation Microservice.
- Select a subscription plan that suits your needs (e.g., Free, Pay-as-you-go).

### 2. **API Endpoint**
The API is accessible at:
```
POST https://<api-endpoint-url>/classify
```

---

## How to Use

### 1. **Headers**
You must include the following headers in your requests:
- `Content-Type: application/json`
- `X-RapidAPI-Key`: Your RapidAPI key, available in your RapidAPI dashboard.
- `X-RapidAPI-Host`: The RapidAPI host for this microservice (e.g., `email-validation-microservice.p.rapidapi.com`).

### 2. **Request Body**
The request body must contain the email address to validate, formatted as JSON:
```json
{
  "email": "test@example.com"
}
```

---

## Example Usage

### Using `curl`
```bash
curl -X POST https://<api-endpoint-url>/classify \
     -H "Content-Type: application/json" \
     -H "X-RapidAPI-Key: YOUR_RAPIDAPI_KEY" \
     -H "X-RapidAPI-Host: email-validation-microservice.p.rapidapi.com" \
     -d '{"email": "test@example.com"}'
```

### Using Postman
1. Open Postman and create a new `POST` request.
2. Set the URL to `https://<api-endpoint-url>/classify`.
3. Add the required headers:
   - `Content-Type: application/json`
   - `X-RapidAPI-Key: YOUR_RAPIDAPI_KEY`
   - `X-RapidAPI-Host: email-validation-microservice.p.rapidapi.com`
4. In the request body, provide the email address in JSON format:
   ```json
   {
     "email": "test@example.com"
   }
   ```
5. Send the request and inspect the response.

---

## Sample Responses

### 1. **Benign Email**
For a valid and benign email:
```json
{
  "email": "test@example.com",
  "flags": [
    {
      "category": "benign",
      "description": "No issues detected."
    }
  ]
}
```

### 2. **Spam Email**
For an email using a suspicious domain:
```json
{
  "email": "user@dangerous.xyz",
  "flags": [
    {
      "category": "spam",
      "description": "Email uses a risky domain: .xyz"
    }
  ]
}
```

### 3. **Likely Malicious Email**
For an email containing Cyrillic characters:
```json
{
  "email": "спам@пример.ru",
  "flags": [
    {
      "category": "likely malicious",
      "description": "Email contains Cyrillic or unusual characters."
    },
    {
      "category": "spam",
      "description": "Email uses a risky domain: .ru"
    }
  ]
}
```

---

## Error Handling

If something goes wrong, the API returns an appropriate error message.

### Example Error Responses:
- **Missing or invalid headers**:
  ```json
  {
    "error": "Invalid or missing X-RapidAPI-Key header"
  }
  ```
- **Invalid request body**:
  ```json
  {
    "error": "Invalid JSON payload"
  }
  ```

---

## RapidAPI Dashboard

Track your API usage, request logs, and quota limits directly in the [RapidAPI Dashboard](https://rapidapi.com/).

---

## Support

If you encounter any issues or need assistance, please reach out via RapidAPI’s support system or submit an issue on the project’s GitHub page.

