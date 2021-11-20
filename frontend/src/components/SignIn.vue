<template>
  <div id="app">
   <div class="login-page">
      <transition name="fade">
         <div v-if="!registerActive" class="wallpaper-login"></div>
      </transition>
      <div class="wallpaper-register"></div>

      <div class="container">
         <div class="row">
            <div class="col-lg-4 col-md-6 col-sm-8 mx-auto">
               <div v-if="!registerActive" class="card login" v-bind:class="{ error: emptyFields }">
                  <h1>Sign In</h1>
                  <form class="form-group">
                     <input v-model="emailLogin" type="email" class="form-control" placeholder="Email" required>
                     <input v-model="passwordLogin" type="password" class="form-control" placeholder="Password" required>
                     <input type="submit" class="btn btn-primary" @click.prevent="doLogin">
                     <p>Don't have an account? <a href="#" @click="registerActive = !registerActive, emptyFields = false">Sign up here</a>
                     </p>
                     <p><a href="#">Forgot your password?</a></p>
                  </form>
               </div>

               <div v-else class="card register">
                  <h1>Sign Up</h1>
                  <form class="form-group">
                     <input v-model="emailReg" type="email" class="form-control" placeholder="Email" required>
                     <input v-model="nameReg" type="name" class="form-control" placeholder="Name" required>
                     <input v-model="passwordReg" type="password" class="form-control" placeholder="Password" required>
                     <input v-model="confirmReg" type="password" class="form-control" placeholder="Confirm Password" required>
                     <input type="submit" class="btn btn-primary" @click.prevent="doRegister">
                     <p>Already have an account? <a href="#" @click="registerActive = !registerActive">Sign in here</a>
                     </p>
                  </form>
               </div>
            </div>
         </div>

      </div>
   </div>
</div>
</template>

<script>
import axios from 'axios';
import { mapActions } from 'vuex'


export default {
  name: 'SignIn',
  components: {
    
  },
  data() {
    return {
      registerActive: false,
      emailLogin: "",
      passwordLogin: "",
      emailReg: "",
      nameReg: "",
      passwordReg: "",
      confirmReg: "",
      emptyFields: false,
    }
  },
  computed:{
  },

  methods: {
   ...mapActions([ 'signIn' ]),

    doLogin() {
       
         if (this.emailLogin === "" || this.passwordLogin === "") {
            this.emptyFields = true;
            return
         } 
         var userData = {
                "email": this.emailLogin,
                "password": this.passwordLogin,
            }
            axios({ method: "POST", url: "https://localhost:8000/login", data: userData, headers: {"content-type": "application/json" } })
            .then((response)=> {
                     window.alert("Welcome to Octopus, " + response.data.name + "!")
                     this.$store.dispatch('signIn',response.data);
                     this.$router.push({name:'home'})
                     //this.$forceUpdate();
                  
                  })
               .catch((error) => {
                  window.alert(error.response.data.error);
}
            )
         },

      doRegister() {
         if (this.nameReg === "" || this.emailReg === "" || this.passwordReg === "" || this.confirmReg === "") {
            window.alert("Please fill in all fields!")
            return
         } 
         if (this.passwordReg != this.confirmReg){
            window.alert("Please confirm your password again!");
            return
         }
         
         var userData = {
                "name": this.nameReg,
                "email": this.emailReg,
                "password": this.passwordReg,
                "watchlist": [],
            }
            axios({ method: "POST", url: "https://localhost:8000/signup", data: userData, headers: {"content-type": "application/json" } })
            .then((response)=> {
                  window.alert(response.data.message);
                  if (response.data.message === "Signed up successfully!"){
                     this.registerActive = false;
                     this.emailReg = "";
                     this.nameReg = "";
                     this.passwordReg = "";
                     this.confirmReg = "";
                     }
                  this.registerActive=false
                 
                  
               })
               .catch((error) => {
                  window.alert(error.response.data.error);
            })
   
  }
  }
}


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
p {
   line-height: 1rem;
}

.card {
    margin:auto;
    margin-top:200px;
}

.form-group {
   input {
      margin-bottom: 20px;
   }
}

.login-page {
   
   
   .fade-enter-active,
   .fade-leave-active {
  transition: opacity .5s;
}
   .fade-enter,
   .fade-leave-to {
      opacity: 0;
   }
   
   .wallpaper-register {
      
      height: 100%;
      position: flex;
      width: 100%;
      z-index: -1;
   }

   h1 {
      margin-bottom: 1.5rem;
   }
}

.error {
   animation-name: errorShake;
   animation-duration: 0.3s;
}

@keyframes errorShake {
   0% {
      transform: translateX(-25px);
   }
   25% {
      transform: translateX(25px);
   }
   50% {
      transform: translateX(-25px);
   }
   75% {
      transform: translateX(25px);
   }
   100% {
      transform: translateX(0);
   }
}
</style>
