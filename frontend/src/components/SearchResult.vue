<template>
  <div class="container">
    <div class="search-wrapper">
        <GoogleMaps class = "map" 
                    :latitude="latitude" :longitude="longitude" 
                    :storeLocations="storeLocations"/>
        <h1>Results for {{productName}}</h1>
    </div>
   
    <h1 class="wmr"><img src='../assets/walmart.jpg'  width="50" height="50" class="image"/>Walmart</h1>    
    <div class = "parent">
        <div><unicon class="angle-left-b" @click.stop = "scrollLeft('wleft')" v-if="walmart_index > 0" name="angle-left-b" width="50" height="50" fill="black"></unicon></div>
        <div class="test-cards-container">

            <ProductItem 
                v-for="item in walmartList"
                v-bind:href="item.name"
                :key="item.id"
                :item="item"   
            />
        </div>
        <div><unicon class="angle-right-b" @click.stop = "scrollRight('wright')" v-if="walmart_index + 6 <= walmartLength - 1" name="angle-right-b" width="50" height="50" fill="black"></unicon></div>
    </div>

    <h1 class="tgt"><img src='../assets/target.jpg'  width="50" height="50" class="image"/>Target</h1>   
    
    <div class = "parent">
        <div><unicon class="angle-left-b" @click.stop = "scrollLeft('tleft')" v-if="target_index > 0" name="angle-left-b" width="50" height="50" fill="black"></unicon></div>
        <div class="test-cards-container">
            <ProductItem 
                v-for="item in targetList"
                v-bind:href="item.name"
                :key="item.id"
                :item="item"   
            />
        </div>
        <div><unicon class="angle-right-b" @click.stop = "scrollRight('tright')" v-if="target_index + 6 <= targetLength - 1" name="angle-right-b" width="50" height="50" fill="black"></unicon></div>
    </div>
    <h1 class="wgr"><img src='../assets/walgreens.jpg' width="50" height="50" class="image"/>Walgreens</h1>

    <div class = "parent">
        <div><unicon class="angle-left-b" @click.stop = "scrollLeft('gleft')" v-if="walgreens_index > 0" name="angle-left-b" width="50" height="50" fill="black"></unicon></div>
        <div class="test-cards-container">
            <ProductItem 
                v-for="item in walgreensList"
                v-bind:href="item.name"
                :key="item.id"
                :item="item"   
            />
        </div>
        <div><unicon class="angle-right-b" @click.stop = "scrollRight('gright')" v-if="walgreens_index + 6 <= walgreensLength - 1" name="angle-right-b" width="50" height="50" fill="black"></unicon></div>
    </div>
  </div>
</template>

<script>
import ProductItem from './product/ProductItem.vue'
import GoogleMaps from './GoogleMaps.vue'
import { mapGetters } from 'vuex'
export default {
    name: "search-result",
    data(){
        return{
        walmart_index: 0,
        target_index: 0,
        walgreens_index: 0,

        }
    },
    props:['productName','latitude', 'longitude'],
    components: {
        ProductItem,
        GoogleMaps
    },
    computed: {
        ...mapGetters(['walmart', 'target', 'walgreens',
                       'store_locations']),
        walmartList() {
            return this.walmart.slice(this.walmart_index, this.walmart_index + 6)
        },
        targetList() {
            return this.target.slice(this.target_index, this.target_index+6)
        },
        walgreensList() {
            return this.walgreens.slice(this.walgreens_index, this.walgreens_index+6)
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
        }

    },
   
    methods: {
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
                        console.log(this.walmart_index);

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

        }
    }
    
}
</script>

<style lang="scss" scoped>
.container{
    font-family:'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif
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
    flex-wrap: wrap;
    justify-content: center;
    padding-top: 10px;
  }

  .angle-left-b{
  
    display: block;
    float:left;

  }

  .angle-right-b{
    display: block;
    float:right;
  
  }

  .map{
      left:25%;
  }

</style>