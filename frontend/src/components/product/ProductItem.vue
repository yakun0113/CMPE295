<template>
  <div class="card">
      <h3>{{item.name}}</h3>
      <img class="image" :src="image" />
      <h5 class="price">{{item.price}}</h5>
      <p class="description">{{item.rating}}</p>
      <a :href="link">Learn More</a>
  <div><button v-if="isLoggedIn && (button==='Add')" class="add" @click.prevent="addToWatchlist">Add to watchlist</button></div>
  <div><button v-if="isLoggedIn && (button==='Delete')" class="delete" @click.prevent="removeFromWatchlist">Remove from watchlist</button></div>
  </div>
</template>

<script>
import axios from 'axios';
import { mapGetters } from 'vuex'
export default {
    props: [ 'item', 'button' ],
    data(){
        return{
            image: this.item.image,
            link: this.item.link,
        }
    },
  computed: {
     ...mapGetters(['user','logged']),
    getUser(){
        return this.user    
    },
    isLoggedIn(){
      return this.logged
    }
      
  },
  methods: {
      goToOrigin(){
          window.location.href = this.item.link;
      },
      addToWatchlist(){
           axios({ method: "PUT", url: "https://localhost:8080/handleWatchlist/"+ this.getUser, data: this.item, headers: {"content-type": "application/json" } })
            .then((response) => {
                this.$toast.show(response.data.message);
                setTimeout(this.$toast.clear,2000)
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error.data}`);
            })
      },
      removeFromWatchlist(){
           axios({ method: "DELETE", url: "https://localhost:8080/handleWatchlist/" + this.getUser + "/" + this.item.link })
            .then((response) => {
                this.$toast.show(response.data.message);
                setTimeout(this.$toast.clear,2000)
            })
            .catch((error) => {
                window.alert(`The API returned an error: ${error.data}`);
            })
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