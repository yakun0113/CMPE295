import requests
import json
from bs4 import BeautifulSoup
from selenium import webdriver
import time
from requests_html import HTMLSession
import socket
from selenium.webdriver.common.action_chains import ActionChains
from multiprocessing import Pool

option = webdriver.ChromeOptions()
driver = webdriver.Chrome(executable_path="/Users/brian80433/Desktop/pythonScraping/chromedriver", options = option)
driver.get("https://www.walmart.com/store/5435-san-jose-ca/search?query="+"beef")

#driver.execute_script("window.scrollBy(0, document.body.scrollHeight/5)")
button = driver.find_element_by_xpath('//*[@id="content"]/div[2]/section[2]/div[2]/div/div/div/div/div/div[2]/form/div/div[3]/button')
button.click()
'''
time.sleep(0.4)
html = driver.execute_script("return document.getElementsByTagName('html')[0].innerHTML")
'''
time.sleep(0.4)
for i in range(10):
        driver.execute_script("window.scrollBy(0, document.body.scrollHeight/10)")
        #time.sleep(0.1)

soup = BeautifulSoup(driver.page_source, 'html.parser')
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

            print(productImage)
            print(productName)
            print(productPrice)
            print(productRating)
            print(productLink)
            print("---------------------------")
        