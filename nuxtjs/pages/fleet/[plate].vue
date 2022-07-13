<script setup lang="ts">
import vehicles from "@/data.json";
const route = useRoute();
const plate = route.params.plate;
const vehicle = vehicles.find((v) => v.plate === plate);


useMeta({
  brand: vehicle ? plate : "404 - Vehicle Not Found",
  meta: [
    {
      plate: "viewport",
      content: "width=device-width",
    },
  ],
});
</script>


<template>
  <div>
      <div class="vehicle-container">
        <div class="image-container">
          <img :src="vehicle.imageUrl" alt="" />
        </div>
        <div class="info-container">
          <h1>{{ vehicle.brand }}</h1>
          <div class="stats-container">
            <h6>Plate number:</h6>
            <p>{{ vehicle.plate }}</p>
          </div>
          <div class="stats-container">
            <h6>Name of driver:</h6>
            <p>{{ vehicle.driver }}</p>
          </div>
          <div class="stats-container">
            <h6>Distance in KM:</h6>
            <p>{{ vehicle.km }}</p>
          </div>
          <div class="stats-container">
            <h6>Engine Type:</h6>
            <p>{{ vehicle.engine }}</p>
          </div>
          <div class="stats-container">
            <h6>Status:</h6>
            <p>{{ vehicle.status }}</p>
          </div>
        </div>
      </div>

   
  </div>
</template>

<style scoped>
.vehicle-container {
  display: flex;
}
.image-container {
  width: 75%;
  height: 95vh;
}
.image-container img {
  width: 100%;
  height: 100%;
}
.vehicle-not-found {
  height: 75vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.info-container {
  padding: 3rem;
  width: 50%;
}
.info-container h1 {
  text-transform: uppercase;
  font-size: 6rem;
  margin-bottom: 3rem;
}
.stats-container {
  display: flex;
  align-items: flex-end;
  margin-bottom: 1rem;
}
.stats-container h6 {
  width: 20rem;
  font-size: 2rem;
  margin: 0;
  margin-right: 5rem;
}
.stats-container p {
  font-size: 2rem;
  margin: 0;
}
.content {
  font-size: 1.5rem;
  margin-top: 4rem;
}
img {
  width: 10rem;
}
</style>