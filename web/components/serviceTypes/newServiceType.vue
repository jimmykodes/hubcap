<template>
  <v-card>
    <v-card-title>{{ value.id ? 'Edit' : 'New' }} Service Type</v-card-title>
    <v-container>
      <v-row>
        <v-col>
          <v-text-field
            :value="value.name"
            label="Name"
            @input="emit('name', $event)"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-text-field
            :value="value.freq_miles"
            label="Frequency - Miles"
            type="number"
            @input="emit('freq_miles', $event)"
          ></v-text-field>
        </v-col>
        <v-col>
          <v-text-field
            :value="value.freq_days"
            label="Frequency - Days"
            type="number"
            @input="emit('freq_days', $event)"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-card-title> Addition Questions </v-card-title>
      <v-row v-if="!value.questions || value.questions.length === 0">
        <v-col> No addition questions </v-col>
      </v-row>
      <v-row v-for="(q, index) in value.questions" :key="q.id" dense>
        <v-col>
          <v-divider></v-divider>
          <question
            :value="value.questions[index]"
            @input="updateQuestion(index, $event)"
            @delete="deleteQuestion(index)"
          />
        </v-col>
      </v-row>
    </v-container>
    <v-card-actions>
      <v-btn text color="primary" @click="addQuestion">
        <v-icon color="pr-2">mdi-plus</v-icon>
        New Question
      </v-btn>
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
import Question from '@/components/serviceTypes/question'

export default {
  name: 'NewServiceType',
  components: { Question },
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
    addQuestion() {
      let q = clone(this.value.questions)
      q = q || []
      q.push({})
      this.emit('questions', q)
    },
    updateQuestion(index, question) {
      const q = clone(this.value.questions)
      q[index] = question
      this.emit('questions', q)
    },
    deleteQuestion(index) {
      const q = clone(this.value.questions)
      q.splice(index, 1)
      this.emit('questions', q)
    },
  },
}
</script>
