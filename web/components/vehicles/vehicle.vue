<template>
  <v-card>
    <v-card-title>
      {{ vehicle.name }}
      <v-spacer />
      <v-menu left offset-x>
        <template #activator="{ on, attrs }">
          <v-icon v-bind="attrs" v-on="on">mdi-dots-vertical</v-icon>
        </template>
        <v-list>
          <v-list-item @click="$emit('edit', vehicle)">
            <v-list-item-title>Edit</v-list-item-title>
          </v-list-item>
          <v-list-item @click="$emit('delete', vehicle)">
            <v-list-item-title>Delete</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-card-title>
    <v-card-subtitle>
      {{ vehicle.year }} {{ vehicle.make }} {{ vehicle.model }}
    </v-card-subtitle>
  </v-card>
</template>

<script>
export default {
  name: 'Vehicle',
  props: {
    vehicle: {
      type: Object,
      required: true,
    },
  },
}
</script>

<style scoped></style>
