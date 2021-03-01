<template>
  <v-card>
    <v-card-title>{{ value.id ? 'Edit' : 'New' }} Vehicle</v-card-title>
    <v-container>
      <v-form @submit="$emit('save')">
        <v-text-field
          :value="value.name"
          label="Name"
          @input="emit('name', $event)"
        ></v-text-field>
        <v-text-field
          :value="value.make"
          label="Make"
          @input="emit('make', $event)"
        ></v-text-field>
        <v-text-field
          :value="value.model"
          label="Model"
          @input="emit('model', $event)"
        ></v-text-field>
        <v-text-field
          :value="value.year"
          label="Year"
          type="number"
          @input="emit('year', $event)"
        ></v-text-field>
      </v-form>
    </v-container>
    <v-card-actions>
      <v-spacer />
      <v-btn @click="$emit('close')">Cancel</v-btn>
      <v-btn
        :disabled="!valid"
        color="primary"
        :loading="loading"
        @click="$emit('save')"
      >
        Save
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { every, clone, toNumber } from 'lodash'

export default {
  name: 'NewVehicle',
  props: {
    loading: {
      type: Boolean,
      default: () => false,
    },
    value: {
      type: Object,
      required: true,
    },
  },
  computed: {
    valid() {
      return every(
        [this.value.name, this.value.make, this.value.model, this.value.year],
        (i) => !!i
      )
    },
  },
  methods: {
    emit(field, newValue) {
      const v = clone(this.value)
      if (field === 'year') {
        newValue = toNumber(newValue)
      }
      v[field] = newValue
      this.$emit('input', v)
    },
  },
}
</script>

<style scoped></style>
