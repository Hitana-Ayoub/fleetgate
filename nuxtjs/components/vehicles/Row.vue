<script setup lang="ts">
interface RowProps {
  index?: number;
  id?: number;
  plate?: string;
  brand?: string;
  driver?: string;
  status?: string;
  isHeader?: boolean;
}
const props = defineProps<RowProps>();
</script>

<template>
  <div class="row" v-if="isHeader">
    <h4 class="header">Plate number</h4>
    <h4 class="header">Brand</h4>
    <h4 class="header">Driver</h4>
    <h4 class="header">Status</h4>
 </div>
  <div
    v-else
    class="row"
    :style="index % 2 === 0 ? { background: 'rgba(128,128,128,0.15)' } : null"
  >
    <NuxtLink :to="`/fleet/${plate}`" class="header plate link">{{
      plate
    }}</NuxtLink>
    <h4 class="header">{{ brand }}</h4>
    <h4 class="header plate">{{ driver }}</h4>
    <h4 class="header">{{ status }}</h4>
  </div>
</template>

<style scoped>
.row {
  border-top: 1px solid rgba(128, 128, 128, 0.336);
  display: flex;
  align-items: center;
}
.header {
  font-size: 1.25rem;
  color: rgb(10, 31, 148);
  margin: 0;
  margin-right: 2rem;
  padding: 0.5rem 1rem;
  font-weight: 700;
  width: 4rem;
}
.plate {
  color: black;
}
.link {
  text-decoration: none;
  color: rgb(228, 157, 27);
  width: auto;
  font-weight: 500;
}
</style>