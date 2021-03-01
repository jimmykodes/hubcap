<template>
  <v-card>
    <v-card-title>
      {{ service.service_type_id }} - {{ service.vehicle_id }}
    </v-card-title>
    <v-card-text>
      {{ service.odometer }}
      {{ service.date }}
      {{ service.data }}
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
