# SimpleServer

A simple HTTP server written in Go. 🚀  
Perfect for beginners who want to learn the basics of web servers and Go programming!

---

## Features
- Serves HTTP requests
- Easy to understand
- Minimal dependencies

---

## Getting Started

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

## How to Test the Application

### 1. Start the Server

Make sure you have Go installed. Then, run the server from the project directory.

### 2. Testing GET Endpoints in the Browser

Server provides two GET endpoints that can be tested directly from the browser by entering the URL. Both endpoints accept query parameters, so you can easily customize the request.

#### 1. `/hello` endpoint
Returns a JSON greeting message. Optionally takes a name parameter via the query string.
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
  
#### 2. `/data` endpoint
Returns a JSON message and echoes back the provided id and type parameters.
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




