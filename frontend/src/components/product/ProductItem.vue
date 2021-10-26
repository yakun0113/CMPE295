<template>
  <div class="card">
      <h3>{{item.name}}</h3>
      <img class="image" :src="image" />
      <h5 class="price">{{item.price}}</h5>
      <p class="description">{{item.rating}}</p>
      <a :href="link">Learn More</a>
  <div><button class="add" @click="addToWatchList">Add to watchlist</button></div>
  </div>
</template>

<script>

export default {
    props: [ 'item' ],
    data(){
        return{
            image: this.item.image,
            link: this.item.link
        }
    },
  computed: {
      
      
  },
  methods: {
      goToOrigin(){
          window.location.href = this.item.link;
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
</style>