export const state = () => ({
  loading: false,
})

export const mutations = {
  start(state) {
    state.loading = true
  },
  end(state) {
    state.loading = false
  },
}
