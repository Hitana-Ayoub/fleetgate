<script setup lang="ts">
import vehicles from "@/data.json";
const VehiclesList = {
  cars: [...vehicles].splice(0, 5),
  trucks: [...vehicles].splice(5, 4),
};
</script>

<template>
  <div class="table">
    <h1>The Fleet:</h1>
    <div class="table-container">
      <div class="table-col">
        <h2> Cars</h2>
        <VehiclesRow :isHeader="true" />
        <VehiclesRow

          v-for="(vehicle, index) in VehiclesList.cars"
          :key="vehicle.id"
          :id="vehicle.id"
          :plate="vehicle.plate"
          :brand="vehicle.brand"
          :driver="vehicle.driver"
          :status="vehicle.status"
          :index="index"
        />
      </div>
      <div class="table-col">
        <h2>Trucks</h2>
        <VehiclesRow :isHeader="true" />
        <VehiclesRow
          v-for="(vehicle, index) in VehiclesList.trucks"
          :key="vehicle.id"
          :id="vehicle.id"
          :plate="vehicle.plate"
          :brand="vehicle.brand"
          :driver="vehicle.driver"
          :status="vehicle.status"
          :index="index"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.table {
  margin: 3rem 0;
}
.table h1 {
  margin-bottom: 2rem;
}
.table-container {
  display: flex;
  justify-content: space-between;
}
.table-col {
  width: 48%;
}
h2{
    color: #198754;
}
</style>