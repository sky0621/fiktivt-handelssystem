<template>
  <div>
    <ItemCard @addItem="moveToNew" />
  </div>
</template>

<script>
import ItemCard from '~/components/item/ItemCard.vue'

export default {
  components: { ItemCard },

  methods: {
    moveToNew() {
      this.$router.push('/item/new')
    }
  }
}
</script>
