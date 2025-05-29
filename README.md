# SimpleServer 🚀

![Go](https://img.shields.io/badge/language-Go-blue.svg)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Build with Go](https://img.shields.io/badge/build-passing-brightgreen.svg)
![Minimal size](https://img.shields.io/badge/minimal-21KB-lightgrey.svg)
![Open Source](https://img.shields.io/badge/open--source-yes-green)
![Say Hello!](https://img.shields.io/badge/says-hello!-orange?logo=handshake)


A simple HTTP server written in Go.   
Perfect for beginners who want to learn the basics of web servers and Go programming!   
**Let’s make your browser say “Hello” like never before!** 👋

---

## 🧠 Features
- Serves HTTP requests
- Easy to understand 
- Minimal dependencies

---

## 🛠️ Requirements

- [Go 1.18+](https://go.dev/dl/) installed on your system
- Internet connection (for cloning the repository)
- (Optional) [Postman](https://www.postman.com/) or `curl` for testing POST requests

---

## 🎬 Getting Started
*Don’t worry, even your grandma could run this server. Just follow these steps!*

1. **Clone this repository:**
   ```bash
   git clone https://github.com/hrosicka/SimpleServer.git
   ```
   ```bash
   cd SimpleServer
   ```
   
2. **Run the server:**
   ```bash
   go run main.go handlers.go types.go
   ```

3. **Access the server:**
   Open your browser and visit [http://localhost:8080](http://localhost:8080)

---

## 🧪 How to Test the Application 

### 1. Start the Server

Make sure you have Go installed. Then, run the server from the project directory.

---

### 2. Testing GET Endpoints in the Browser

Server provides two GET endpoints that can be tested directly from the browser by entering the URL. Both endpoints accept query parameters, so you can easily customize the request.

#### `/hello` endpoint (GET):
**Returns a personalized JSON greeting message. Optionally takes a name parameter via the query string.**   
*Try changing the name for a different greeting. The server never gets tired of saying hello!*
- Open your browser.
- In the address bar, enter (replace ```Hanka``` with any name you like):

  ```
  http://localhost:8080/hello?name=Hanka
  ```

- **Response:** You’ll see a JSON message like:
  ```json
  {
  "message": "Hello Hanka!"
  }
  ```
  
#### `/data` endpoint (GET):
**Returns a JSON message and echoes back the provided id and type parameters.**   
*Send your data and see it come right back at you! It’s like playing catch with a server.*
- Open your browser.
- In the address bar, enter (replace ```123456``` and ```user``` with any name you like):

  ```
  http://localhost:8080/data?id=123456&type=user
  ```

- **Response:** You’ll see a JSON message like:
  ```json
  {
     "message": "Data retrieved.",
     "data": {
       "id": "123456",
       "type": "user"
     }
  }
  ```
  
---

### 2. Testing POST Endpoints in the Browser
#### `/submit` endpoint (POST):
**Accepts a JSON object with name, surname, and email in the request body. Returns a confirmation message and saves the data.**   
*Submit your name, surname, and email. The server will remember… at least until you delete the file!*
**With Postman:**

1. Open Postman.
2. Set the method to POST.
3. Enter the URL:
```http://localhost:8080/submit```
4. Go to the Body tab, select raw, and choose JSON.
5. Enter sample data:
```json
{
    "name": "Martin",
    "surname": "West",
    "email": "martin.west@example.com"
}
```
6. Click **Send**.
7. You should see a JSON response like:
```json
{
    "message": "Data received for user Martin West (martin.west@example.com).",
    "data": {
        "name": "Martin",
        "surname": "West",
        "email": "martin.west@example.com"
    }
}
```

**With curl:**
```
curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice", "surname": "Blue", "email": "alice.blue@example.com"}' http://localhost:8080/submit
```

