<template>
  <v-container>
    <page-header title="Service Types" @new="dialog.save = true" />
    <v-row>
      <v-col v-for="st in serviceTypes" :key="st.id" cols="12" md="4">
        <service-type
          :service-type="st"
          @edit="editServiceType"
          @delete="confirm"
        />
      </v-col>
    </v-row>
    <v-row v-if="serviceTypes.length === 0 && !loading.serviceTypes">
      <v-col>
        <v-list>
          <v-list-item class="font-italics">No Service Types</v-list-item>
        </v-list>
      </v-col>
    </v-row>
    <v-dialog v-model="dialog.save" max-width="500px">
      <new-service-type
        v-model="serviceType"
        :loading="loading.save"
        @close="dialog.save = false"
        @save="saveServiceType"
      />
    </v-dialog>
    <v-dialog v-model="dialog.delete" max-width="500px">
      <v-card>
        <v-card-title> Confirm Delete</v-card-title>
        <v-card-subtitle>
          Are you sure you want to delete this service type?
        </v-card-subtitle>
        <v-card-text>
          This will remove the service type and all service records for it. This
          cannot be undone.
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn text @click="dialog.delete = false">Cancel</v-btn>
          <v-btn color="error" text @click="deleteServiceType">Delete</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { delay, every, values } from 'lodash'
import NewServiceType from '~/components/serviceTypes/newServiceType'
import ServiceType from '~/components/serviceTypes/serviceType'
import PageHeader from '~/components/pageHeader'
import api from '~/api'

export default {
  name: 'ServiceTypes',
  components: { PageHeader, NewServiceType, ServiceType },
  data: () => ({
    dialog: {
      save: false,
      delete: false,
    },
    serviceTypes: [],
    serviceType: {},
    loading: {
      serviceTypes: false,
      save: false,
    },
  }),
  watch: {
    dialog: {
      deep: true,
      handler(newVal) {
        if (every(values(newVal), (v) => !v)) {
          // everything is false, so we just closed a dialog
          // delay clearing the form so we don't see a flicker
          delay(() => (this.serviceType = {}), 100)
        }
      },
    },
    'loading.serviceTypes'(newVal) {
      if (newVal) {
        this.$store.commit('loading/start')
      } else {
        this.$store.commit('loading/end')
      }
    },
  },
  created() {
    this.getServiceTypes()
  },
  methods: {
    getServiceTypes() {
      this.loading.serviceTypes = true
      api.serviceTypes
        .list()
        .then((res) => (this.serviceTypes = res))
        .finally(() => (this.loading.serviceTypes = false))
    },
    editServiceType(serviceType) {
      this.serviceType = serviceType
      this.dialog.save = true
    },
    deleteServiceType() {
      api.serviceTypes.delete(this.serviceType.id).then(() => {
        this.getServiceTypes()
        this.dialog.delete = false
      })
    },
    confirm(serviceType) {
      this.serviceType = serviceType
      this.dialog.delete = true
    },
    _save() {
      if (this.serviceType.id) {
        return api.serviceTypes.update(this.serviceType.id, this.serviceType)
      }
      return api.serviceTypes.create(this.serviceType)
    },
    saveServiceType() {
      this.loading.save = true
      this._save()
        .then(() => {
          this.getServiceTypes()
          this.dialog.save = false
        })
        .finally(() => (this.loading.save = false))
    },
  },
}
</script>

<style scoped></style>
