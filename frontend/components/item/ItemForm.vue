<template>
  <div>
    <v-row>
      <v-form>
        <v-text-field v-model="item.name" label="商品名"></v-text-field>
        <v-text-field v-model="item.price" label="価格"></v-text-field>
        <v-select
          v-model="item.itemHolderId"
          label="所有者"
          :items="itemHolders"
          item-text="name"
          item-value="id"
        ></v-select>
      </v-form>
    </v-row>
    <v-row>
      <v-col><v-btn @click="$emit('cancel')">CANCEL</v-btn></v-col>
      <v-col><v-btn @click="save">SAVE</v-btn></v-col>
    </v-row>
  </div>
</template>

<script>
import ItemHolders from '~/apollo/queries/itemHolders.gql'
import CreateItem from '~/apollo/mutations/items.gql'

export default {
  props: {
    itemId: {
      type: String,
      default: () => ''
    }
  },

  data() {
    return {
      item: {
        name: '',
        price: 0,
        itemHolderId: ''
      }
    }
  },

  methods: {
    async save() {
      try {
        const response = await this.$apollo.mutate({
          mutation: CreateItem,
          variables: {
            name: this.item.name,
            price: this.item.price,
            itemHolderId: this.item.itemHolderId
          }
        })
        console.log(response)
        this.$toast.success('商品を登録しました。')
        this.$emit('save')
      } catch (error) {
        console.log(error)
        this.$toast.success('商品の登録に失敗しました。')
      }
    }
  },

  apollo: {
    itemHolders: {
      prefetch: true,
      query: ItemHolders
    }
  }
}
</script>
