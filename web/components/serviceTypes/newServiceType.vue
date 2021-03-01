<template>
  <v-card>
    <v-card-title>New Service Type</v-card-title>
    <v-container>
      <v-text-field v-model="name" label="Name"></v-text-field>
      <v-text-field
        v-model="miles"
        label="Frequency - Miles"
        type="number"
      ></v-text-field>
      <v-text-field
        v-model="days"
        label="Frequency - Days"
        type="number"
      ></v-text-field>
    </v-container>
    <v-card-actions>
      <v-spacer />
      <v-btn @click="$emit('close')">Cancel</v-btn>
      <v-btn
        :loading="loading"
        :disabled="!valid"
        color="primary"
        @click="$emit('save', serviceType)"
      >
        Save
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { toNumber } from 'lodash'

export default {
  name: 'NewServiceType',
  props: {
    loading: {
      type: Boolean,
      required: false,
      default() {
        return false
      },
    },
  },
  data: () => ({
    name: undefined,
    miles: undefined,
    days: undefined,
  }),
  computed: {
    valid() {
      return !!this.name
    },
    serviceType() {
      return {
        name: this.name,
        freq_miles: toNumber(this.miles),
        freq_days: toNumber(this.days),
      }
    },
  },
}
</script>
