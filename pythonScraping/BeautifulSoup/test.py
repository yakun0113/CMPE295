import requests
import json
from bs4 import BeautifulSoup

headers = {"referer":"https://www.target.com",
            "User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_5_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"}
session = requests.Session()
firstUrl = "https://www.target.com"
url = "https://www.target.com/s?searchTerm=toilet+paper"
session.get(firstUrl, headers = headers)
page = session.get(url, headers = headers)
soup = BeautifulSoup(page.content, 'html.parser')
    
productInfo = soup.select('ul[class="Row-uds8za-0 jBYETz h-padding-t-tight"] li[data-test="list-entry-product-card"]')
    
for result in productInfo:
      
    #productName = result.select_one('div div[class="h-display-flex"] a').text
    #productPrice = result.select_one('span[class="price-group"]').text
    #productRating = result.select_one('span[class="visuallyhidden seo-avg-rating"]').text
    #productStatus = result.select_one('div[class="free-pickup-today"]')
    #productLink = result.select_one('div[class="search-result-product-title gridview"] a')["href"]
    print("productName")
