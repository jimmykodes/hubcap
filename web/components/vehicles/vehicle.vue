<template>
  <v-card>
    <v-card-title>
      {{ vehicle.name }}
      <v-spacer />
      <card-menu
        @edit="$emit('edit', vehicle)"
        @delete="$emit('delete', vehicle)"
      />
    </v-card-title>
    <v-card-subtitle>
      {{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}
    </v-card-subtitle>
  </v-card>
</template>

<script>
export default {
  name: 'Vehicle',
  props: {
    vehicle: {
      type: Object,
      required: true,
    },
  },
}
</script>

<style scoped></style>
