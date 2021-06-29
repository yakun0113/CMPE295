import requests
from bs4 import BeautifulSoup

ss = requests.session()
headers = {"referer":"https://www.walmart.com/search",
           "User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36"}

urlFirst = "https://www.walmart.com/search"
url = "https://www.walmart.com/search/?query=toilet%20paper"

ss.get(urlFirst, headers=headers)
res = ss.get(url, headers=headers)
soup = BeautifulSoup(res.text, 'html.parser')
productInfoList = soup.select('li div div[class="search-result-gridview-item clearfix arrange-fill"]')

for product in productInfoList:
    productName = product.select_one('div[class="search-result-product-title gridview"]').a.text
    productPrice = product.select_one('span[class="price-group"]').text
    print(productName)
    print(productPrice)
    print('=======')