import requests
import json
from bs4 import BeautifulSoup

def scrape_walmart(product_name, session, headers):

    firstUrl = "https://www.walmart.com"
    url = "https://www.walmart.com/search/?query="+product_name
    session.get(firstUrl, headers = headers)
    page = session.get(url, headers = headers)
    soup = BeautifulSoup(page.content, 'html.parser')
    

    productInfo = soup.select('div[class="search-product-result"] ul li div div[class="search-result-gridview-item clearfix arrange-fill"]')
    
    for result in productInfo:
      
        productName = result.select_one('div[class="search-result-product-title gridview"] a').text
        productPrice = result.select_one('span[class="price-group"]').text
        productRating = result.select_one('span[class="visuallyhidden seo-avg-rating"]').text
        productStatus = result.select_one('div[class="free-pickup-today"]')
        productLink = result.select_one('div[class="search-result-product-title gridview"] a')["href"]
        productLink = firstUrl + productLink

        if productStatus is not None:
            productStatus = "In stock"
                  
        else:
            productStatus = "Out of stock"
        
        product = {
                "name": productName,
                "price": productPrice,      
                "rating": productRating,
                "status": productStatus,
                "link": productLink
                  } 
        product_list["Walmart"].append(product)    

    for product in product_list["Walmart"]:
        if product["status"] == "In stock":
            locationLink = product["link"]
            page = session.get(locationLink, headers = headers)
            soup = BeautifulSoup(page.content, 'html.parser')
            product_Info = soup.select('div[class="hf-Bot"]')
            for res in product_Info:
                productLocation = res.select_one('div[class="fulfillment-text"]').p.text
                productLocation = productLocation.split("In stock at ")[1]
                product_list["Walmart at " + productLocation] = product_list.pop("Walmart")
            break

def scrape_target(product_name, headers):
    firstUrl = "https://www.target.com"
    url = "https://www.target.com/s?searchTerm="+product_name

def scrape_safeway(product_name, headers):
    firstUrl = "https://www.safeway.com"
    safeway_url = "https://www.safeway.com/shop/search-results.html?q="+product_name

def scrape_kroger(product_name, headers):
    firstUrl = "https://www.kroger.com"
    kroger_url = "https://www.kroger.com/search?query="+product_name


if __name__ == '__main__':

    headers = {"referer":"https://www.google.com",
               "User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_5_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"}
    session = requests.Session()

    # read product name
    with open('product_name.txt') as f:
        product_name = f.read()

    product_list = {"Walmart":[], "Target":[], "SafeWay":[], "Kroger":[]}

    scrape_walmart(product_name, session, headers)

    # scrape_target(product_name, session, headers)

    # scrape_safeway(product_name, session, headers)

    # scrape_kroger(product_name, session, headers)

    with open('productsInfo.json', 'w') as outfile:
        json.dump(product_list, outfile, indent=4)
