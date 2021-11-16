<template>
  <div class="container">
    <div class='account-wrapper'>
      <h1>My account</h1>
      <h3>Name:</h3>
      <input type="name" v-model="userName" :placeholder="getUser.name"/>
      <h3>Email:</h3>
      <h5>{{getUser.email}}</h5>
      <h3>New password:</h3>
      <input type="password" v-model="newPassword" placeholder= "Enter new password"/>
      <h3>Confirm new password:</h3>
      <input type="password" v-model="confirmPassword" placeholder= "Confirm new password"/>
      <h3>Current password: </h3>
      <input type="password" v-model="currentPassword" placeholder= "Enter current password"/>
      <div><button class="sm" @click="updateAccount">Update Account</button></div>
      <div><button class="sm" @click="deleteAccount">Delete Account</button></div>           
    </div> 
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import axios from 'axios';

export default {
  data(){
    
    return{
        newPassword: "",
        confirmPassword: "",
        currentPassword: "",
        userName: "",
        userData: null,

    } },
  computed: {
  ...mapGetters(['user']),
  
    getUser(){
      return this.user
    },
  },
  methods: {
  
    updateAccount(){

      if (this.userName === ""){
        this.userName = this.getUser.name
      }
      if(this.newPassword === "" && this.confirmPassword != ""){
          window.alert("Please enter a new password")
          return
      }

      if(this.newPassword != "" && this.confirmPassword === ""){
        window.alert("Please confirm your new password")
        return
      }
      if(this.newPassword != this.confirmPassword){
        window.alert("Confirmation doesn't match with new password")
        return
      }

      if (this.currentPassword === ""){
        window.alert("Enter current password before updating")
        return
      }

      if (this.newPassword === ""){
         this.userData = {
          "newName": this.userName,
          "currentPassword": this.currentPassword,
          }
      }
      else{
          this.userData = {
          "newName": this.userName,
          "currentPassword": this.currentPassword,
          "newPassword": this.newPassword
          }
      }
        axios({ method: "PUT", url: "https://localhost:8000/users/" + this.getUser.user_id, data: this.userData, headers: {"content-type": "application/json", "token":this.getUser.token } })
            .then((response)=> {
              window.alert(response.data.message)                  
                  })
            .catch((error) => {
                  window.alert(error.response.data.error);
}
            )
    },

    deleteAccount(){
       if(confirm("Do you really want to delete?")){
           axios({ method: "DELETE", url: "https://localhost:8000/users/" + this.getUser.user_id, headers: {"content-type": "application/json", "token":this.getUser.token } })
            .then((response)=> {
              window.alert(response.data.message)     
              this.$store.dispatch('signOut');
              this.$router.push('/')             
                  
                  })
            .catch((error) => {
                  window.alert(error.response.data.error);
}
            )
    }
    }
  }
}
</script>

<style lang="scss" scoped>

.account-wrapper {
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
</style>