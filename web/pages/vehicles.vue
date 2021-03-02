<template>
  <v-container>
    <v-row>
      <v-col v-for="v in vehicles" :key="v.id" cols="12" md="6">
        <vehicle :vehicle="v" @edit="editVehicle" @delete="confirm" />
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
      <v-btn text color="primary" @click.stop="dialog.save = true">
        <v-icon class="pr-1">mdi-plus</v-icon>
        New Vehicle
      </v-btn>
    </v-row>
    <v-dialog v-model="dialog.save" max-width="500px">
      <new-vehicle
        v-model="vehicle"
        :loading="loading.save"
        @close="dialog.save = false"
        @save="saveVehicle"
      />
    </v-dialog>
    <v-dialog v-model="dialog.delete" max-width="500px">
      <v-card>
        <v-card-title> Confirm Delete </v-card-title>
        <v-card-subtitle>
          Are you sure you want to delete this vehicle?
        </v-card-subtitle>
        <v-card-text>
          This will remove the vehicle and all service records for it. This
          cannot be undone.
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn text @click="dialog.delete = false">Cancel</v-btn>
          <v-btn color="error" text @click="deleteVehicle">Delete</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { delay, every, values } from 'lodash'
import NewVehicle from '~/components/vehicles/newVehicle'
import Vehicle from '~/components/vehicles/vehicle'
const url = '/vehicles'
export default {
  name: 'Vehicles',
  components: { Vehicle, NewVehicle },
  data: () => ({
    dialog: {
      save: false,
      delete: false,
    },
    loading: {
      vehicles: false,
      save: false,
    },
    vehicles: [],
    vehicle: {},
  }),
  watch: {
    dialog: {
      deep: true,
      handler(newVal) {
        if (every(values(newVal), (v) => !v)) {
          // everything is false, so we just closed a dialog
          // delay clearing the form so we don't see a flicker
          delay(() => (this.vehicle = {}), 100)
        }
      },
    },
  },
  created() {
    this.getVehicles()
  },
  methods: {
    getVehicles() {
      this.loading.vehicles = true
      return this.$axios
        .$get(url)
        .then((vehicles) => {
          this.vehicles = vehicles
          this.dialog.save = false
        })
        .catch((err) => console.error(err))
        .finally(() => (this.loading.vehicles = false))
    },
    editVehicle(vehicle) {
      this.vehicle = vehicle
      this.dialog.save = true
    },
    confirm(vehicle) {
      this.vehicle = vehicle
      this.dialog.delete = true
    },
    deleteVehicle() {
      this.$axios
        .$delete(`${url}/${this.vehicle.id}`)
        .then(() => {
          this.getVehicles()
          this.dialog.delete = false
        })
        .catch((err) => console.error(err))
    },
    saveVehicle() {
      this.loading.save = true
      this._save()
        .then(() => this.getVehicles())
        .catch((err) => console.error(err))
        .finally(() => (this.loading.save = false))
    },
    _save() {
      if (this.vehicle.id) {
        return this.$axios.$put(`${url}/${this.vehicle.id}`, this.vehicle)
      }
      return this.$axios.$post(url, this.vehicle)
    },
  },
}
</script>

<style scoped></style>
