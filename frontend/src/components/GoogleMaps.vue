<template>
  <GMapMap 
    class = "map"
    :center="center"
    :options="options"
    :zoom="13"
    map-type-id="terrain"
    style="width: 100vw; height: 20rem"
  >
      <GMapMarker
        :key="index"
        v-for="(m, index) in marker"
        :position="m.position"
        :clickable="true"
        @click="center = m.position"
      />

      <GMapMarker
        :key="index"
        v-for="(m, index) in walmart_marker"
        :position="m.position"
        :clickable="true"
        :icon= '{
          url: "https://s3.amazonaws.com/www-inside-design/uploads/2018/04/walmart-square.jpg",
          scaledSize: {width: 30, height: 30},
      }'
         @click="visit(storeLinks['walmart'])"
      />

      <GMapMarker
        :key="index"
        v-for="(m, index) in target_marker"
        :position="m.position"
        :clickable="true"
        :icon= '{
          url: "https://i.pinimg.com/originals/11/98/8f/11988f04000898c4106a7dd89f819469.jpg",
          scaledSize: {width: 30, height: 30},
      }'
        @click="visit(storeLinks['target'])"
      />

      <GMapMarker
        :key="index"
        v-for="(m, index) in walgreens_marker"
        :position="m.position"
        :clickable="true"
        :icon= '{
          url: "https://www.tkasphalt.com/wp-content/uploads/2017/03/Walgreens-logo.jpg",
          scaledSize: {width: 30, height: 30},
      }'
        @click="visit(storeLinks['walgreens'])"
      />

  </GMapMap>

</template>


<script>

export default {
  props:["latitude","longitude","storeLocations","storeLinks"],
  data() {
    return {
      
      center: {lat: parseFloat(this.latitude), lng: parseFloat(this.longitude)},
      marker: [
        {
          position: {
            lat: parseFloat(this.latitude), 
            lng: parseFloat(this.longitude),
            label: "Your location"

          }
        }],
        walmart_marker:[
          {
          position: {
             lat: (this.storeLocations['walmart']['latitude']), 
             lng: (this.storeLocations['walmart']['longitude']),
             label: "Walmart",
          },

        }],
        target_marker:[
         {
          position: {
             lat: (this.storeLocations['target']['latitude']), 
             lng: (this.storeLocations['target']['longitude']),
             label: "Target",

          },

          
        }],
        walgreens_marker:[
         {
          position: {
             lat: (this.storeLocations['walgreens']['latitude']), 
             lng: (this.storeLocations['walgreens']['longitude']),
             label: "Walgreens",
          },

        }],

    }

  },
  methods:{
        visit(links){
            window.open(links)
        },
  }
  
}
</script>

<style scoped>
.map{
  position: relative;
  height: 20rem;
  width: 50%;
  left:25%;
}
</style>