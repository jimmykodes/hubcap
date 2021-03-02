<template>
  <v-card>
    <v-card-title>{{ value.id ? 'Edit' : 'New' }} Service Type</v-card-title>
    <v-container>
      <v-text-field
        :value="value.name"
        label="Name"
        @input="emit('name', $event)"
      ></v-text-field>
      <v-text-field
        :value="value.freq_miles"
        label="Frequency - Miles"
        type="number"
        @input="emit('freq_miles', $event)"
      ></v-text-field>
      <v-text-field
        :value="value.freq_days"
        label="Frequency - Days"
        type="number"
        @input="emit('freq_days', $event)"
      ></v-text-field>
    </v-container>
    <v-card-actions>
      <v-spacer />
      <v-btn @click="$emit('close')">Cancel</v-btn>
      <v-btn
        :loading="loading"
        :disabled="!valid"
        color="primary"
        @click="$emit('save')"
      >
        Save
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { clone, includes, toNumber } from 'lodash'

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
    value: {
      type: Object,
      required: true,
    },
  },
  computed: {
    valid() {
      return !!this.value.name
    },
  },
  methods: {
    emit(field, val) {
      const v = clone(this.value)
      if (includes(['freq_miles', 'freq_days'], field)) {
        val = toNumber(val)
      }
      v[field] = val
      this.$emit('input', v)
    },
  },
}
</script>
