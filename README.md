# CMPE295

To run this app, 
1. Please download chrome driver and set the executable path to your local path.
2. Please set the path to python your own path.
3. Since the number of searches an api key can have is limited (100 searches/ month), please use your own SerpApi API key to scrape data:
    - step1: go to SerpApi website: https://serpapi.com
    - step2: create a new account and get a free API key
    - step3: go to backend/scraper.py and replace the API key with your API key.
4. Replace Google maps API key to your own key. 
    - https://developers.google.com/maps 
5. Run two commands before going to the next step. The two commands are in the following page.
    - https://medium.com/rungo/secure-https-servers-in-go-a783008b36da
   
  ``` 
  openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
  ```
  ```
  openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt
  ```
6. In terminal, go to the backend file and run the command " go run server.go" to start the app.
