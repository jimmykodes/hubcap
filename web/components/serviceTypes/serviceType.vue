<template>
  <v-card min-height="100%">
    <v-card-title>
      {{ serviceType.name }}
    </v-card-title>
    <v-card-text>{{ text }}</v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'ServiceType',
  props: {
    serviceType: {
      type: Object,
      required: true,
    },
  },
  computed: {
    text() {
      const hasMiles = this.serviceType.freq_miles !== 0
      const hasDays = this.serviceType.freq_days !== 0
      let str = ''
      if (hasDays) {
        str += `${this.serviceType.freq_days} days`
      }
      if (hasDays && hasMiles) {
        str += ' or '
      }
      if (hasMiles) {
        str += `${this.serviceType.freq_miles} miles`
      }
      return str
    },
  },
}
</script>

<style scoped></style>
