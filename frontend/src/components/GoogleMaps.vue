<template>
  <GMapMap
    :center="center"
    :options="options"
    :zoom="10"
    map-type-id="terrain"
    style="width: 100vw; height: 20rem"
  >
    <GMapCluster :zoomOnClick="true">
      <GMapMarker
        :key="index"
        v-for="(m, index) in markers"
        :position="m.position"
        :clickable="true"
        :draggable="true"
        @click="center = m.position"
      />
    </GMapCluster>
  </GMapMap>
</template>


<script>
export default {
  data() {
    return {
      center: {lat: 37.3380242, lng: -121.9187964},
      markers: [
        {
          position: {
            lat: 37.3380242, lng: -121.9187964
          },
        }
        , // Along list of clusters
      ]
    }
  }
}
</script>

