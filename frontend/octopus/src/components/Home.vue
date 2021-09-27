<template>
  <div class="container">
    <Nav/>
    <div class="search-wrapper">
        <input type="text" v-model="search" placeholder="Search everithing at Octopus"/>
        <button> Search</button>
    </div>

    <h1>Walmart</h1>
    <div class="test-cards-container">
        <ProductItem 
            v-for="item in walmartList"
            v-bind:href="item.name"
            :key="item.id"
            :item="item"   
        />
    </div>
    <h1>Target</h1>
    <div class="test-cards-container">
        <ProductItem 
            v-for="item in targetList"
            v-bind:href="item.name"
            :key="item.id"
            :item="item"   
        />
    </div>
    <h1>Wallgreens</h1>
    <div class="test-cards-container">
        <ProductItem 
            v-for="item in wallgreenList"
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
        ...mapGetters(['walmart', 'target', 'wallgreen']),
        walmartList() {
            return this.walmart.slice(0,6)
        },
        targetList() {
            return this.target.slice(0,6)
        },
        wallgreenList() {
            return this.wallgreen.slice(0,6)
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

  .test-cards-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    padding-top: 10px;
  }
</style>