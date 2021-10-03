<template>
  <div class="container">
    <div class="search-wrapper">
        <form v-on:submit.prevent="submitForm">

            <input type="text" v-model="itemName" placeholder="Search everithing at Octopus"/>
            <input type="text" v-model="latitude" placeholder="Latitude" />
            <input type="text" v-model="longitude" placeholder="Longitude" />
            <div><button class="sm" @click="locateMe">Get my location</button></div>
          
            
            
            <div><button class="sm" @click="submit">Search</button></div>
            <h1>Results for {{itemName}}(1000+)</h1>
        </form>
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
import ProductItem from './product/ProductItem.vue'
import { mapGetters, mapActions } from 'vuex'
import axios from 'axios';

export default {
    data() {
        return {
        test: null,
        itemName: '',
        location:null,
        gettingLocation: false,
        errorStr:null,
        latitude:null,
        longitude:null,
        }
    },
    components: {
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
    },
    methods: {
        ...mapActions([ ' setProduct ' ]),
        submit(){
            const json = require('../Store/products.json');
            this.$store.dispatch('setProduct',json);


            axios.post("https://localhost:8080/search", {
                itemName: this.itemName,
                latitude: this.latitude,
                longitude: this.longitude,
            })
            .then((response) => {
                console.log(response.data)
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error.data}`);
            })

            console.log(this.itemName,this.latitude,this.longitude)

        },
        async getLocation() {
      
        return new Promise((resolve, reject) => {

            if(!("geolocation" in navigator)) {
            reject(new Error('Geolocation is not available.'));
            }

            navigator.geolocation.getCurrentPosition(pos => {
            resolve(pos);
            }, err => {
            reject(err);
            });

        });
        },
        async locateMe() {

        this.gettingLocation = true;
        try {
            this.gettingLocation = false;
            this.location = await this.getLocation();
            this.latitude = (this.location.coords.latitude).toString();
            this.longitude = (this.location.coords.longitude).toString();
        } catch(e) {
            this.gettingLocation = false;
            this.errorStr = e.message;
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
  .test-cards-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    padding-top: 10px;
  }
</style>