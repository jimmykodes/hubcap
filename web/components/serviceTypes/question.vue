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
    <template v-if="isMultipleChoice">
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
    <template v-if="isRange">
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
  },
  data: () => ({
    questionTypes: [
      { text: 'Multiple Choice', value: 'multiple' },
      { text: 'Range', value: 'range' },
      { text: 'Short Text', value: 'text' },
      { text: 'Long Text', value: 'textarea' },
      { text: 'Number', value: 'number' },
      { text: 'True False', value: 'bool' },
    ],
  }),
  computed: {
    isMultipleChoice() {
      return this.value.type === 'multiple'
    },
    isRange() {
      return this.value.type === 'range'
    },
  },
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
