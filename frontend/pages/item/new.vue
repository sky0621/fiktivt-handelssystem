<template>
  <div>
    <ItemForm @save="moveToList" @cancel="moveToList" />
  </div>
</template>

<script>
import ItemForm from '~/components/item/ItemForm.vue'

export default {
  components: { ItemForm },

  methods: {
    moveToList() {
      this.$router.push('/item')
    }
  }
}
</script>
