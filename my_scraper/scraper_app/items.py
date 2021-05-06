from scrapy.item import Item, Field

class LivingSocialDeal(Item):
    """Livingsocial container (dictionary-like object) for scraped data"""
    name = Field()
    link = Field()
    location = Field()
    original_price = Field()
    deal_price = Field()
    brand = Field()
    description = Field()
    img = Field()