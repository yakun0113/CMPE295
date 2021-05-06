from scrapy.spiders import Spider
from scrapy.selector import Selector
from scrapy.loader import ItemLoader
from itemloaders.processors import MapCompose, Join

from scraper_app.items import LivingSocialDeal

class LivingSocialSpider(Spider):
    """Spider for regularly updated livingsocial.com site, San Francisco Page"""
    name = "livingsocial"
    
    allowed_domains = ["livingsocial.com"]
    start_urls = ["https://www.livingsocial.com/browse/san-jose?lat=37.339&lng=-121.895&query=san+jose&address=San+Jose&division=san-jose&locale=en_US"]

    deals_list_xpath = '//figure'
    item_fields = {
        'name': './/div[@class="cui-udc-title one-line-truncate"]/text()',
        'link': './/div[@class="cui-content"]/a/@href',
        'location': './/span[@class="cui-location-name"]/text()',
        'img': '(.//span[@class="a-list-item"]//img/@src)[1]',
        'brand': '(.//td[@class="a-span9"]/span[@class="a-size-base"]/text())[1]',
        'original_price': './/td/span[@class="priceBlockStrikePriceString a-text-strike"]/text()',
        'deal_price': './/td[@class="a-span12"]/span/text()',
        'description': './/ul[@class="a-unordered-list a-vertical a-spacing-mini"]/li'
        
    }

    """
    #Amazon single product page
    allowed_domains = ["amazon.com"]
    start_urls = ["https://www.amazon.com/OGX-Nourishing-Coconut-Shampoo-Ounce/dp/B000TGEJT8/ref=sr_1_53?dchild=1&keywords=shampoo&qid=1617924870&sr=8-53"]

    deals_list_xpath = '//body'
    item_fields = {
        'name': './/h1/span/text()',
        'img': '(.//span[@class="a-list-item"]//img/@src)[1]',
        'brand': '(.//td[@class="a-span9"]/span[@class="a-size-base"]/text())[1]',
        'link': './/h1/span/text()',
        'location': './/b/text()',
        'original_price': './/td/span[@class="priceBlockStrikePriceString a-text-strike"]/text()',
        'deal_price': './/td[@class="a-span12"]/span/text()',
        'description': '(.//ul[@class="a-unordered-list a-vertical a-spacing-mini"]/li/span[@class="a-list-item"]/text())[1]'
    }
    """
    
    def parse(self, response):
        """
        Default callback used by Scrapy to process downloaded responses

        Testing contracts:

        @url http://www.livingsocial.com/cities/15-san-francisco
        @returns items 1
        @scrapes title link

        """
        selector = Selector(response)

        # iterate over deals
        for deal in selector.xpath(self.deals_list_xpath):
            loader = ItemLoader(LivingSocialDeal(), selector=deal)

            # define processors
            loader.default_input_processor = MapCompose(str.strip)
            loader.default_output_processor = Join()

            # iterate over fields and add xpaths to the loader
            for field, xpath in self.item_fields.items():
                loader.add_xpath(field, xpath)
            yield loader.load_item()
