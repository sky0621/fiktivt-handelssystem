<template>
  <div>
    <v-row>
      <v-col
        ><v-text-field v-model="search" label="Filter"></v-text-field
      ></v-col>
    </v-row>
    <v-row>
      <v-col
        ><v-data-table
          :headers="headers"
          :items="itemHolders"
          :items-per-page="10"
          :search="search"
        >
        </v-data-table
      ></v-col>
    </v-row>
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
      itemHolders: []
    }
  },

  apollo: {
    itemHolders: {
      prefetch: true,
      query: queryItemHolders
    }
  }
}
</script>
