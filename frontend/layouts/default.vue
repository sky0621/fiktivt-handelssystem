<template>
  <v-app>
    <!-- サイドメニュー -->
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
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
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <!-- ヘッダー -->
    <v-app-bar :clipped-left="clipped" fixed app>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
    </v-app-bar>

    <!-- コンテンツ -->
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>

    <!-- フッター -->
    <v-footer :fixed="fixed" app>
      <span>&copy; 2019</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      clipped: false,
      drawer: false,
      fixed: false,
      items: [
        {
          icon: 'mdi-home',
          title: 'Home',
          to: '/'
        },
        {
          icon: 'mdi-gift-outline',
          title: 'Items',
          to: '/item'
        },
        {
          icon: 'mdi-chess-king',
          title: 'ItemHolders',
          to: '/itemholder'
        },
        {
          icon: 'mdi-video-input-component',
          title: 'Composite',
          to: '/composite'
        }
      ],
      miniVariant: false,
      title: 'fiktivt-handelssystem'
    }
  }
}
</script>
