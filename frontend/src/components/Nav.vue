<template>
  <div id="container">
    <div id='nav'>
      <btn id="left" btnColor="btn btn-large btn-info" @click="goHome">Octopus</btn>
      <div>
       <div id="right" class="nav-right">
          <span v-if="isLoggedIn" class="nav-line-0" >My watchlist</span>
          <span v-if="isLoggedIn" class="nav-line-1" >My account</span>
          <span v-if="!isLoggedIn" class="nav-line-2" @click="signIn">Sign in</span>
          <span v-else class="nav-line-2" @click="signOut">Sign out</span>
       </div>
      </div>   
    </div> 
  </div>
</template>

<script>
import btn from './Button.vue'
export default {
  props:{
    isLoggedIn:{
      default: false
    }
  }
  ,
  components: {
    btn
  },
  computed: {

  },
  methods: {
    signIn(){
      this.$router.push('/sign-in')
    },

    signOut(){
      this.$emit('signOut')
      this.$router.push('/')
      window.alert("Signed out successfully!")
    },

    goHome(){
      this.$router.push('/')
    }
  }
}
</script>


<style scoped lange="sass">
#nav {
  z-index: 2;
  display: block;
  padding: 10px;
  width: 100%;
  height: 60px;
  top: 0px;
  left: 0%;
  position: fixed;
  background-color: rgb(101, 159, 192);
  color: white;
  line-height: 20px;
}

#left {
    float: left;
    text-align: left;
    font-size: 35px;
    font-weight: bold;
    padding-top: 8px;
    padding-left: 15px;
    padding-right: 20px;
    font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
}


#right {
    float: right;
    margin-right: 40px;
    margin-top: 15px; 
    color: white;
    font-family:'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
    cursor: pointer;
}

#cart {
  float: right;
  height: 0%;
  padding-top: 12px;
  padding-left: 20px;
  padding-right: 40px;
  margin-top: 0%; 
  margin-right: 5px; 
}
.nav-right{
  display:flex
}
.nav-line-0 {
  float: left;
  clear: both;
  display: block;
  padding-top: 10px;
  padding-right: 20px;
}
.nav-line-1 {
  float: center;
  clear: both;
  display: block;
  padding-top: 10px;
  padding-right: 20px;
}
.nav-line-2 {
  float: right;
  clear: both;
  display: block;
  padding-top: 10px;
  padding-right: 20px;
}
.text-center {
  text-align: center;
}

.btn-circle {
    width: 25px;
    height: 25px;
    border-radius: 50%;
    position: absolute;
    top: -5px;
    right: 12px;
    background-color: rgb(249, 250, 250);
    color: #000;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .leave-enter-active, .leave-leave-active {
    transition: all 3.2s;
  }
  .leave-enter, .leave-leave-to {
    opacity: 0;
    transform: translateX(-50%);
  }
  .parent{
      display:flex;
  }

  .test-cards-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    padding-top: 10px;
  }

  .angle-left-b{
  
    display: block;
    float:left;

  }

  .angle-right-b{
    display: block;
    float:right;
  
  }
</style>
