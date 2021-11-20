<template>
  <div class="card">
      <h3>{{item.name}}</h3>
      <img class="image" :src="item.image" />
      <h5 class="price">{{item.price}}</h5>
      <p class="description">{{item.rating}}</p>
      <a :href="item.link" @click.prevent="visitProductPage(item.link)">Learn More</a>
      <div>
        <button v-if="button != 'Delete'" class="add" @click.prevent="addToWatchlist">Add to watchlist</button>
        <button v-else class="delete" @click.prevent="removeFromWatchlist(item)">Remove from watchlist</button>
      </div>
  </div>
</template>

<script>
import axios from 'axios';
import { mapGetters } from 'vuex'
export default {
    props: [ 'item', 'button' ],
  
  computed: {
     ...mapGetters(['user','logged','watchlist']),
     
     watchList() {
            return this.watchlist.slice(this.watchlist_index, this.watchlist_index + 6)
        },
    //getUser(){
      //  return this.user    
    //},
    //isLoggedIn(){
    //  return this.logged
    //}
      
  },
  methods: {
      //goToOrigin(){
      //    window.location.href = this.item.link;
      //},
      addToWatchlist(){
           var data = {
               "image": this.item.image,
               "name": this.item.name,
               "price": this.item.price,
               "rating": this.item.rating,
               "link": this.item.link,
           }
           const user_id = document.cookie.split('=')[1]

           axios({ method: "POST", url: "https://localhost:8000/watchlist/" + user_id, data: data, headers: {"content-type": "application/json"} })
            .then((response) => {
                this.$toast.show(response.data.message);
                setTimeout(this.$toast.clear,2000)
            })
            .catch((error) => {
                  window.alert(error.response.data.error);
            })
      },
      removeFromWatchlist(item){
           const user_id = document.cookie.split('=')[1]

           axios({ method: "DELETE", url: "https://localhost:8000/watchlist/" + user_id + "/" + item.item_id, headers: {"content-type": "application/json"} })
            .then((response) => {
                const json = response.data.result;
                this.$store.dispatch('setWatchlist',json);
                this.$toast.show(response.data.message);
                setTimeout(this.$toast.clear,2000)
            })
            .catch((error) => {
                  window.alert(error.response.data.error);
            })
  },
    visitProductPage(link){
        window.open(link)
    }
}
}
</script>

<style lang="scss">
h3 {
  text-align: center;
  color: black;
}
    .card{
        width: 80%;
        margin: 10%;
        padding: 10px;
        border-radius: 5px;
        background-color: white;
        box-shadow: 0 0 5px grey;

        h5.price{
            color: grey;
        }

        p.description {
            font-size: .85rem;
            padding: 10px;
        }
        button {
           width: 250px;
           padding-bottom: 10px;
           margin-top: 29px;
        }
       
        button.text-muted {
            color: white;
            background-color: grey;
        }
        
    }
    h4.test-detail-button {
            padding: 10px;
            color: rgb(46, 126, 153);
            font-weight: bold;
            font-size: 1.15rem;    
            cursor: pointer;
        }

    @media (min-width: 500px) {
        .card {
            width: 350px;
            margin: 10px;
        }
    }
    .image{
        z-index:0;
        width:80px;
        height:80px;
    }
.add{
    border: 0;
    cursor: pointer;
    box-sizing: border-box;
    border-radius: 3px;
    color: white;
    background-color: rgb(124, 168, 107);
    font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
}

.delete{
    border: 0;
    cursor: pointer;
    box-sizing: border-box;
    border-radius: 3px;
    color: white;
    background-color:rgb(187, 17, 11);
    font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
}

</style>