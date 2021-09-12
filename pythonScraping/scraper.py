import requests
import json
from bs4 import BeautifulSoup
from selenium import webdriver
import time
from threading import Thread
from multiprocessing import Process


def get_walmart(product_name, option, latitude, longitude):

    Walmart = requests.get("https://serpapi.com/search.json?engine=google_maps&q=walmart&ll=%40"+latitude+"%2C"+longitude+"%2C15z&type=search&api_key=" + serpApi_key)
    Walmart = Walmart.json()
    Walmart_url = Walmart["local_results"][0]["website"]
    Walmart_url = Walmart_url.split("?")[0] + "/search?query=" + product_name
    Walmart_location = Walmart["local_results"][0]["gps_coordinates"]
    product_list["Walmart"]["location"] = Walmart_location

    driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/pythonScraping/chromedriver", options = option)
    driver.get(Walmart_url)
    button = driver.find_element_by_xpath('//*[@id="content"]/div[2]/section[2]/div[2]/div/div/div/div/div/div[2]/form/div/div[3]/button')
    button.click()
    for i in range(10):
        driver.execute_script("window.scrollBy(0, document.body.scrollHeight/10)")
        time.sleep(0.1)
    walmart_soup = BeautifulSoup(driver.page_source, 'html.parser')
    driver.quit()
    scrape_walmart(walmart_soup)

def get_target(product_name, option, latitude, longitude):

    Target = requests.get("https://serpapi.com/search.json?engine=google_maps&q=target&ll=%40"+latitude+"%2C"+longitude+"%2C15.1z&type=search&api_key=5c7e9eb2cdd847fc042006bf5ef4734994c8ce655aa67c0c27122d8c37d2368e")
    Target = Target.json()
    url = Target["local_results"][0]["place_id_search"]
    page = requests.get(url+"&api_key="+serpApi_key)
    page = page.json()
    Target_url=page["place_results"]["website"]
    Target_location = Target["local_results"][0]["gps_coordinates"]
    product_list["Target"]["location"] = Target_location

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
    scrape_target(target_soup)

def get_walgreens(product_name, option, latitude, longitude):

    Walgreens = requests.get("https://serpapi.com/search.json?engine=google_maps&q=walgreens&ll=%40"+latitude+"%2C"+longitude+"%2C15z&type=search&api_key=" + serpApi_key)
    Walgreens = Walgreens.json()
    Walgreens_url = Walgreens["local_results"][0]["website"]
    Walgreens_location = Walgreens["local_results"][0]["gps_coordinates"]
    product_list["Walgreens"]["location"] = Walgreens_location
    
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
    scrape_walgreens(walgreens_soup)

def scrape_walmart(soup):
    
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

        productImage = "https://www.walmart.com" + productImage
        productLink = "https:" + productLink

        if productStatus == None:
            continue
        else:
            productStatus = productStatus.text
            product = {
                'image': productImage,
                "name": productName,
                "price": productPrice,      
                "rating": productRating,
                "link": productLink,
                  } 
            product_list["Walmart"]['products'].append(product)    
    
def scrape_target(soup):

    productInfo = soup.select('ul[class="Row-uds8za-0 jBYETz h-padding-t-tight"] li[class="Col-favj32-0 iXmsJV h-padding-a-none h-display-flex"]')
    for result in productInfo:
        productImage = result.select_one('img')["src"]    
        productNameLink = result.select_one('a[class="Link__StyledLink-sc-4b9qcv-0 styles__StyledTitleLink-h3r0um-1 iBIqkb rwewC h-display-block h-text-bold h-text-bs"]')
        productName = productNameLink.text
        productPrice = result.select_one('div[data-test="current-price"]').span.text
        productRating = result.select_one('span[class="utils__ScreenReaderOnly-sc-1p6kq06-0 iyDnVM"]').text
        productLink = "https://www.target.com"+productNameLink["href"]

        product = {
                "image": productImage,
                "name": productName,
                "price": productPrice,      
                "rating": productRating,
                "link": productLink
                  } 
        product_list["Target"]['products'].append(product)    

def scrape_walgreens(soup):

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
                "image": productImage,
                "name": productName,
                "price": productPrice,      
                "rating": productRating,
                "link": productLink
                } 

            product_list["Walgreens"]['products'].append(product)

        else: 
            continue

if __name__ == '__main__':

    with open('product_name.txt') as f:
        info = f.read()

    info = info.split(",")
    product_name = info[0]
    latitude = info[1]
    longitude = info[2]

    serpApi_key = "9772af35a91fc905851c21a281375c414974e6c78f1a8ffe6ae1670bf3148eb5"

    product_list = {"Walmart":{"location":"", "products":[]}, "Target":{"location":"", "products":[]}, "Walgreens":{"location":"", "products":[]}}

    option = webdriver.ChromeOptions()
    option.add_argument("--headless")
    
    t1 = Thread(target=get_walmart, args=(product_name, option, latitude, longitude))
    t2 = Thread(target=get_target, args=(product_name, option, latitude, longitude))
    t3 = Thread(target=get_walgreens, args=(product_name, option, latitude, longitude))
   
    t1.start()
    t2.start()
    t3.start()

    t1.join()
    t2.join()
    t3.join()

    with open('products.json', 'w') as outfile:
        json.dump(product_list, outfile, indent=4)
