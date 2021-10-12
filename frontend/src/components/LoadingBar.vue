<template>
    <div class="container">
        <h6>Fetching information for you...</h6>
        <div>{{percent}}%</div>
        <div class="loading-bar">
            <div class="percentage" :style="{'width': percentage +'%'}"></div>
        </div>
    </div>
</template>

<script>
    export default{
    
        props:{
            percentage:{
                default:0,
                type:Number
            }
        },
        
        computed:{
            percent(){
                return this.percentage.toFixed();
            }
        },
       
      }
</script>

<style lang="scss" scoped>
  .container{
      text-align:center;
      font-size: 2rem;
      color:#555;
  }
  .loading-bar{
      position: relative;
      left: 33%;
      width: 33%;
      height: 20px;
      border-radius: 15px;
      overflow:hidden;
      border-bottom: 1px soild #ddd;
      box-shadow: inset 0 1px 2px rgba($color: #000, $alpha: .4), 0 -1px 1px #fff, 0 1px 0 #fff; 
  }

  .percentage{
      position: relative;
      top: 1px; left: 1px; right: 1px;
      display: block;
      height: 100%;
      border-radius: 15px;
      background-color:#a5df41;
      background-size: 10px 10px;
      background-image: linear-gradient(135deg, rgba($color: #fff, $alpha: .15) 25%,
            transparent 50%, rgba($color: #fff, $alpha: .15) 50%, rgba($color: #fff, $alpha: .15) 75%, transparent 75%,
            transparent);
      animation: animate-stripes 3s linear infinite;
  }

  @keyframes animate-stripes {
      0% {background-position: 0 0; }
      100% {background-position: 60px 0; }
  }

</style>