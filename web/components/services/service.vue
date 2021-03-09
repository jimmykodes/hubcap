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
      Odometer: {{ service.odometer }}<br />
      Date: {{ new Date(service.date * 1000).toLocaleDateString() }}<br />
      <template v-for="(v, k) in service.data">
        <span :key="k"> {{ k }}: {{ v }} <br /> </span>
      </template>
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
