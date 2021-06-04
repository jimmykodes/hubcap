<template>
  <v-app dark>
    <v-navigation-drawer permanent mini-variant fixed app>
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content></v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar fixed app>
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn v-if="loggedIn" @click="logout">Log Out</v-btn>
      <template v-else>
        <v-btn href="/api/oauth/login/github">Login with GitHub</v-btn>
        <v-btn href="/api/oauth/login/google">Login with Google</v-btn>
      </template>
      <v-progress-linear
        :active="loading"
        indeterminate
        absolute
        bottom
        color="deep-purple accent-4"
      ></v-progress-linear>
    </v-app-bar>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { mapState } from 'vuex'

export default {
  data() {
    return {
      drawer: false,
      allItems: [
        {
          icon: 'mdi-home',
          to: '/',
        },
        {
          icon: 'mdi-car',
          to: '/vehicles',
        },
        {
          icon: 'mdi-cog',
          to: '/serviceTypes',
        },
        {
          icon: 'mdi-format-list-bulleted',
          to: '/services',
        },
        {
          icon: 'mdi-chart-bar',
          to: '/graphs',
        },
      ],
      title: 'Vehicle Maintenance',
    }
  },
  computed: {
    loggedIn() {
      return !!this.$store.state.auth.user
    },
    items() {
      if (!this.loggedIn) {
        return this.allItems.slice(0, 1)
      }
      return this.allItems
    },
    ...mapState('loading', {
      loading: 'loading',
    }),
  },
  created() {
    this.$axios.$get('/users/me').then((user) => {
      this.$store.commit('auth/setUser', user)
    })
  },
  methods: {
    logout() {
      this.$store.commit('auth/clearUser')
      window.location.href = '/api/logout'
    },
  },
}
</script>
