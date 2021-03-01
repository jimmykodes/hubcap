<template>
  <v-container>
    <v-row>
      <v-col v-for="st in serviceTypes" :key="st.id" cols="12" md="4">
        <service-type :service-type="st" />
      </v-col>
    </v-row>
    <v-row v-if="serviceTypes.length === 0">
      <v-col>
        <v-list>
          <v-list-item class="font-italics">No Service Types</v-list-item>
        </v-list>
      </v-col>
    </v-row>
    <v-row>
      <v-spacer />
      <v-btn color="primary" text @click.stop="dialog = true">
        <v-icon class="pr-1">mdi-plus</v-icon>
        New Service Type
      </v-btn>
    </v-row>
    <v-dialog v-model="dialog" max-width="500px">
      <new-service-type
        :loading="loading.save"
        @close="dialog = false"
        @save="saveServiceType"
      />
    </v-dialog>
  </v-container>
</template>

<script>
import ServiceType from '~/components/serviceTypes/serviceType'
import NewServiceType from '~/components/serviceTypes/newServiceType'
const url = '/service_types'
export default {
  name: 'ServiceTypes',
  components: { NewServiceType, ServiceType },
  data: () => ({
    dialog: false,
    serviceTypes: [],
    loading: {
      serviceTypes: false,
      save: false,
    },
  }),
  created() {
    this.getServiceTypes()
  },
  methods: {
    getServiceTypes() {
      this.loading.serviceTypes = true
      this.$axios
        .$get(url)
        .then((res) => (this.serviceTypes = res))
        .catch((err) => console.error(err))
        .finally(() => (this.loading.serviceTypes = false))
    },
    saveServiceType(serviceType) {
      this.loading.save = true
      this.$axios
        .$post(url, serviceType)
        .then(() => {
          this.getServiceTypes()
          this.dialog = false
        })
        .catch((err) => console.error(err))
        .finally(() => (this.loading.save = false))
    },
  },
}
</script>

<style scoped></style>
