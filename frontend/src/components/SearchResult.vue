<template>
  <div class="container">
    <div class="search-wrapper">
        <GoogleMaps class = "map" 
                    :latitude="latitude" :longitude="longitude" 
                    :storeLocations="storeLocations" :storeLinks="storeLinks"/>
        <h1>Results for {{productName}}</h1>
        <div class="dropdown">
            <unicon name="filter"></unicon>
            <div class="dropdown-content">
                <div class="sort" @click="priceSort">
                    <unicon name="sort-amount-up"></unicon>
                    <h5>Price low to high</h5>
                </div>
                <div class="sort" @click="ratingSort">
                    <unicon name="sort-amount-down"></unicon>
                    <h5>Rating high to low</h5>
                </div>
                <div class="sort" @click="noSort">
                    <unicon name="bars"></unicon>
                    <h5>No filter</h5>
                </div>

            </div>
        </div>
    </div>
   
    <h1 class="wmr">
        <img src='../assets/walmart.jpg' width="50" height="50" class="image"/>Walmart</h1>    
    <h3 v-if="storeLinks['walmart']==''">Cannot get store</h3>
    <h3 v-else-if="walmartLength === 0">Cannot get result, please visit the link in the map</h3>

    <div class = "parent">
        <div v-if="isSortByPrice" class="test-cards-container">
            <ProductItem 
                v-for="item in priceSortWalmart"
                v-bind:href="item.name"
                :key="item.id"
                :item="item" 
            />
        </div>
        <div v-else-if="isSortByRating" class="test-cards-container">
            <ProductItem 
                v-for="item in ratingSortWalmart"
                v-bind:href="item.name"
                :key="item.id"
                :item="item" 
            />
        </div>
        <div v-else class="test-cards-container">
            <ProductItem 
                v-for="item in walmartList"
                v-bind:href="item.name"
                :key="item.id"
                :item="item" 
            />
        </div>
    </div>
    <unicon  class="angle-left-b" @click.stop = "scrollLeft('wleft')" v-if="walmart_index > 0 && walmartLength != 0" name="angle-double-left" height="40px" width="40px" fill="black" hover-fill="limegreen"></unicon>
    <unicon  class="angle-right-b" @click.stop = "scrollRight('wright')" v-if="walmart_index + 6 <= walmartLength-1" name="angle-double-right" height="40px" width="40px" fill="black" hover-fill="limegreen"></unicon>


    <h1 class="tgt"><img src='../assets/target.jpg' width="50" height="50" class="image"/>Target</h1>   
    <h3 v-if="storeLinks['target']==''">Cannot get store</h3>
    <h3 v-else-if="targetLength === 0">Cannot get result, please visit the link in the map</h3>
    <div class = "parent">
        <div v-if="isSortByPrice" class="test-cards-container">
            <ProductItem 
                v-for="item in priceSortTarget"
                v-bind:href="item.name"
                :key="item.id"
                :item="item"   
           />
        </div>
        <div v-else-if="isSortByRating" class="test-cards-container">
            <ProductItem 
                v-for="item in ratingSortTarget"
                v-bind:href="item.name"
                :key="item.id"
                :item="item"   
           />
        </div>
        <div v-else class="test-cards-container">
            <ProductItem 
                v-for="item in targetList"
                v-bind:href="item.name"
                :key="item.id"
                :item="item"   
           />
        </div>
    </div>
    <unicon class="angle-left-b" @click.stop = "scrollLeft('tleft')" v-if="target_index > 0" name="angle-double-left" height="40px" width="40px" fill="black" hover-fill="limegreen"></unicon>
    <unicon class="angle-right-b" @click.stop = "scrollRight('tright')" v-if="target_index + 6 <= targetLength+1" name="angle-double-right" height="40px" width="40px" fill="black" hover-fill="limegreen"></unicon>

    <h1 class="wgr"><img src='../assets/walgreens.jpg' width="50" height="50" class="image"/>Walgreens</h1>
    <h3 v-if="storeLinks['walgreens']==''">Cannot get store</h3>
    <h3 v-else-if="walgreensLength === 0">Cannot get result, please visit the link in the map</h3>
    <div class = "parent">    
        <div v-if="isSortByPrice" class="test-cards-container">
            <ProductItem 
                v-for="item in priceSortWalgreens"
                v-bind:href="item.name"
                :key="item.id"
                :item="item" 
            />
        </div>
        <div v-else-if="isSortByRating" class="test-cards-container">
            <ProductItem 
                v-for="item in ratingSortWalgreens"
                v-bind:href="item.name"
                :key="item.id"
                :item="item" 
            />
        </div>
        <div v-else class="test-cards-container">
            <ProductItem 
                v-for="item in walgreensList"
                v-bind:href="item.name"
                :key="item.id"
                :item="item" 
            />
        </div>
    </div>
     
    <unicon class="angle-left-b" @click.stop = "scrollLeft('gleft')" v-if="walgreens_index > 0" name="angle-double-left" height="40px" width="40px"  fill="black" hover-fill="limegreen"></unicon>
    <unicon class="angle-right-b" @click.stop = "scrollRight('gright')" v-if="walgreens_index + 6 < walgreensLength+1" name="angle-double-right" height="40px" width="40px" fill="black" hover-fill="limegreen"></unicon>
    
  </div>
</template>

<script>
import ProductItem from './product/ProductItem.vue'
import GoogleMaps from './GoogleMaps.vue'
import { mapGetters } from 'vuex'
//import Button from './Button.vue'
export default {
    name: "search-result",
    data(){
        return{
        walmart_index: 0,
        target_index: 0,
        walgreens_index: 0,
        isSortByPrice: false,
        isSortByRating: false,
        }
    },
    props:['productName','latitude', 'longitude',],
    components: {
        ProductItem,
        GoogleMaps,
        
        //Button
    },
    computed: {
        ...mapGetters(['walmart', 'target', 'walgreens', 'walmart_sort_by_price', 'target_sort_by_price', 'walgreens_sort_by_price','walmart_sort_by_rating','target_sort_by_rating','walgreens_sort_by_rating',
                       'store_locations','store_links']),
        walmartList() {
            return this.walmart.slice(this.walmart_index, this.walmart_index + 6)
        },
        targetList() {
            return this.target.slice(this.target_index, this.target_index+6)
        },
        walgreensList() {
            return this.walgreens.slice(this.walgreens_index, this.walgreens_index+6)
        },
        priceSortWalmart() {
            return this.walmart_sort_by_price.slice(this.walmart_index, this.walmart_index + 6)
        },
        priceSortTarget(){
            return this.target_sort_by_price.slice(this.target_index, this.target_index + 6)
        },
        priceSortWalgreens(){
            return this.walgreens_sort_by_price.slice(this.walgreens_index, this.walgreens_index + 6)
        },
        ratingSortWalmart(){
            return this.walmart_sort_by_rating.slice(this.walmart_index, this.walmart_index + 6)
        },
        ratingSortTarget(){
            return this.target_sort_by_rating.slice(this.target_index, this.target_index + 6)
        },
        ratingSortWalgreens(){
            return this.walgreens_sort_by_rating.slice(this.walgreens_index, this.walgreens_index + 6)
        },
        walmartLength(){
            return this.walmart.length
        },
        targetLength(){
            return this.target.length
        },
        walgreensLength(){
            return this.walgreens.length
        },
        storeLocations() {
            return this.store_locations
        },
        storeLinks(){
            return this.store_links
        }
    },
   
    methods: {
        noSort(){
            this.walmart_index = 0
            this.target_index = 0
            this.walgreens_index = 0
            this.isSortByRating = false
            this.isSortByPrice = false
         

        },

        priceSort(){
            this.walmart_index = 0
            this.target_index = 0
            this.walgreens_index = 0
            this.isSortByPrice=true
            this.isSortByRating = false
           

        },
        ratingSort(){
            this.walmart_index = 0
            this.target_index = 0
            this.walgreens_index = 0
            this.isSortByRating=true
            this.isSortByPrice=false
           

        },
        scrollRight(id){
           switch (id){
               case "wright":
                   this.walmart_index+=6;
                   break;
               case "tright":
                   this.target_index+=6;
                   break;
               case "gright":
                   this.walgreens_index+=6;
           }
        },
        scrollLeft(id){
            switch (id){
               case "wleft":
                   if(this.walmart_index >= 6){
                        this.walmart_index-=6;
                   }
                   break;   
               case "tleft":
                   if(this.target_index >= 6){
                        this.target_index-=6;
                   }
                   break;
               case "gleft":
                   if(this.walgreens_index >= 6 ){
                        this.walgreens_index-=6;
                   }
           }
        },
       
   
        }
    }
    
</script>

<style lang="scss" scoped>
.container{
    font-family:'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif
}
.dropdown {
  position: relative;
  display: inline-block;
  margin-left: 900px;

}
.dropdown-content {
  display: none;
  position: absolute;
  background-color: #f9f9f9;
  min-width: 200px;
  height: 270px;
  box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
  border-radius: 5px;
  z-index: 2;
}
.sort:hover{
    background-color:rgb(66, 161, 106);
    width: 200px;
    height:90px;
    border-radius: 5px;
    color:white;
}
.dropdown:hover .dropdown-content {
  display: block;
}
.sm{
   border: 0;
    cursor: pointer;
    box-sizing: border-box;
    border-radius: 3px;
    color: white;
    background-color: rgb(106, 165, 82);
    justify-content: space-around;
    position: relative; 
}
.search-wrapper {
    position: relative;
    margin-top: 100px;
    input {
      padding: 3px;
      width: 500px;
      height: 28px;
      color: rgba(0,0,0,.70);
      border: 1px solid rgba(0,0,0,.12);
      transition: .15s all ease-in-out;
      background: white;
      &::placeholder{
        font-size: 12px;
        color: rgba(10, 10, 10, 0.5);
        font-weight: 800;
      }
    }
  }
  .image{
      vertical-align:middle;
  }
    .wgr{
        background-color: rgb(56, 47, 102);
        color: rgb(247, 41, 26);
        padding: 5px;
    }
    .tgt{
        background-color: rgb(196, 16, 16);
        color: white;
        padding: 5px;
    }
    .wmr{
        background-color: rgb(18, 109, 245);
        color: white;
        padding: 5px;
    }
  .parent{
      display:flex;
  }
  .test-cards-container {
    display: flex;
    clear: both;
    flex-wrap: wrap;
    justify-content: center;
    float:center;
  }
  .map{
      left:25%;
  }

 
</style>