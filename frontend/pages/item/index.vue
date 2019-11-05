<template>
  <div>
    <ItemCard @event="moveTo" />
  </div>
</template>

<script>
import ItemCard from '~/components/item/ItemCard.vue'

export default {
  components: { ItemCard },

  methods: {
    moveTo() {
      console.log('moveTo!')
      this.$router.push('/item/new')
    }
  }
}
</script>
