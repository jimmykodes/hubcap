<template>
  <v-card>
    <v-card-title>{{ value.id ? 'Edit' : 'New' }} Service</v-card-title>
    <v-container>
      <v-row dense>
        <v-col>
          <v-text-field
            :value="value.odometer"
            type="number"
            label="Odometer"
            @input="emit('odometer', $event, number)"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row dense>
        <v-col>
          <v-select
            :value="value.vehicle_id"
            :items="vehicles"
            item-value="id"
            item-text="name"
            label="Vehicle"
            @input="emit('vehicle_id', $event)"
          ></v-select>
        </v-col>
      </v-row>
      <v-row dense>
        <v-col>
          <v-select
            :value="value.service_type_id"
            :items="serviceTypes"
            item-value="id"
            item-text="name"
            label="Service Type"
            @input="emit('service_type_id', $event)"
          ></v-select>
        </v-col>
      </v-row>
      <v-row dense>
        <v-col>
          <v-menu
            v-model="datePicker"
            :close-on-content-click="false"
            :nudge-right="40"
            transition="scale-transition"
            offset-y
            min-width="auto"
          >
            <template #activator="{ on, attrs }">
              <v-text-field
                :value="dateString(value.date)"
                label="Date"
                prepend-icon="mdi-calendar"
                readonly
                v-bind="attrs"
                v-on="on"
              ></v-text-field>
            </template>
            <v-date-picker
              :value="dateString(value.date)"
              no-title
              show-current
              color="primary"
              @input="
                datePicker = false
                emit('date', $event, timestamp)
              "
            ></v-date-picker>
          </v-menu>
        </v-col>
      </v-row>
      <v-row v-for="(q, index) in questions" :key="index" dense>
        <v-col v-if="q.type === 'multiple'">
          <v-radio-group
            column
            :label="q.title"
            :value="getData(q.title)"
            class="px-3"
            @change="emitData(q.title, $event)"
          >
            <v-radio
              v-for="(o, i) in q.options"
              :key="i"
              :label="o.text"
              :value="o.value"
            ></v-radio>
          </v-radio-group>
        </v-col>
        <v-col v-else-if="q.type === 'range'">
          <v-slider
            :value="getData(q.title)"
            class="px-3"
            :min="q.min"
            :max="q.max"
            :step="q.step"
            :label="q.title"
            :tick-labels="tickLabels(q.min, q.max + q.step, q.step)"
            @input="emitData(q.title, $event)"
          ></v-slider>
        </v-col>
        <v-col v-else-if="q.type === 'text'">
          <v-text-field
            :value="getData(q.title)"
            :label="q.title"
            class="px-3"
            outlined
            @input="emitData(q.title, $event)"
          ></v-text-field>
        </v-col>
        <v-col v-else-if="q.type === 'textarea'">
          <v-textarea
            :value="getData(q.title)"
            :label="q.title"
            class="px-3"
            outlined
            @input="emitData(q.title, $event)"
          ></v-textarea>
        </v-col>
        <v-col v-else-if="q.type === 'number'">
          <v-text-field
            :value="getData(q.title)"
            outlined
            :label="q.title"
            type="number"
            class="px-3"
            @input="emitData(q.title, $event, number)"
          ></v-text-field>
        </v-col>
        <v-col v-else-if="q.type === 'bool'">
          <v-switch
            :value="getData(q.title)"
            :label="q.title"
            class="px-3"
            @change="emitData(q.title, $event)"
          ></v-switch>
        </v-col>
        <v-col v-else-if="q.type === 'calculated'">
          <calculated-input
            :label="q.title"
            :field1="number(getData(q.field1))"
            :field2="number(getData(q.field2))"
            :operator="q.operator"
            @input="emitData(q.title, $event)"
          />
        </v-col>
      </v-row>
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
import { every, toNumber, get, find, range, clone } from 'lodash'

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
    value: {
      type: Object,
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
    serviceTypeID: undefined,
    datePicker: false,
  }),
  computed: {
    valid() {
      return every(
        [
          this.value.vehicle_id,
          this.value.service_type_id,
          this.value.date,
          this.value.odometer,
        ],
        (i) => !!i
      )
    },
    questions() {
      const st = find(this.serviceTypes, { id: this.value.service_type_id })
      return get(st, 'questions', [])
    },
  },
  methods: {
    number(val) {
      return toNumber(val)
    },
    timestamp(val) {
      return Math.round(new Date(val).getTime() / 1000)
    },
    dateString(val) {
      if (val) {
        return new Date(val * 1000).toISOString().slice(0, 10)
      }
    },
    getData(key) {
      if (key === 'Odometer') {
        return this.value.odometer
      }
      return get(this.value, `data.${key}`)
    },
    emitData(field, val, transformer) {
      const d = this.value.data || {}
      if (transformer) {
        val = transformer(val)
      }
      d[field] = val
      this.emit('data', d)
    },
    emit(field, val, transformer) {
      const v = clone(this.value)
      if (transformer) {
        val = transformer(val)
      }
      v[field] = val
      this.$emit('input', v)
    },
    tickLabels(start, end, step) {
      return range(start, end, step)
    },
  },
}
</script>
