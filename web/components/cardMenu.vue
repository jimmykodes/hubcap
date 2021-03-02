<template>
  <v-menu left offset-x>
    <template #activator="{ on, attrs }">
      <v-icon v-bind="attrs" v-on="on">mdi-dots-vertical</v-icon>
    </template>
    <v-list>
      <v-list-item @click="$emit('edit')">
        <v-list-item-title>Edit</v-list-item-title>
      </v-list-item>
      <v-list-item @click="$emit('delete')">
        <v-list-item-title>Delete</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>

<script>
export default {
  name: 'CardMenu',
}
</script>

<style scoped></style>
