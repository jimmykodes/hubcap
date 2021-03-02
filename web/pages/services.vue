<template>
  <v-container>
    <v-row>
      <v-col v-for="s in services" :key="s.id" cols="12" md="6">
        <service :service="s" />
      </v-col>
      <v-col v-if="services.length === 0">
        <v-list>
          <v-list-item>
            <v-list-item-content class="font-italics">
              No Services Yet
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
    <v-row>
      <v-spacer />
      <v-btn text color="primary" @click.stop="dialog = true">
        <v-icon class="pr-1">mdi-plus</v-icon>
        New Service
      </v-btn>
    </v-row>
    <v-dialog v-model="dialog" max-width="500px">
      <new-service
        :vehicles="vehicles"
        :service-types="serviceTypes"
        :loading="loading.save"
        @close="dialog = false"
        @save="saveService"
      />
    </v-dialog>
  </v-container>
</template>

<script>
import Service from '~/components/services/service'
import NewService from '~/components/services/newService'
export default {
  name: 'Services',
  components: { NewService, Service },
  data: () => ({
    loading: {
      vehicles: false,
      serviceTypes: false,
      services: false,
      save: false,
    },
    dialog: false,
    vehicles: [],
    serviceTypes: [],
    services: [],
  }),
  created() {
    this.getServices()
    this.getServiceTypes()
    this.getVehicles()
  },
  methods: {
    getVehicles() {
      this.$axios
        .$get('/vehicles')
        .then((vehicles) => (this.vehicles = vehicles))
        .catch((err) => console.error(err))
        .finally(() => (this.loading.vehicles = false))
    },
    getServiceTypes() {
      this.$axios
        .$get('/service_types')
        .then((serviceTypes) => (this.serviceTypes = serviceTypes))
        .catch((err) => console.error(err))
        .finally(() => (this.loading.serviceTypes = false))
    },
    getServices() {
      this.loading.services = true
      this.$axios
        .$get('/services')
        .then((services) => (this.services = services))
        .catch((err) => console.error(err))
        .finally(() => (this.loading.services = false))
    },
    saveService(service) {
      this.loading.save = false
      this.$axios
        .$post('/services', service)
        .then(() => {
          this.dialog = false
          this.getServices()
        })
        .catch((err) => console.error(err))
        .finally(() => (this.loading.save = false))
    },
  },
}
</script>

<style scoped></style>
