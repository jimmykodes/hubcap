const userKey = 'user'

function getUser() {
  const user = sessionStorage.getItem(userKey)
  if (!user) {
    return null
  }
  try {
    return JSON.parse(user)
  } catch (e) {
    console.error(e)
    return null
  }
}

export const state = () => ({
  user: getUser(),
})

export const mutations = {
  setUser(state, user) {
    state.user = user
    sessionStorage.setItem(userKey, JSON.stringify(user))
  },
}
