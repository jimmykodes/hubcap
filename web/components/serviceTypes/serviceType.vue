<template>
  <v-card min-height="100%">
    <v-card-title>
      {{ serviceType.name }}
      <v-spacer />
      <card-menu
        @edit="$emit('edit', serviceType)"
        @delete="$emit('delete', serviceType)"
      />
    </v-card-title>
    <v-card-subtitle>{{ text }}</v-card-subtitle>
    <v-card-text>{{ details }}</v-card-text>
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
    details() {
      let count = 'No'
      let plural = true
      if (this.serviceType.questions && this.serviceType.questions.length > 0) {
        count = `${this.serviceType.questions.length}`
        if (this.serviceType.questions.length === 1) {
          plural = false
        }
      }
      return `${count} Additional ${plural ? 'Questions' : 'Question'}`
    },
  },
}
</script>

<style scoped></style>
