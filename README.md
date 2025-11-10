# URL Shortener API

A simple URL shortener API built with Golang, Gorilla Mux, and Docker. This API allows users to shorten long URLs and redirect using generated short codes.

##  Features
- Shorten long URLs
- Redirect using short URLs
- In-memory storage for simplicity
- No analytics


##  Installation & Running

```
docker build . -t shorturl
docker run -P shorturl
```

The server will start on **http://localhost:8080**.

##  API Endpoints

### Shorten a URL
**Endpoint:** `POST /shorten`

**Request:**
```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
```

**Response:**
```json
{
  "short_url": "Pwl3OV"
}
```

### Redirect to Original URL
**Endpoint:** `GET /{short_code}`

**Example:**
```bash
curl -L http://localhost:8080/Pwl3OV
```
Redirects to `https://example.com`.
