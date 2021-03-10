<template>
  <div>
    <span class="overline">{{ title }}:</span> {{ value }}
  </div>
</template>

<script>
export default {
  name: 'DataItem',
  props: {
    title: {
      type: String,
      required: true,
    },
    value: {
      type: [Number, String],
      required: false,
      default() {
        return undefined
      },
    },
  },
}
</script>
