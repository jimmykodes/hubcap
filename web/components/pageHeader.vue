<template>
  <v-row>
    <v-col cols="auto">
      <h2 class="text--h2">{{ title }}</h2>
    </v-col>
    <v-spacer />
    <v-col cols="auto">
      <v-btn text color="primary" @click.stop="$emit('new')">
        <v-icon class="pr-1">mdi-plus</v-icon>
        New
      </v-btn>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'PageHeader',
  props: {
    title: {
      type: String,
      required: true,
    },
  },
}
</script>
