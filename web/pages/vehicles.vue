<template>
  <v-container>
    <v-row>
      <v-col v-for="v in vehicles" :key="v.id" cols="12" md="6">
        <vehicle :vehicle="v" />
      </v-col>
      <v-col v-if="vehicles.length === 0">
        <v-list>
          <v-list-item>
            <v-list-item-content class="font-italics">
              No Vehicles Yet
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
    <v-row>
      <v-spacer />
      <v-btn text color="primary" @click.stop="dialog = true">
        <v-icon class="pr-1">mdi-plus</v-icon>
        New Vehicle
      </v-btn>
    </v-row>
    <v-dialog v-model="dialog" max-width="500px">
      <new-vehicle
        :loading="loading.save"
        @close="dialog = false"
        @save="saveVehicle"
      />
    </v-dialog>
  </v-container>
</template>

<script>
import NewVehicle from '~/components/vehicles/newVehicle'
import Vehicle from '~/components/vehicles/vehicle'

export default {
  name: 'Vehicles',
  components: { Vehicle, NewVehicle },
  data: () => ({
    dialog: false,
    loading: {
      vehicles: false,
      save: false,
    },
    vehicles: [],
  }),
  created() {
    this.getVehicles()
  },
  methods: {
    getVehicles() {
      this.loading.vehicles = true
      return this.$axios
        .$get('/vehicles')
        .then((vehicles) => (this.vehicles = vehicles))
        .catch((err) => console.error(err))
        .finally(() => {
          this.loading.vehicles = false
          this.dialog = false
        })
    },
    saveVehicle(vehicle) {
      this.loading.save = true
      this.$axios
        .$post('/vehicles', vehicle)
        .then(() => this.getVehicles())
        .catch((err) => console.error(err))
        .finally(() => (this.loading.save = false))
    },
  },
}
</script>

<style scoped></style>
