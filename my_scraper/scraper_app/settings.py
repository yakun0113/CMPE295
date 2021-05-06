BOT_NAME = 'livingsocial'

SPIDER_MODULES = ['scraper_app.spiders']

DATABASE = {
    'drivername': 'postgresql',
    'host': 'localhost',
    'port': '5432',
    'username': 'postgres',
    'password': '0617',
    'database': 'octopus'
}

ITEM_PIPELINES = {'scraper_app.pipelines.LivingSocialPipeline':100}
""" http://127.0.0.1:52569/?key=edaf5d69-4eff-482c-8ff0-eec98e721a2d"""