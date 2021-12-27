<template>
  <v-text-field
    outlined
    readonly
    class="px-3"
    :label="label"
    :value="value()"
  ></v-text-field>
</template>

<script>
export default {
  name: 'CalculatedInput',
  props: {
    label: {
      type: String,
      required: true,
    },
    field1: {
      type: Number,
      required: true,
      default: 0,
    },
    field2: {
      type: Number,
      required: true,
      default: 0,
    },
    operator: {
      type: String,
      required: true,
    },
  },
  data: () => ({
    last: null,
  }),
  methods: {
    value() {
      let val
      switch (this.operator) {
        case 'm':
          val = this.field1 * this.field2
          break
        case 'd':
          val = this.field1 / this.field2
          break
        case 'a':
          val = this.field1 + this.field2
          break
        case 's':
          val = this.field1 - this.field2
          break
      }
      if (val !== this.last) {
        this.last = val
        this.$emit('input', val)
      }
      return val
    },
  },
}
</script>
