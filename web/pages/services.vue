<template>
  <v-container>
    <page-header title="Services" @new="dialog.save = true" />
    <v-row>
      <v-col v-for="s in services" :key="s.id" cols="12" md="6">
        <service :service="s" @edit="editService" @delete="deleteService" />
      </v-col>
      <v-col v-if="services.length === 0 && !loading.services">
        <v-list>
          <v-list-item>
            <v-list-item-content class="font-italics">
              No Services Yet
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
    <v-dialog v-model="dialog.save" width="500px">
      <new-service
        v-model="service"
        :vehicles="vehicles"
        :service-types="serviceTypes"
        :loading="loading.save"
        @close="dialog.save = false"
        @save="saveService"
      />
    </v-dialog>
    <v-dialog v-model="dialog.delete" max-width="500px">
      <v-card>
        <v-card-title> Confirm Delete</v-card-title>
        <v-card-subtitle>
          Are you sure you want to delete this service?
        </v-card-subtitle>
        <v-card-text> This cannot be undone.</v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn text @click="dialog.delete = false">Cancel</v-btn>
          <v-btn color="error" text @click="doDelete">Delete</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { delay, cloneDeep, every, values } from 'lodash'
import Service from '~/components/services/service'
import NewService from '~/components/services/newService'
import PageHeader from '~/components/pageHeader'
import api from '~/api'

export default {
  name: 'Services',
  components: { PageHeader, NewService, Service },
  data: () => ({
    loading: {
      vehicles: false,
      serviceTypes: false,
      services: false,
      save: false,
      delete: false,
    },
    dialog: {
      delete: false,
      save: false,
    },
    vehicles: [],
    serviceTypes: [],
    services: [],
    service: {},
  }),
  watch: {
    dialog: {
      deep: true,
      handler(newVal) {
        if (every(values(newVal), (v) => !v)) {
          // everything is false, so we just closed a dialog
          // delay clearing the form so we don't see a flicker
          delay(() => (this.service = {}), 100)
        }
      },
    },
    'loading.services'(newVal) {
      if (newVal) {
        this.$store.commit('loading/start')
      } else {
        this.$store.commit('loading/end')
      }
    },
  },
  created() {
    this.getServices()
    this.getServiceTypes()
    this.getVehicles()
  },
  methods: {
    getVehicles() {
      this.loading.vehicles = true
      api.vehicles
        .list()
        .then((vehicles) => (this.vehicles = vehicles))
        .finally(() => (this.loading.vehicles = false))
    },
    getServiceTypes() {
      this.loading.serviceTypes = true
      api.serviceTypes
        .list()
        .then((serviceTypes) => (this.serviceTypes = serviceTypes))
        .finally(() => (this.loading.serviceTypes = false))
    },
    getServices() {
      this.loading.services = true
      api.services
        .list()
        .then((services) => (this.services = services))
        .finally(() => (this.loading.services = false))
    },
    saveService() {
      this.loading.save = false
      this._save()
        .then(() => {
          this.dialog.save = false
          this.getServices()
        })
        .finally(() => (this.loading.save = false))
    },
    _save() {
      if (this.service.id) {
        return api.services.update(this.service.id, this.service)
      }
      return api.services.create(this.service)
    },
    editService(service) {
      this.service = cloneDeep(service)
      this.dialog.save = true
    },
    deleteService(service) {
      this.service = service
      this.dialog.delete = true
    },
    doDelete() {
      api.services
        .delete(this.service.id)
        .then(() => {
          this.dialog.delete = false
          this.getServices()
        })
        .finally(() => (this.loading.delete = false))
    },
  },
}
</script>

<style scoped></style>
