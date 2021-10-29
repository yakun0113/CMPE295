<template>
  <div id="container">
    <div id='account'>
      <btn id="left" btnColor="btn btn-large btn-info" @click="goHome">Octopus</btn>
      <div>
       <div id="right" class="nav-right">
          <span v-if="isLoggedIn" class="nav-line-0" @click="watchlist">My watchlist</span>
          <span v-if="isLoggedIn" class="nav-line-1" @click="account">My account</span>
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
    },

    watchlist(){
      this.$router.push('/watchlist')
    },

    account(){
      this.$router.push('/account')
      }
  }
}
</script>