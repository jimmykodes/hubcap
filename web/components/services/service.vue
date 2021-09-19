<template>
  <v-card>
    <v-card-title>
      {{ service.service_type_name }} - {{ service.vehicle_name }}
      <v-spacer />
      <card-menu
        v-if="!readOnly"
        @edit="$emit('edit', service)"
        @delete="$emit('delete', service)"
      />
    </v-card-title>
    <v-card-text>
      <data-item title="Odometer" :value="service.odometer"></data-item>
      <data-item
        title="Date"
        :value="new Date(service.date * 1000).toISOString().slice(0, 10)"
      ></data-item>
      <data-item
        v-for="(v, k) in service.data"
        :key="k"
        :title="k"
        :value="v"
      ></data-item>
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
    readOnly: {
      type: Boolean,
      default: () => false,
    },
  },
}
</script>

<style scoped></style>
