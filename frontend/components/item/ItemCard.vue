<template>
  <div>
    <v-row>
      <v-btn
        fab
        color="lightgreen accent-2"
        top
        center
        @click="$emit('addItem')"
      >
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card>
          <v-list two-line subheader>
            <v-list-item v-for="item in items" :key="item.id" link>
              <v-list-item-avatar>
                <v-icon>mdi-gift-outline</v-icon>
              </v-list-item-avatar>
              <v-list-item-content>
                <v-list-item-title>{{ item.name }}</v-list-item-title>
                <v-list-item-subtitle>￥{{ item.price }}</v-list-item-subtitle>
              </v-list-item-content>
              <v-list-item-content>
                <v-list-item-title>
                  {{ item.itemHolder.name }}
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import Items from '~/apollo/queries/items.gql'
export default {
  data() {
    return {
      items: []
    }
  },

  apollo: {
    items: {
      prefetch: true,
      query: Items
    }
  }
}
</script>
