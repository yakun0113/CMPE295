import requests
import json
from bs4 import BeautifulSoup
from selenium import webdriver
import time
from threading import Thread
from multiprocessing import Process, Manager

def get_walmart(product_name, option, latitude, longitude):
    
    Walmart = requests.get("https://serpapi.com/search.json?engine=google_maps&q=walmart&ll=%40"+latitude+"%2C"+longitude+"%2C15z&type=search&api_key=" + serpApi_key)
    Walmart = Walmart.json()
    Walmart_url = Walmart["local_results"][0]["website"]
    Walmart_url = Walmart_url.split("?")[0] + "/search?query=" + product_name
    Walmart_location = Walmart["local_results"][0]["gps_coordinates"]

    driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/pythonScraping/chromedriver", options = option)
    driver.get(Walmart_url)
    button = driver.find_element_by_xpath('//*[@id="content"]/div[2]/section[2]/div[2]/div/div/div/div/div/div[2]/form/div/div[3]/button')
    button.click()
    for i in range(10):
        driver.execute_script("window.scrollBy(0, document.body.scrollHeight/10)")
        time.sleep(0.1)

    walmart_soup = BeautifulSoup(driver.page_source, 'html.parser')
    driver.quit()
    scrape_walmart(walmart_soup, Walmart_location)

def get_target(product_name, option, latitude, longitude):

    Target = requests.get("https://serpapi.com/search.json?engine=google_maps&q=target&ll=%40"+latitude+"%2C"+longitude+"%2C15z&type=search&api_key=" + serpApi_key)
    Target = Target.json()
    url = Target["local_results"][0]["place_id_search"]
    page = requests.get(url+"&api_key="+serpApi_key)
    page = page.json()
    Target_url=page["place_results"]["website"]
    Target_location = Target["local_results"][0]["gps_coordinates"]

    driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/pythonScraping/chromedriver", options = option)
    driver.get(Target_url)
    driver.execute_script("window.scrollBy(0, document.body.scrollHeight/5)")
    time.sleep(0.5)
    button = driver.find_element_by_xpath("//*[@id='__next']/div[3]/div/div[1]/div[2]/div/div[5]/button[2]")
    button.click()
    driver.get("https://www.target.com/s?searchTerm="+product_name+"&facetedValue=5zkty")

    for i in range(15):
        driver.execute_script("window.scrollBy(0, document.body.scrollHeight/15)")
        time.sleep(0.1)
    target_soup = BeautifulSoup(driver.page_source, 'html.parser')
    driver.quit()
    scrape_target(target_soup, Target_location)

def get_walgreens(product_name, option, latitude, longitude):
  
    Walgreens = requests.get("https://serpapi.com/search.json?engine=google_maps&q=walgreens&ll=%40"+latitude+"%2C"+longitude+"%2C15z&type=search&api_key=" + serpApi_key)
    Walgreens = Walgreens.json()
    Walgreens_url = Walgreens["local_results"][0]["website"]
    Walgreens_location = Walgreens["local_results"][0]["gps_coordinates"]
    
    driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/pythonScraping/chromedriver", options = option)
    driver.get(Walgreens_url)
    search = driver.find_element_by_xpath('//*[@id="shop-placeholder"]')
    search.send_keys(product_name)
    button = driver.find_element_by_xpath('//*[@id="storeServiceDiv"]/div[1]/div[1]/div[3]/div/div/button')
    button.click()

    for i in range(15):
        driver.execute_script("window.scrollBy(0, document.body.scrollHeight/15)")
        time.sleep(0.1)

    walgreens_soup = BeautifulSoup(driver.page_source, 'html.parser')
    driver.quit()
    scrape_walgreens(walgreens_soup, Walgreens_location)

def scrape_walmart(soup, Walmart_location):
    p_list = []
    productId = 0
    productInfo = soup.select('ul[class="results-list grid-view"] div[class="Grid-col u-size-6-12 u-size-4-12-s u-size-3-12-m u-size-3-12-l"]')
    for result in productInfo:
        productImage = result.select_one('img[class="tile-img"]')["src"]    
        productName = result.select_one('div[class="tile-title"]').text
        productPrice = result.select_one('span[class="price-group"]')
        productRating = result.select_one('span[class="stars-container"]')
        productStatus = result.select_one('div[class="tile-in-store-info"]')
        productLink = result.select_one('div[class="tile u-focusTile"] a[class="display-block"]')['href']

        if productPrice == None:
            productPrice = "None"
        else:
            productPrice = productPrice.text
        if productRating == None:
            productRating = "None"
        else:
            productRating = productRating['aria-label']
        productImage = (productImage.split("//"))[1]
        productLink = "https://www.walmart.com/" + productLink

        if productStatus == None:
            continue
        else:
            productStatus = productStatus.text
            product = {
                "id": productId,
                'image': productImage,
                "name": productName,
                "price": productPrice,      
                "rating": productRating,
                "link": productLink,
                  } 
            p_list.append(product)    
            productId += 1

    
    store = {
        "store": "Walmart",
        "location": Walmart_location,
        "products": p_list
    }

    product_list.append(store)
    
def scrape_target(soup, Target_location):
    p_list = []
    productId = 100

    productInfo = soup.select('ul[class="Row-uds8za-0 jBYETz h-padding-t-tight"] li[class="Col-favj32-0 iXmsJV h-padding-a-none h-display-flex"]')
    for result in productInfo:
        productImage = result.select_one('img')["src"]    
        productNameLink = result.select_one('a[class="Link__StyledLink-sc-4b9qcv-0 styles__StyledTitleLink-h3r0um-1 iBIqkb rwewC h-display-block h-text-bold h-text-bs"]')
        productName = productNameLink.text
        productPrice = result.select_one('div[data-test="current-price"]').span.text
        productRating = result.select_one('span[class="utils__ScreenReaderOnly-sc-1p6kq06-0 iyDnVM"]').text
        productLink = "https://www.target.com"+productNameLink["href"]

        product = {
                "id": productId,
                "image": productImage,
                "name": productName,
                "price": productPrice,      
                "rating": productRating,
                "link": productLink
                  } 
        p_list.append(product)    
        productId += 1
    store = {
        "store": "Target",
        "location": Target_location,
        "products": p_list
    }

    product_list.append(store)


def scrape_walgreens(soup, Walgreens_location):
    p_list = []
    productId = 200

    productInfo = soup.select('div[class="item card card__product in-stores"]')
    for result in productInfo:
        productImage = result.select_one('figure[class="product__img"] img')["src"]
        productImage = "https:" + productImage
        productName = result.select_one('div[name="product-title"]').text
        productPrice = result.select_one('div[class="product__price-contain"] span span span').text
        productRating = result.select_one('span[class="product__rating"] figure img')
        productStatus = result.select_one('div[class="avail-cta"] strong')
        productLink = result.select_one('div[class="product__text"] a')["href"]
        productLink = "https://www.walgreens.com"+productLink
        if productRating == None:
            productRating = "No Rating"
        else:
            productRating = productRating["alt"]

        if productStatus is not None:
            product = {
                "id": productId,
                "image": productImage,
                "name": productName,
                "price": productPrice,      
                "rating": productRating,
                "link": productLink
                } 

            p_list.append(product)    
            productId += 1

        else: 
            continue
    store = {
        "store": "Walgreens",
        "location": Walgreens_location,
        "products": p_list
    }

    product_list.append(store)

if __name__ == '__main__':
          
    with open('product_name.txt') as f:
        info = f.read()

    info = info.split(",")
    product_name = info[0]
    latitude = info[1]
    longitude = info[2]
        
    product_list = []
    serpApi_key = "826f229af259ebcfec493907d7b1f5e923f3efd20b8ea6b67417a6659870e7dc"
    option = webdriver.ChromeOptions()
    # option.add_argument("--headless")

    args = (product_name, option, latitude, longitude)

    t1 = Thread(target=get_walmart, args= args)
    t2 = Thread(target=get_target, args= args)
    t3 = Thread(target=get_walgreens, args= args)
   
    t1.start()
    t2.start()
    t3.start()
    
    t1.join()
    t2.join()
    t3.join()
    
    with open('products.json', 'w') as outfile:
        json.dump(product_list, outfile, indent=4)
    

 

