<template>
  <div class="container">
    <div class="search-wrapper">
        <h1>My watchlist</h1>
        <h3 v-if="isEmpty"> No items in your watchlist</h3>
    </div>
   
    <div class = "parent">
        <div><unicon class="angle-left-b" @click.stop = "scrollLeft()" v-if="watchlist_index > 0" name="angle-left-b" width="50" height="50" fill="black"></unicon></div>
        <div class="test-cards-container">

            <ProductItem 
                v-for="item in watchList"
                v-bind:href="item.name"
                :key="item.id"
                :item="item"   
                :button='button'
            />
        </div>
        <div><unicon class="angle-right-b" @click.stop = "scrollRight()" v-if="watchlist_index + 6 <= watchlistLength - 1" name="angle-right-b" width="50" height="50" fill="black"></unicon></div>
    </div>
  </div>
</template>

<script>
import ProductItem from './product/ProductItem.vue'

import { mapGetters } from 'vuex'
export default {
    name: "watchlist",
    data(){
        return{
        watchlist_index: 0,

        }
    },
    props:['button'],
    components: {
        ProductItem,
    },
    computed: {
        ...mapGetters(['watchlist']),
        watchList() {
            return this.watchlist.slice(this.watchlist_index, this.watchlist_index + 6)
        },
 
        watchlistLength(){
            return this.watchlist.length
        },

        isEmpty() {
            if (this.watchlist.length === 0){
                return true
            }
            else{
                return false
            }
        }
    },
   
    methods: {

        scrollRight(){
            this.watchlist_index+=6;
           }
        ,
        scrollLeft(){
            if(this.watchlist_index >= 6){
                this.watchlist_index-=6;
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
    clear: both;
    flex-wrap: wrap;
    justify-content: center;
    float:center;
  }

  .angle-left-b{
  
    display: block;
    float:left;
    clear: both;
    padding-top: 300px;

  }

  .angle-right-b{
    display: block;
    float:right;
    padding-top: 300px;
    clear: both;

  }

  .map{
      left:25%;
  }

</style>