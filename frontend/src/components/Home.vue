<template>
  <div class="container">
    <div class="search-wrapper">
        <h1>Search products near you!</h1>
        <GMapMap 
            class = "maps"
            :center="center"
            :options="options"
            :zoom="zoom"
            map-type-id="terrain"
            style="width: 100vw; height: 20rem"
            @click = "getCoordinates"
        >
            <GMapMarker
            :key="index"
            v-for="(m, index) in marker"
            :position="m.position"
            :clickable="true"
            :draggable="true"
            @click="center = m.position"
      />
        </GMapMap>
            <input type="text" v-model="itemName" placeholder="Search any product"/>
            <div><button class="sm" @click="locateMe">Get my location</button></div>           
            <div><button class="sm" @click="search">Search</button></div>
    </div>
    <LoadingBar v-show = "showBar" :percentage = "percentage"/>

  </div>
</template>

<script>
import LoadingBar from './LoadingBar'
import { mapActions } from 'vuex'
import { mapGetters } from 'vuex'

import axios from 'axios';
export default {
    name:"home",
    data() {
        return {
        test: null,
        itemName: '',
        location:null,
        gettingLocation: false,
        errorStr:null,
        latitude:null,
        longitude:null,
        showBar: false,
        start: false,
        percentage: 0,
        center: {lat: 37.0902 , lng:-95.7129},
        zoom:4,
        marker: [
            {
                position:{
                    lat:null,
                    lng:null,

                }
            }
        ],

        }
    },
    components: {
        LoadingBar,
       // GoogleMaps,
    },

    computed: {
          ...mapGetters(['logged','user']),
        isLoggedIn(){
            return this.logged
        },
        getUser(){
            return this.user    
        },
    },
  
    methods: {
        ...mapActions([ ' setProduct ' ]),
        search(){
            if (this.latitude === null || this.latitude === null){
                window.alert("Please choose a location!")
                return
            }
            if (this.itemName === ""){
                window.alert("Please enter an item to search!")
                return
            }
            
            if (this.isLoggedIn === false){
                window.alert("Sign in to search!")
                return
            }
            this.showBar = true
            var intval = setInterval(()=>{
                    if(this.percentage < 100)
                        this.percentage += .1;
                       
                    else
                        clearInterval(intval);
                },25);

            var data = {
               "itemName": this.itemName,
                "latitude": (this.latitude).toString(),
                "longitude": (this.longitude).toString(),
            }
            axios({ method: "POST", url: "https://localhost:8000/search", data: data, headers: {"content-type": "application/json" } })
            .then((response) => {
                this.percentage = 100
                const json = response.data;
                this.$store.dispatch('setProduct',json);
                this.$router.push({name:'search-result', 
                                   params:{
                                       productName: this.itemName,
                                       latitude: this.latitude,
                                       longitude: this.longitude,
                                       }});
               
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
            this.latitude = this.location.coords.latitude;
            this.longitude = this.location.coords.longitude;
            this.center = {lat: this.latitude, lng: this.longitude};
            this.zoom = 12;
            this.marker[0].position.lat = this.latitude;
            this.marker[0].position.lng = this.longitude;
        } catch(e) {
            this.gettingLocation = false;
            this.errorStr = e.message;
        }
        },

        getCoordinates(event){
            this.latitude = event.latLng.lat();
            this.longitude = event.latLng.lng();
            this.center = {lat: this.latitude, lng: this.longitude};
            this.zoom = 12;
            this.marker[0].position.lat = this.latitude;
            this.marker[0].position.lng = this.longitude;
        }
        
    }
}
</script>

<style lang="scss" scoped>
.container{
    font-family:'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif
}
.maps{
  position: relative;
  height: 20rem;
  width: 50%;
  left:25%;
}
.sm{
    position: relative;
    top:10px;
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


</style>