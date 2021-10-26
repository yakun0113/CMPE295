# CMPE295

To run this app, 

1. Since the number of searches an api key can have is limited (100 searches/ month), please use your own SerpApi API key to scrape data:
    - step1: go to SerpApi website: https://serpapi.com
    - step2: create a new account and get a free API key
    - step3: go to backend/scraper.py and replace the API key with your API key.
2. Replace Google maps API key to your own key. 
    - https://developers.google.com/maps 
3. Run two commands before going to the next step. The two commands are in the following page.
    - https://medium.com/rungo/secure-https-servers-in-go-a783008b36da
   
  ``` 
  openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr
  ```
  ```
  openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt
  ```
  Replace the lochost.crt, localhost.key to the ones generated by the above commands.
  
4. Go to frontend file and execute two commands:
  - npm install --save @fawmi/vue-google-maps
  - npm run build
  
5. In terminal, go to the backend/go file and run the command "go run *.go" to start the server.
6. In terminal, go to the backend/python file and run the command "python3 scraping_server.py" to start the scraping server.
