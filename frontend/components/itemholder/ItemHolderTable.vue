<template>
  <div>
    <p><v-text-field v-model="search" label="Filter"></v-text-field></p>
    <p>
      <v-data-table
        :headers="headers"
        :items="itemholders"
        :items-per-page="10"
        :search="search"
      >
      </v-data-table>
    </p>
  </div>
</template>

<script>
import queryItemHolders from '~/apollo/queries/itemHolders.gql'
export default {
  data() {
    return {
      search: '',
      headers: [
        { text: 'ID', align: 'center', sortable: true, value: 'id' },
        { text: 'Name', align: 'center', sortable: true, value: 'name' },
        {
          text: 'Nickname',
          align: 'center',
          sortable: true,
          value: 'nickname'
        }
      ],
      itemholders: []
    }
  },

  apollo: {
    itemholders: {
      query: queryItemHolders
    }
  }
}
</script>
