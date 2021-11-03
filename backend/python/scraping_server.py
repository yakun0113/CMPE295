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
        file = open("products.json",'r')
        product_data = file.read()
        self._set_response()
        #product_data = json.dumps(product_data)
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


