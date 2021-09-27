<template>
  <div class="container">
    <Nav/>
    <div class="search-wrapper">
        <input type="text" v-model="search" placeholder="Search everithing at Octopus"/>
        <div><button> Search</button></div>
    </div>

    <h1 class="wgr"><img src='../assets/walgreens.jpg' width="50" height="50" class="image"/>Walgreens</h1>
    
    <div class="test-cards-container">
        <ProductItem 
            v-for="item in walgreensList"
            v-bind:href="item.name"
            :key="item.id"
            :item="item"   
        />
    </div>

    <h1 class="tgt"><img src='../assets/target.jpg'  width="50" height="50" class="image"/>Target</h1>   
    

    <div class="test-cards-container">
        <ProductItem 
            v-for="item in targetList"
            v-bind:href="item.name"
            :key="item.id"
            :item="item"   
        />
    </div>

    <h1 class="wmr"><img src='../assets/walmart.jpg'  width="50" height="50" class="image"/>Walmart</h1>    
    

    <div class="test-cards-container">
        <ProductItem 
            v-for="item in walmartList"
            v-bind:href="item.name"
            :key="item.id"
            :item="item"   
        />
    </div>
  </div>
</template>

<script>
import Nav from './Nav.vue'
import ProductItem from './product/ProductItem.vue'
import { mapGetters } from 'vuex'
export default {
    data() {
        return {
        test: null,
        search: '',
        }
    },
    components: {
        Nav,
        ProductItem
    },
    computed: {
        ...mapGetters(['walmart', 'target', 'walgreens']),
        walmartList() {
            return this.walmart.slice(0,6)
        },
        targetList() {
            return this.target.slice(0,6)
        },
        walgreensList() {
            return this.walgreens.slice(0,6)
        }
    }
}
</script>

<style lang="scss" scoped>
.search-wrapper {
    position: relative;
    margin-top: 22px;
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
  .test-cards-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    padding-top: 10px;
  }
</style>