<template>
  <v-layout>
    <v-text-field v-model="search" label="Filter"></v-text-field>
    <v-data-table
      :headers="headers"
      :items="items"
      :items-per-page="2"
      :search="search"
    >
    </v-data-table>
  </v-layout>
</template>

<script>
import queryItems from '~/apollo/queries/items.gql'
export default {
  data() {
    return {
      search: '',
      headers: [
        { text: 'ID', align: 'center', sortable: true, value: 'id' },
        { text: 'Name', align: 'center', sortable: true, value: 'name' },
        { text: 'Price', align: 'center', sortable: true, value: 'price' },
        {
          text: 'ItemHolder ID',
          align: 'center',
          sortable: true,
          value: 'itemHolder.id'
        },
        {
          text: 'ItemHolder Name',
          align: 'center',
          sortable: true,
          value: 'itemHolder.name'
        },
        {
          text: 'ItemHolder Nickname',
          align: 'center',
          sortable: true,
          value: 'itemHolder.nickname'
        }
      ],
      items: []
    }
  },

  apollo: {
    items: {
      query: queryItems
    }
  }
}
</script>
