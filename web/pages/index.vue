<template>
  <v-row>
    <v-col cols="12">
      <template v-if="loggedIn">
        <v-container>
          <v-row dense>
            <v-col>
              <v-row>
                <v-col> <h3 class="headline">My Vehicles</h3> </v-col>
                <v-spacer />
                <v-col cols="auto">
                  <v-btn text color="primary" @click="newVehicle">
                    <v-icon class="pr-1">mdi-plus</v-icon> New
                  </v-btn>
                </v-col>
              </v-row>
              <v-row>
                <v-col v-for="v in vehicles.slice(0, 3)" :key="v.id" cols="4">
                  <vehicle :vehicle="v" read-only />
                </v-col>
              </v-row>
            </v-col>
          </v-row>
          <v-divider class="my-3" />
          <v-row dense>
            <v-col>
              <v-row>
                <v-col>
                  <h3 class="headline">My Service Types</h3>
                </v-col>
                <v-spacer />
                <v-col cols="auto">
                  <v-btn text color="primary" @click="newServiceType">
                    <v-icon class="pr-1">mdi-plus</v-icon> New
                  </v-btn>
                </v-col>
              </v-row>
              <v-row>
                <v-col
                  v-for="st in serviceTypes.slice(0, 3)"
                  :key="st.id"
                  cols="4"
                >
                  <service-type :service-type="st" read-only />
                </v-col>
              </v-row>
            </v-col>
          </v-row>
          <v-divider class="my-3" />
          <v-row dense>
            <v-col>
              <v-row>
                <v-col>
                  <h3 class="headline">My Services</h3>
                </v-col>
                <v-spacer />
                <v-col cols="auto">
                  <v-btn text color="primary" @click="newService">
                    <v-icon class="pr-1">mdi-plus</v-icon> New
                  </v-btn>
                </v-col>
              </v-row>
              <v-row>
                <v-col v-for="s in services.slice(0, 3)" :key="s.id" cols="4">
                  <service :service="s" read-only />
                </v-col>
              </v-row>
            </v-col>
          </v-row>
        </v-container>
      </template>
      <template v-else>
        <v-container>
          <v-row justify="center">
            <v-col cols="12" md="6">
              <v-card>
                <v-card-title>Welcome To Hubcap!</v-card-title>
                <v-card-subtitle>
                  Log in to view your vehicles, service types, and services
                </v-card-subtitle>
                <v-card-text>
                  <p>
                    You may not need hubcaps for your car to run just fine, but
                    it sure looks a lot better when they are there. In the same
                    way, your vehicle service log may
                    <span class="font-italic">work</span>, but it could probably
                    use a solid set of hubcaps. That's where we come in.
                  </p>
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </template>
      <v-dialog v-model="dialog.service" max-width="500">
        <new-service
          v-model="service"
          :loading="loading.service"
          :service-types="serviceTypes"
          :vehicles="vehicles"
          @save="saveService"
          @close="dialog.service = false"
        />
      </v-dialog>
      <v-dialog v-model="dialog.vehicle" max-width="500">
        <new-vehicle
          v-model="vehicle"
          :loading="loading.vehicle"
          @save="saveVehicle"
          @close="dialog.vehicle = false"
        />
      </v-dialog>
      <v-dialog v-model="dialog.serviceType" max-width="500">
        <new-service-type
          v-model="serviceType"
          :loading="loading.serviceType"
          @save="saveServiceType"
          @close="dialog.serviceType = false"
        />
      </v-dialog>
    </v-col>
  </v-row>
</template>

<script>
import api from '~/api'
import NewService from '~/components/services/newService'
import Service from '~/components/services/service'
import NewServiceType from '~/components/serviceTypes/newServiceType'
import ServiceType from '~/components/serviceTypes/serviceType'
import NewVehicle from '~/components/vehicles/newVehicle'
import Vehicle from '~/components/vehicles/vehicle'

export default {
  components: {
    NewServiceType,
    NewVehicle,
    NewService,
    Service,
    ServiceType,
    Vehicle,
  },
  data: () => ({
    fab: false,
    vehicles: [],
    services: [],
    serviceTypes: [],
    service: {},
    vehicle: {},
    serviceType: {},
    dialog: {
      service: false,
      vehicle: false,
      serviceType: false,
    },
    loading: {
      service: false,
      vehicle: false,
      serviceType: false,
    },
  }),
  computed: {
    loggedIn() {
      return !!this.$store.state.auth.user
    },
  },
  watch: {
    loggedIn(newVal) {
      if (newVal) {
        this.init()
      }
    },
  },
  mounted() {
    if (!this.loggedIn) {
      return
    }
    this.init()
  },
  methods: {
    init() {
      this.getServiceTypes()
      this.getServices()
      this.getVehicles()
    },
    getVehicles() {
      api.vehicles.list().then((vehicles) => (this.vehicles = vehicles))
    },
    getServices() {
      api.services.list().then((services) => (this.services = services))
    },
    getServiceTypes() {
      api.serviceTypes
        .list()
        .then((serviceTypes) => (this.serviceTypes = serviceTypes))
    },
    newVehicle() {
      this.vehicle = {}
      this.dialog.vehicle = true
    },
    newService() {
      this.service = {}
      this.dialog.service = true
    },
    newServiceType() {
      this.serviceType = {}
      this.dialog.serviceType = true
    },
    saveVehicle() {
      this.loading.vehicle = true
      api.vehicles
        .create(this.vehicle)
        .then(() => {
          this.dialog.vehicle = false
          this.getVehicles()
        })
        .finally(() => (this.loading.vehicle = false))
    },
    saveService() {
      this.loading.service = true
      api.services
        .create(this.service)
        .then(() => {
          this.dialog.service = false
          this.getServices()
        })
        .finally(() => (this.loading.service = false))
    },
    saveServiceType() {
      this.loading.serviceType = true
      api.serviceTypes
        .create(this.serviceType)
        .then(() => {
          this.dialog.serviceType = false
          this.getServiceTypes()
        })
        .finally(() => (this.loading.serviceType = false))
    },
  },
}
</script>
