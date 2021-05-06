from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine.url import URL
from sqlalchemy import create_engine, Column, Integer, String, DateTime, text, delete
import scraper_app.settings


DeclarativeBase = declarative_base()

def db_connect():
    """
    Performs database connection using database settings from settings.py.
    Returns sqlalchemy engine instance
    """
    return create_engine(URL(**scraper_app.settings.DATABASE))

def create_deals_table(engine):
    sql = text('DROP TABLE IF EXISTS items;')
    result = engine.execute(sql)
    DeclarativeBase.metadata.create_all(engine)

class Deals(DeclarativeBase):
    """Sqlalchemy deals model"""
    __tablename__ = "items"

    id = Column(Integer, primary_key=True)
    name = Column('title', String)
    brand = Column('brand',String, nullable=True)
    link = Column('link', String, nullable=True)
    location = Column('location', String, nullable=True)
    original_price = Column('original_price',String, nullable=True)
    deal_price = Column('deal_price',String, nullable=True)
    description = Column('description',String, nullable=True)
    img = Column('img',String, nullable=True)