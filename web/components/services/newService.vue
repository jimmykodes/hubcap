<template>
  <v-card>
    <v-card-title>New Service</v-card-title>
    <v-container>
      <v-form>
        <v-text-field
          v-model="odometer"
          type="number"
          label="Odometer"
        ></v-text-field>
        <v-select
          v-model="vehicle"
          :items="vehicles"
          item-value="id"
          item-text="name"
          label="Vehicle"
        ></v-select>
        <v-select
          v-model="serviceType"
          :items="serviceTypes"
          item-value="id"
          item-text="name"
          label="Service Type"
        ></v-select>
        <v-date-picker v-model="date"></v-date-picker>
        <v-textarea v-model="data"></v-textarea>
      </v-form>
    </v-container>
    <v-card-actions>
      <v-spacer />
      <v-btn @click="$emit('close')">Cancel</v-btn>
      <v-btn
        :loading="loading"
        :disabled="!valid"
        color="primary"
        @click="$emit('save', service)"
      >
        Save
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { every, toNumber } from 'lodash'

export default {
  name: 'NewService',
  props: {
    vehicles: {
      type: Array,
      required: true,
    },
    serviceTypes: {
      type: Array,
      required: true,
    },
    loading: {
      type: Boolean,
      required: false,
      default() {
        return false
      },
    },
  },
  data: () => ({
    date: undefined,
    odometer: undefined,
    serviceType: undefined,
    vehicle: undefined,
    data: undefined,
  }),
  computed: {
    valid() {
      return every(
        [this.vehicle, this.serviceType, this.date, this.odometer],
        (i) => !!i
      )
    },
    service() {
      const data = {
        service_type_id: this.serviceType,
        vehicle_id: this.vehicle,
        odometer: toNumber(this.odometer),
        date: Math.round(new Date(this.date).getTime()),
      }
      try {
        data.data = JSON.parse(this.data)
      } catch {
        data.data = this.data
      }
      return data
    },
  },
}
</script>
