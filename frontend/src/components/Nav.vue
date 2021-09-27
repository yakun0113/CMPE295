<template>
  <div id="container">
    <div id='nav'>
      <h1 id="left"> Octopus</h1>
      <div>
       <div id="right">
          <span class="nav-line-1">Hello,</span>
          <span v-if="!isLoggedIn" class="nav-line-2" @click="signIn">Sign in</span>
          <span v-else class="nav-line-2">{{currentUser}}</span>
       </div>
      </div>   
    </div> 
  </div>
</template>

<script>

export default {
  components: {
    
  },
  computed: {

  },
  methods: {

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
    padding-left: 28px;
}


#right {
    float: right;
    margin-right: 40px;
    margin-top: 15px; 
    color: white;
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
.nav-line-1 {
  float: left;
  font-size: 13px;
  line-height: 8px;
  height: 11px;
  font-weight: 400;
  clear: both;
  display:block;
  
}
.nav-line-2 {
  clear: both;
  display: block;
  padding-right: 10px;
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
</style>
