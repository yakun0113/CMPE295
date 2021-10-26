from scraper import scrape
from http.server import BaseHTTPRequestHandler, HTTPServer
import logging
import json
class S(BaseHTTPRequestHandler):
    def _set_response(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()

    def do_POST(self):
        content_length = int(self.headers['Content-Length']) # <--- Gets the size of data
        post_data = self.rfile.read(content_length) # <--- Gets the data itself
        post_data = post_data.decode('utf-8')
        post_data = json.loads(post_data)      
        product_name = post_data["itemName"]
        latitude = post_data["latitude"]
        longitude = post_data["longitude"]
        #product_data = scrape(product_name, latitude, longitude)
        #product_data = scrape(product_name, "37.3382", "-121.8863")

        
        product_data = [
        {
        "store": "Walmart",
         "location": {
             "latitude": 37.3304293,
             "longitude": -121.86017489999999
         },
         "products": [
             {
                 "id": 0,
                 "image": "https://i5.walmartimages.com/asr/b6ccc571-3db7-4c95-98c4-a18dc0f8ef85_2.890efc06d40e935239c92d473006aaea.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
                 "name": "All Natural* 93% Lean/7% Fat Lean Ground Beef Tray, 1 lb",
                 "price": "$6.40",
                 "rating": "3.6 Stars. 33 reviews.",
                 "link": "https://www.walmart.com//ip/All-Natural-93-Lean-7-Fat-Lean-Ground-Beef-Tray-1-lb/824841960"
             }
             ]
             },
        {
         "store": "Target",
         "location": {
             "latitude": 37.3386461,
             "longitude": -121.9044629
         },
         "products": [
             {
                 "id": 100,
                 "image": "https://target.scene7.com/is/image/Target/GUEST_a302f4ac-2ce0-44d3-b41b-e3c423f7248b",
                 "name": "All Natural 93/7 Ground Beef - 1lb - Good & Gather\u2122",
                 "price": "$6.69",
                 "rating": "3.6 out of 5 stars with 82 ratings",
                 "link": "https://www.target.com/p/all-natural-93-7-ground-beef-1lb-good-38-gather-8482/-/A-13288296#lnk=sametab"
             }]}, 
             {
         "store": "Walgreens",
         "location": {
             "latitude": 37.3237775,
             "longitude": -121.91337499999999
         },
         "products": [
             {
                 "id": 200,
                 "image": "https://pics.drugstore.com/prodimg/413351/450.jpg",
                 "name": "Dinty MooreBeef Stew\u00a0\u00a0-\u00a010 Ounces  ",
                 "price": "2/$5.00 or 1/$2.99",
                 "rating": "4.4 out of 5 total 67 reviews",
                 "link": "https://www.walgreens.com/store/c/dinty-moore-beef-stew/ID=prod1495467-product"
             }]}]
        

        self._set_response()
        product_data = json.dumps(product_data)
        product_data = product_data.encode('utf-8')
        self.wfile.write(product_data)

def run(server_class=HTTPServer, handler_class=S, port=8000):
    logging.basicConfig(level=logging.INFO)
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    logging.info(' Starting scraping server at port '+ str(port))
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        pass
    httpd.server_close()
    logging.info('Stopping httpd...\n')

if __name__ == '__main__':
    from sys import argv

    if len(argv) == 2:
        run(port=int(argv[1]))
    else:
        run()