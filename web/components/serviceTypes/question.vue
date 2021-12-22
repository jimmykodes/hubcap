<template>
  <v-container>
    <v-row dense>
      <v-col>
        <v-text-field
          :value="value.title"
          label="Title"
          @input="emit('title', $event)"
        ></v-text-field>
      </v-col>
      <v-col>
        <v-select
          :items="questionTypes"
          :value="value.type"
          label="Question Type"
          @input="emit('type', $event)"
        ></v-select>
      </v-col>
      <v-col cols="auto" align-self="center">
        <v-btn icon small @click="$emit('delete')">
          <v-icon small>mdi-trash-can-outline</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <template v-if="value.type === 'multiple'">
      <v-row v-for="(o, index) in value.options" :key="index" dense>
        <v-col>
          <v-text-field
            :value="o.text"
            label="Option Text"
            @input="updateOption(index, 'text', $event)"
          ></v-text-field>
        </v-col>
        <v-col>
          <v-text-field
            :value="o.value"
            label="Option Value"
            @input="updateOption(index, 'value', $event)"
          ></v-text-field>
        </v-col>
        <v-col cols="auto" align-self="center">
          <v-btn icon small @click="removeOption(index)">
            <v-icon small>mdi-trash-can-outline</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-row dense>
        <v-col>
          <v-btn @click="addOption">Add Option</v-btn>
        </v-col>
      </v-row>
    </template>
    <template v-else-if="value.type === 'range'">
      <v-row dense>
        <v-col>
          <v-text-field
            :value="value.min"
            label="Min"
            @input="emit('min', $event)"
          ></v-text-field>
        </v-col>
        <v-col>
          <v-text-field
            :value="value.max"
            label="Max"
            @input="emit('max', $event)"
          ></v-text-field>
        </v-col>
        <v-col>
          <v-text-field
            :value="value.step"
            label="Step"
            @input="emit('step', $event)"
          ></v-text-field>
        </v-col>
      </v-row>
    </template>
    <template v-else-if="value.type === 'calculated'">
      <v-row dense>
        <v-col>
          <v-select
            label="Field 1"
            :items="calcOptions"
            :value="value.field1"
            item-text="title"
            item-value="title"
            @input="emit('field1', $event)"
          ></v-select>
        </v-col>
        <v-col>
          <v-select
            label="Operator"
            :items="operators"
            :value="value.operator"
            @input="emit('operator', $event)"
          ></v-select>
        </v-col>
        <v-col>
          <v-select
            label="Field 2"
            :items="calcOptions"
            item-text="title"
            item-value="title"
            :value="value.field2"
            @input="emit('field2', $event)"
          ></v-select>
        </v-col>
      </v-row>
    </template>
  </v-container>
</template>

<script>
import { clone } from 'lodash'

export default {
  name: 'Question',
  props: {
    value: {
      type: Object,
      required: true,
    },
    calcOptions: {
      type: Array,
      required: true,
    },
  },
  data: () => ({
    questionTypes: [
      { text: 'Multiple Choice', value: 'multiple' },
      { text: 'Range', value: 'range' },
      { text: 'Short Text', value: 'text' },
      { text: 'Long Text', value: 'textarea' },
      { text: 'Number', value: 'number' },
      { text: 'True False', value: 'bool' },
      { text: 'Calculated', value: 'calculated' },
    ],
    operators: [
      { text: 'Multiply (*)', value: 'm' },
      { text: 'Divide (/)', value: 'd' },
      { text: 'Add (+)', value: 'a' },
      { text: 'Subtract (-)', value: 's' },
    ],
  }),
  methods: {
    emit(f, v) {
      const q = clone(this.value)
      q[f] = v
      this.$emit('input', q)
    },
    addOption() {
      let o = clone(this.value.options)
      o = o || []
      o.push({})
      this.emit('options', o)
    },
    removeOption(index) {
      const o = clone(this.value.options)
      o.splice(index, 1)
      this.emit('options', o)
    },
    updateOption(index, field, value) {
      const o = clone(this.value.options)
      o[index][field] = value
      this.emit('options', o)
    },
  },
}
</script>

<style scoped></style>
