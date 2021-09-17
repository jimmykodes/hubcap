import axios from 'axios'

const $axios = axios.create({ baseURL: '/api' })

function endpoints(url) {
  return {
    list() {
      return $axios
        .get(url)
        .then((result) => result.data)
        .catch((err) => console.error(err))
    },
    delete(id) {
      return $axios.delete(`${url}/${id}`).catch((err) => console.error(err))
    },
    create(obj) {
      return $axios.post(url, obj).catch((err) => console.error(err))
    },
    update(id, obj) {
      return $axios.put(`${url}/${id}`, obj).catch((err) => console.error(err))
    },
  }
}

export default {
  vehicles: endpoints('/vehicles'),
  services: endpoints('/services'),
  serviceTypes: endpoints('/service_types'),
}
