<template>
  <v-card>
    <v-card-title>New Vehicle</v-card-title>
    <v-container>
      <v-form @submit="$emit('save', vehicle)">
        <v-text-field v-model="name" label="Name"></v-text-field>
        <v-text-field v-model="make" label="Make"></v-text-field>
        <v-text-field v-model="model" label="Model"></v-text-field>
        <v-text-field v-model="year" label="Year" type="number"></v-text-field>
      </v-form>
    </v-container>
    <v-card-actions>
      <v-spacer />
      <v-btn @click="$emit('close')">Cancel</v-btn>
      <v-btn
        :disabled="!valid"
        color="primary"
        :loading="loading"
        @click="$emit('save', vehicle)"
      >
        Save
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { every, toNumber } from 'lodash'

export default {
  name: 'NewVehicle',
  props: {
    loading: {
      type: Boolean,
      default: () => false,
    },
  },
  data: () => ({
    name: undefined,
    make: undefined,
    model: undefined,
    year: undefined,
  }),
  computed: {
    valid() {
      return every([this.name, this.make, this.model, this.year], (i) => !!i)
    },
    vehicle() {
      return {
        name: this.name,
        make: this.make,
        model: this.model,
        year: toNumber(this.year),
      }
    },
  },
}
</script>

<style scoped></style>
