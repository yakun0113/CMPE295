<template>
<div class="container">
  <Nav :isLoggedIn="isLoggedIn" @signOut = "isLoggedIn = false"/>

  <router-view @signIn = "isLoggedIn = true" />

</div>

</template>

<script>
import Nav from './components/Nav.vue'
//import GoogleMaps from './components/GoogleMaps.vue'
export default {
  data(){
    return{
      isLoggedIn:false
    }
  },
  components: {
    Nav,
    //GoogleMaps,
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
