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

               <div v-else class="card register" v-bind:class="{ error: emptyFields }">
                  <h1>Sign Up</h1>
                  <form class="form-group">
                     <input v-model="emailReg" type="email" class="form-control" placeholder="Email" required>
                     <input v-model="nameReg" type="name" class="form-control" placeholder="Name" required>
                     <input v-model="passwordReg" type="password" class="form-control" placeholder="Password" required>
                     <input v-model="confirmReg" type="password" class="form-control" placeholder="Confirm Password" required>
                     <input type="submit" class="btn btn-primary" @click.prevent="doRegister">
                     <p>Already have an account? <a href="#" @click="registerActive = !registerActive, emptyFields = false">Sign in here</a>
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
    doLogin() {
         if (this.emailLogin === "" || this.passwordLogin === "") {
            this.emptyFields = true;
            return
         } 
         var userData = {
                "email": this.emailLogin,
                "password": this.passwordLogin,
            }
            axios({ method: "POST", url: "https://localhost:8080/signIn", data: userData, headers: {"content-type": "application/json" } })
            .then((response)=> {
                  if (response.data === "Wrong password!"){
                     window.alert(response.data);
                     this.passwordLogin = "";
                     return
                     }
                  else if (response.data === "Account not registered!"){
                     window.alert(response.data);
                     this.emailLogin = "";
                     this.passwordLogin = "";
                     return
                  }
                  else{
                     window.alert("Welcome, "+response.data.name + "!");
                     this.$emit('signIn')
                     this.$router.push({name:'home'})
                  }
               })
               .catch((error) => {
                window.alert(`The API returned an error: ${error.data}`);
            })
         
         }
      },
      
      doRegister() {
         if (this.nameReg === "" || this.emailReg === "" || this.passwordReg === "" || this.confirmReg === "") {
            this.emptyFields = true;
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
            }
            axios({ method: "POST", url: "https://localhost:8080/signUp", data: userData, headers: {"content-type": "application/json" } })
            .then((response)=> {
                  window.alert(response.data);
                  if (response.data === "Signed up successfully!"){
                     this.registerActive = false;
                     }
               })
               .catch((error) => {
                window.alert(`The API returned an error: ${error.data}`);
            })
   
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
