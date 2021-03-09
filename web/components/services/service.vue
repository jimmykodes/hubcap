<template>
  <v-card>
    <v-card-title>
      {{ service.service_type_name }} - {{ service.vehicle_name }}
      <v-spacer />
      <card-menu
        @edit="$emit('edit', service)"
        @delete="$emit('delete', service)"
      />
    </v-card-title>
    <v-card-text>
      <div>Odometer: {{ service.odometer }}</div>
      <div>Date: {{ new Date(service.date * 1000).toLocaleDateString() }}</div>
      <div v-for="(v, k) in service.data" :key="k">{{ k }}: {{ v }}</div>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'Service',
  props: {
    service: {
      type: Object,
      required: true,
    },
  },
}
</script>

<style scoped></style>
