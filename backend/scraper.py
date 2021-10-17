import requests
import json
from bs4 import BeautifulSoup
from selenium import webdriver
import time
from threading import Thread

product_list = []
serpApi_key = "60d95a6c4f23372236426717be3e4349217eb73796fadb016cbacb0cc71d34ae"
option = webdriver.ChromeOptions()
option.add_argument("--headless")
option.add_argument("window-size=2560,1440")

def get_walmart(product_name, option, latitude, longitude):

    Walmart = requests.get("https://serpapi.com/search.json?engine=google_maps&q=walmart&ll=%40"+latitude+"%2C"+longitude+"%2C15z&type=search&api_key=" + serpApi_key)
    Walmart = Walmart.json()
    Walmart_url = Walmart["local_results"][0]["website"]
    product_name = product_name.replace(" ","+")
    Walmart_url = Walmart_url.split("?")[0] + "/search?query=" + product_name
    Walmart_location = Walmart["local_results"][0]["gps_coordinates"]

    driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/CMPE295-main/backend/chromedriver", options = option)
    driver.get(Walmart_url)
    button = driver.find_element_by_xpath('//*[@id="content"]/div[2]/section[2]/div[2]/div/div/div/div/div/div[2]/form/div/div[3]/button')
    button.click()
    time.sleep(0.5)
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

    driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/CMPE295-main/backend/chromedriver", options = option)
    driver.get(Target_url)
    driver.execute_script("window.scrollBy(0, document.body.scrollHeight/5)")
    time.sleep(1)
    button = driver.find_element_by_xpath("//*[@id='__next']/div[3]/div/div[1]/div[2]/div/div[5]/button[2]")
    button.click()
    product_name = product_name.replace(" ","+")
    driver.get("https://www.target.com/s?searchTerm="+product_name+"&facetedValue=5zkty")
    time.sleep(1)

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
    
    driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/CMPE295-main/backend/chromedriver", options = option)
    driver.get(Walgreens_url) 
    search = driver.find_element_by_xpath('//*[@id="shop-placeholder"]')
    search.send_keys(product_name)
    button = driver.find_element_by_xpath('//*[@id="storeServiceDiv"]/div[1]/div[1]/div[3]/div/div/button')
    button.click()
    time.sleep(1)
    pickUp = driver.find_element_by_xpath('//*[@id="IN_STORES"]')
    pickUp.click()
    time.sleep(1)

    for i in range(15):
        driver.execute_script("window.scrollBy(0, document.body.scrollHeight/15)")
        time.sleep(0.1)

    walgreens_soup = BeautifulSoup(driver.page_source, 'html.parser')
    driver.quit()
    scrape_walgreens(walgreens_soup, Walgreens_location)

def scrape_walmart(soup, Walmart_location):

    p_list = []
    productId = 0
    productInfo = soup.select('a[class="display-block"]')
    
    for result in productInfo:
        productImage = result.select_one('img[class="tile-img"]')["src"]    
        productName = result.select_one('div[class="tile-title"]').text
        productPrice = result.select_one('span[class="price-group"]')
        productRating = result.select_one('span[class="stars-container"]')
        #productStatus = result.select_one('div[class="tile-in-store-info"]')
        productLink = result['href']

        if productPrice == None:
            productPrice = "Price available in store"
        else:
            productPrice = productPrice.text
        if productRating == None:
            productRating = "No rating"
        else:
            productRating = productRating['aria-label']
        productImage = "https:" + productImage
        productLink = "https://www.walmart.com/" + productLink
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

        ''' 
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
    '''
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
        productImage = result.select_one('img')   
        productNameLink = result.select_one('a[class="Link__StyledLink-sc-4b9qcv-0 styles__StyledTitleLink-h3r0um-1 iBIqkb rwewC h-display-block h-text-bold h-text-bs"]')
        productName = productNameLink.text
        productPrice = result.select_one('div[data-test="current-price"]').span
        productRating = result.select_one('span[class="utils__ScreenReaderOnly-sc-1p6kq06-0 iyDnVM"]')
        productLink = "https://www.target.com"+productNameLink["href"]

        if productPrice == None:
            productPrice = "Price available in store"
        else:
            productPrice = productPrice.text

        if productRating == None:
            productRating = "No rating"
        else:
            productRating = productRating.text

        if productImage == None:
            continue
        else:
            productImage = productImage['src']

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
        productImage = result.select_one('figure[class="product__img"] img')
        productName = result.select_one('div[name="product-title"]').text
        productPrice = result.select_one('div[class="product__price-contain"] span span span')
        productRating = result.select_one('span[class="product__rating"] figure img')
        productStatus = result.select_one('div[class="avail-cta"] strong')
        productLink = result.select_one('div[class="product__text"] a')["href"]
        productLink = "https://www.walgreens.com"+productLink

        if productPrice == None:
            productPrice = "Price available in store"
        else:
            productPrice = productPrice.text

        if productRating == None:
            productRating = "No Rating"
        else:
            productRating = productRating["alt"]
        if productImage == None:
            continue
        else:
            productImage = "https:" + productImage['src']

        if productStatus is not None :
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

def scrape(product_name, latitude, longitude):

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
    
    return product_list
    


 

