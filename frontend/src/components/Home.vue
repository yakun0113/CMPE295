<template>
  <div class="container">
    <div class="search-wrapper">
        <form v-on:submit.prevent="submitForm">

            <input type="text" name="itemName" v-model="itemName" class="input" placeholder="Search everithing at Octopus"/>
            <input type="text" name="latitude" v-model="latitude" class="input" placeholder="Latitude" />
            <input type="text" name="longitude" v-model="longitude" class="input" placeholder="Longitude" />
            <div><button class="sm" @click="locateMe">Get my location</button></div>           
            <div><button class="sm" @click="submit">Search</button></div>
           
        </form>
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
        submit: function(){
            this.$router.push('/search-result');
            const json = require('../Store/products.json');
            this.$store.dispatch('setProduct',json);
            var data = {
                "itemName": this.itemName,
                "latitude": (this.latitude),
                "longitude": (this.longitude),
            }
            axios({ method: "POST", url: "https://localhost:8080/search", data: data, headers: {"content-type": "text/plain" } })
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
    background-color: rgb(124, 168, 107);
    font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;

}
.search-wrapper {
    position: relative;
    margin-top: 100px;
    input {
      padding: 3px;
      width: 200px;
      height: 25px;
      color: rgba(0,0,0,.70);
      border: 1px solid rgba(0,0,0,.12);
      transition: .15s all ease-in-out;
      background: white;
        margin:0.1rem;
      &::placeholder{
        font-size: 11px;
        color: rgba(10, 10, 10, 0.5);
        font-weight: 600;
        font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;

      }
    }
  }
  .image{
      vertical-align:middle;
  }
    .wgr{
        background-color: rgb(61, 92, 138);
        color: rgb(255, 255, 255);
        padding: 5px;
    }

  .test-cards-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    padding-top: 10px;
  }
</style>